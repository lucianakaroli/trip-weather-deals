package main

import (
	"log"
	"net/http"
	"trip-weather-deals/internal/flights"
	"trip-weather-deals/internal/recommendation"
	"trip-weather-deals/internal/transport"
	"trip-weather-deals/internal/weather"
)

func main() {
	weatherClient := weather.MockClient{}
	flightClient := flights.MockClient{}

	service := recommendation.NewService(weatherClient, flightClient)
	handler := transport.NewHandler(service)

	http.HandleFunc("/recommendations", handler.Recommendations)

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
