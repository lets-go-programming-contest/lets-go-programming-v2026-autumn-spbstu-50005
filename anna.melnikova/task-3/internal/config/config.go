package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

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

func (c *Config) Validate() error {
	if c.InputFile == "" || c.OutputFile == "" {
		return ErrInvalidConfig
	}

	return nil
}

func Load(path string) (*Config, error) {
	if path == "" {
		return nil, ErrEmptyConfigPath
	}

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("open config: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	var cfg Config

	decoder := yaml.NewDecoder(file)

	decoder.KnownFields(true)

	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}
