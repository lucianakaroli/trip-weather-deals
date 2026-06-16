package weather

type Forecast struct {
	City      string `json:"city"`
	WillRain  bool   `json:"will_rain"`
	Condition string `json:"condition"`
}

type Client interface {
	GetForecast(city string) (Forecast, error)
}

type MockClient struct{}

func (m MockClient) GetForecast(city string) (Forecast, error) {
	mockData := map[string]Forecast{
		"Recife": {
			City:      "Recife",
			WillRain:  false,
			Condition: "sunny",
		},
		"Maceio": {
			City:      "Maceio",
			WillRain:  false,
			Condition: "partly cloudy",
		},
		"Sao Paulo": {
			City:      "Sao Paulo",
			WillRain:  true,
			Condition: "rainy",
		},
		"Curitiba": {
			City:      "Curitiba",
			WillRain:  true,
			Condition: "storm",
		},
	}

	if forecast, ok := mockData[city]; ok {
		return forecast, nil
	}

	return Forecast{
		City:      city,
		WillRain:  true,
		Condition: "unknown",
	}, nil
}
