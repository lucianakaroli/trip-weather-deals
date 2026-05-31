package recommendation

import (
	"trip-weather-deals/internal/flights"
	"trip-weather-deals/internal/weather"
)

type Result struct {
	City      string  `json:"city"`
	WillRain  bool    `json:"will_rain"`
	Condition string  `json:"condition"`
	Price     float64 `json:"price"`
}

type WeatherClient interface {
	GetForecast(city string) (weather.Forecast, error)
}

type FlightClient interface {
	GetPrice(city string) (flights.Offer, error)
}

type Service struct {
	weatherClient WeatherClient
	flightClient  FlightClient
}

func NewService(weatherClient WeatherClient, flightClient FlightClient) *Service {
	return &Service{
		weatherClient: weatherClient,
		flightClient:  flightClient,
	}
}

func (s *Service) Recommend(cities []string, maxPrice float64) ([]Result, error) {
	var results []Result

	for _, city := range cities {
		forecast, err := s.weatherClient.GetForecast(city)
		if err != nil {
			return nil, err
		}

		offer, err := s.flightClient.GetPrice(city)
		if err != nil {
			return nil, err
		}

		if !forecast.WillRain && offer.Price <= maxPrice {
			results = append(results, Result{
				City:      city,
				WillRain:  forecast.WillRain,
				Condition: forecast.Condition,
				Price:     offer.Price,
			})
		}
	}

	return results, nil
}
