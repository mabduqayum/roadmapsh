package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Storage struct {
		File string `mapstructure:"file"`
	} `mapstructure:"storage"`
}

func Load() *Config {
	v := viper.New()

	// Set the config file name and path
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("config")

	// Read the config file
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshalling config: %v", err)
	}

	return &cfg
}
