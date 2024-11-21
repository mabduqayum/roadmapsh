package config

import (
	"fmt"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	Storage struct {
		TrackerDir string `koanf:"tracker_dir"`
		TasksFile  string `koanf:"tasks_file"`
	} `koanf:"storage"`
}

func Load() (*Config, error) {
	conf := koanf.Conf{
		Delim:       ".",
		StrictMerge: true,
	}
	k := koanf.NewWithConf(conf)
	// Load default configuration
	configPath := "config/config.yaml"
	if err := k.Load(file.Provider(configPath), yaml.Parser()); err != nil {
		// If config file doesn't exist, use default values
		k.Set("storage.tracker_dir", ".db")
		k.Set("storage.tasks_file", "tasks.json")
		fmt.Println("Using default values due to error:", err)
	}

	// Load environment variables
	if err := k.Load(env.Provider("TASK_", ".", func(s string) string {
		return s
	}), nil); err != nil {
		return nil, fmt.Errorf("error loading environment variables: %w", err)
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	// Validate config
	if cfg.Storage.TrackerDir == "" || cfg.Storage.TasksFile == "" {
		return nil, fmt.Errorf("invalid configuration: TrackerDir and TasksFile must be set")
	}

	return &cfg, nil
}
