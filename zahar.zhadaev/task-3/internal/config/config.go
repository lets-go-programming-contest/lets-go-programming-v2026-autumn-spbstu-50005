package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrMissingConfigPath = errors.New("config file path is required")
	ErrEmptyIOPaths      = errors.New("input-file and output-file must be set")
)

type Settings struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadSettings(path string) (Settings, error) {
	trimmedPath := strings.TrimSpace(path)
	if trimmedPath == "" {
		return Settings{}, ErrMissingConfigPath
	}

	raw, err := os.ReadFile(trimmedPath)
	if err != nil {
		return Settings{}, fmt.Errorf("read config %q: %w", trimmedPath, err)
	}

	var cfg Settings
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return Settings{}, fmt.Errorf("unmarshal config %q: %w", trimmedPath, err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return Settings{}, ErrEmptyIOPaths
	}

	return cfg, nil
}
