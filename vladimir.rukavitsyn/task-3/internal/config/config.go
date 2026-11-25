package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ErrEmptyField = errors.New("input-file and output-file must be set")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (c Config) validate() error {
	if c.InputFile == "" || c.OutputFile == "" {
		return ErrEmptyField
	}

	return nil
}

func Load(path string) (Config, error) {
	var cfg Config

	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return cfg, fmt.Errorf("read config: %w", err)
	}

	if err := yaml.Unmarshal(content, &cfg); err != nil {
		return cfg, fmt.Errorf("parse config: %w", err)
	}

	if err := cfg.validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}