package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port            string `mapstructure:"PORT"`
	WeatherAPIKey   string `mapstructure:"WEATHER_API_KEY"`
	RedisURL        string `mapstructure:"REDIS_URL"`
	CacheDuration   int    `mapstructure:"CACHE_DURATION"`
	RateLimit       int    `mapstructure:"RATE_LIMIT"`
	RateLimitPeriod int    `mapstructure:"RATE_LIMIT_PERIOD"`
}

func LoadConfig() (Config, error) {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}
