package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	TrackerDir string `koanf:"tracker_dir"`
	TasksFile  string `koanf:"tasks_file"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	// Load default configuration
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		// If config file doesn't exist, use default values
		k.Set("tracker_dir", ".task-tracker")
		k.Set("tasks_file", "tasks.json")
	}

	// Load environment variables
	if err := k.Load(env.Provider("TASK_", ".", func(s string) string {
		return s
	}), nil); err != nil {
		return nil, err
	}

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
