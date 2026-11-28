package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

var ErrEmptyFilePath = errors.New("input or output file path is empty")

type Config struct {
	InputFilePath  string `yaml:"input-file"`
	OutputFilePath string `yaml:"output-file"`
}

func (cfg *Config) validate() error {
	cfg.InputFilePath = strings.TrimSpace(cfg.InputFilePath)
	cfg.OutputFilePath = strings.TrimSpace(cfg.OutputFilePath)

	if cfg.InputFilePath == "" || cfg.OutputFilePath == "" {
		return ErrEmptyFilePath
	}

	return nil
}

func Load(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config %q: %w", path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config %q: did not find expected key: %w", path, err)
	}

	if err := cfg.validate(); err != nil {
		return Config{}, fmt.Errorf("validate config %q: %w", path, err)
	}

	return cfg, nil
}
