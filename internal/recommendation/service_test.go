package recommendation

import (
	"testing"
	"trip-weather-deals/internal/flights"
	"trip-weather-deals/internal/weather"
)

type weatherMock struct{}

func (w weatherMock) GetForecast(city string) (weather.Forecast, error) {
	data := map[string]weather.Forecast{
		"Recife":    {City: "Recife", WillRain: false, Condition: "sunny"},
		"Maceio":    {City: "Maceio", WillRain: false, Condition: "cloudy"},
		"Sao Paulo": {City: "Sao Paulo", WillRain: true, Condition: "rainy"},
	}

	return data[city], nil
}

type flightMock struct{}

func (f flightMock) GetPrice(city string) (flights.Offer, error) {
	data := map[string]flights.Offer{
		"Recife":    {City: "Recife", Price: 700},
		"Maceio":    {City: "Maceio", Price: 650},
		"Sao Paulo": {City: "Sao Paulo", Price: 300},
	}

	return data[city], nil
}

func TestRecommend(t *testing.T) {
	service := NewService(weatherMock{}, flightMock{})

	cities := []string{"Recife", "Maceio", "Sao Paulo"}
	maxPrice := 700.0

	results, err := service.Recommend(cities, maxPrice)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("expected 2 recommendations, got %d", len(results))
	}

	foundRecife := false
	foundMaceio := false

	for _, result := range results {
		if result.City == "Recife" {
			foundRecife = true
		}
		if result.City == "Maceio" {
			foundMaceio = true
		}
	}

	if !foundRecife {
		t.Errorf("expected Recife to be recommended")
	}

	if !foundMaceio {
		t.Errorf("expected Maceio to be recommended")
	}
}
