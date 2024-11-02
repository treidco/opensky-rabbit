package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		invalidArgs()
	} else if os.Args[1] == "producer" {
		producer()
	} else if os.Args[1] == "consumer" {
		consumer()
	} else {
		invalidArgs()
	}

}

func invalidArgs() {
	fmt.Println("Usage: Supply a single command line argument of either 'producer' or 'consumer'")
	return
}

func producer() {
	flights := getData()

	// Interrogate the data
	for _, flight := range flights {
		log.Printf("[x] Sent data for flight %s", flight.Callsign)
		message := fmt.Sprintf("Flight %s from %s to %s", flight.Callsign, flight.Departure, flight.Arrival)
		send(message)
	}
}

func consumer() {
	receive()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
