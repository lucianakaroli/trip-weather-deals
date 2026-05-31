package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trip-weather-deals/internal/recommendation"
)

type Handler struct {
	service *recommendation.Service
}

func NewHandler(service *recommendation.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Recommendations(w http.ResponseWriter, r *http.Request) {
	cities := []string{"Recife", "Maceio", "Sao Paulo", "Curitiba"}

	maxPriceParam := r.URL.Query().Get("maxPrice")
	maxPrice := 800.0

	if maxPriceParam != "" {
		if parsed, err := strconv.ParseFloat(maxPriceParam, 64); err == nil {
			maxPrice = parsed
		}
	}

	results, err := h.service.Recommend(cities, maxPrice)
	if err != nil {
		http.Error(w, "failed to generate recommendations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
