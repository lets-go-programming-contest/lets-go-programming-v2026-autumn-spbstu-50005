package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyConfigPath = errors.New("config path is empty")
	ErrInvalidConfig   = errors.New("invalid config signature")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		return nil, ErrEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var cfg Config

	decoder := yaml.NewDecoder(bytes.NewReader(data))

	decoder.KnownFields(true)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return nil, ErrInvalidConfig
	}

	return &cfg, nil
}
