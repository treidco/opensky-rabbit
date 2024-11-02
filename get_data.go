package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getData() []Flight {
	end := time.Now().Unix()
	begin := end - 3600

	url := fmt.Sprintf("https://opensky-network.org/api/flights/all?begin=%d&end=%d", begin, end)
	username := os.Getenv("OPENSKY_USERNAME")
	password := os.Getenv("OPENSKY_PASSWORD")
	if username == "" || password == "" {
		log.Fatal("Environment variables OPENSKY_USERNAME and OPENSKY_PASSWORD must be set")
	}
	auth := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Add("Authorization", "Basic "+auth)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var flights []Flight
	err = json.Unmarshal(body, &flights)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	return flights
}
