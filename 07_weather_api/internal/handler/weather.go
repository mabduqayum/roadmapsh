package handler

import (
	"encoding/json"
	"net/http"

	"weather_api/internal/model"
	"weather_api/internal/service"
)

type WeatherHandler struct {
	service *service.WeatherService
}

func NewWeatherHandler(service *service.WeatherService) *WeatherHandler {
	return &WeatherHandler{service: service}
}

func (h *WeatherHandler) GetWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	if city == "" {
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{Error: "city parameter is required"})
		return
	}

	weather, err := h.service.GetWeather(r.Context(), city)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	writeJSON(w, http.StatusOK, weather)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
