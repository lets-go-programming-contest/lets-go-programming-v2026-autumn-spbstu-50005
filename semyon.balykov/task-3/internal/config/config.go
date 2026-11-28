package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyConfigPath = errors.New("config path is empty")
	ErrInvalidConfig   = errors.New("invalid config")
)

type Config struct {
	InputFilePath  string `yaml:"input-file"`
	OutputFilePath string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	if strings.TrimSpace(path) == "" {
		return Config{}, ErrEmptyConfigPath
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg Config

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	if err := decoder.Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("%w: %w", ErrInvalidConfig, err)
	}

	cfg.InputFilePath = strings.TrimSpace(cfg.InputFilePath)
	cfg.OutputFilePath = strings.TrimSpace(cfg.OutputFilePath)

	if cfg.InputFilePath == "" || cfg.OutputFilePath == "" {
		return Config{}, ErrInvalidConfig
	}

	return cfg, nil
}
