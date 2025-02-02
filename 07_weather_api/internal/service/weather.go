package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"weather_api/internal/cache"
	"weather_api/internal/model"
)

type WeatherService struct {
	apiKey string
	cache  *cache.RedisCache
}

func NewWeatherService(apiKey string, cache *cache.RedisCache) *WeatherService {
	return &WeatherService{
		apiKey: apiKey,
		cache:  cache,
	}
}

func (s *WeatherService) GetWeather(ctx context.Context, city string) (*model.WeatherData, error) {
	// Try cache first
	if cached, err := s.cache.Get(ctx, city); err == nil && cached != nil {
		return cached, nil
	}

	// Call external API
	url := fmt.Sprintf("https://weather.visualcrossing.com/VisualCrossingWebServices/rest/services/timeline/%s?unitGroup=metric&key=%s", city, s.apiKey)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather API returned status code: %d", resp.StatusCode)
	}

	// Parse response and create WeatherData
	var result model.WeatherData
	// Parse the actual response and populate result
	// This will depend on the actual API response structure

	// Cache the result
	if err := s.cache.Set(ctx, city, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
