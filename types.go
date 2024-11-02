package main

type Flight struct {
	ICAO24       string  `json:"icao24"`
	FirstSeen    int64   `json:"firstSeen"`
	LastSeen     int64   `json:"lastSeen"`
	Callsign     string  `json:"callsign"`
	Departure    string  `json:"estDepartureAirport"`
	Arrival      string  `json:"estArrivalAirport"`
	DepartureLat float64 `json:"estDepartureAirportHorizDistance"`
	ArrivalLat   float64 `json:"estArrivalAirportVertDistance"`
}
