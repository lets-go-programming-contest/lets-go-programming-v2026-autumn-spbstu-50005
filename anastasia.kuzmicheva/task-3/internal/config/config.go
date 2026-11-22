package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigPathEmpty = errors.New("config path is empty")
	ErrInputRequired   = errors.New("input-file is required")
	ErrOutputRequired  = errors.New("output-file is required")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func validate(cfg Config) error {
	if cfg.InputFile == "" {
		return ErrInputRequired
	}

	if cfg.OutputFile == "" {
		return ErrOutputRequired
	}

	return nil
}

func Load(path string) (Config, error) {
	if path == "" {
		return Config{}, ErrConfigPathEmpty
	}

	data, err := os.ReadFile(path)

	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("parse yaml: %w", err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if err := validate(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
