package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	GitHubAPIURL string `koanf:"github_api_url"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	// Default configuration
	config := &Config{
		GitHubAPIURL: "https://api.github.com",
	}

	// Try to load configuration from current dir
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err == nil {
		if err := k.Unmarshal("", config); err != nil {
			return nil, err
		}
	}

	return config, nil
}
