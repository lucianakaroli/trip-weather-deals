package flights

type Offer struct {
	City  string  `json:"city"`
	Price float64 `json:"price"`
}

type Client interface {
	GetPrice(city string) (Offer, error)
}

type MockClient struct{}

func (m MockClient) GetPrice(city string) (Offer, error) {
	mockData := map[string]Offer{
		"Recife": {
			City:  "Recife",
			Price: 750.00,
		},
		"Maceio": {
			City:  "Maceio",
			Price: 680.00,
		},
		"Sao Paulo": {
			City:  "Sao Paulo",
			Price: 400.00,
		},
		"Curitiba": {
			City:  "Curitiba",
			Price: 550.00,
		},
	}

	if offer, ok := mockData[city]; ok {
		return offer, nil
	}

	return Offer{
		City:  city,
		Price: 9999.99,
	}, nil
}
