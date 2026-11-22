package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var ErrEmptyPath = errors.New("input-file and output-file must be set")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (cfg Config) Validate() error {
	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return ErrEmptyPath
	}

	return nil
}

func Load(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return cfg, fmt.Errorf("read config file: %w", err)
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("decode config file: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return cfg, err
	}

	return cfg, nil
}
