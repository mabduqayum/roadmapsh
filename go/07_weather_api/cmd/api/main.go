package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/time/rate"

	"weather_api/internal/cache"
	"weather_api/internal/config"
	"weather_api/internal/handler"
	"weather_api/internal/service"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Initialize cache
	redisCache, err := cache.NewRedisCache(cfg.RedisURL, cfg.CacheDuration)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize service and handler
	weatherService := service.NewWeatherService(cfg.WeatherAPIKey, redisCache)
	weatherHandler := handler.NewWeatherHandler(weatherService)

	// Setup router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(rateLimitMiddleware(cfg.RateLimit, cfg.RateLimitPeriod))

	// Routes
	r.Get("/weather", weatherHandler.GetWeather)

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}

func rateLimitMiddleware(requests int, period int) func(http.Handler) http.Handler {
	limiter := rate.NewLimiter(rate.Every(time.Duration(period)*time.Second), requests)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !limiter.Allow() {
				http.Error(w, "Too many requests", http.StatusTooManyRequests)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
