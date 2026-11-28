package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var (
	ErrConfigRead   = errors.New("failed to read config file")
	ErrConfigParse  = errors.New("invalid config file format")
	ErrNoInputFile  = errors.New("input-file is not specified in config")
	ErrNoOutputFile = errors.New("output-file is not specified in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func ReadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%w %q: %w", ErrConfigRead, path, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("%w %q: %w", ErrConfigParse, path, err)
	}

	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" {
		return nil, ErrNoInputFile
	}

	if cfg.OutputFile == "" {
		return nil, ErrNoOutputFile
	}

	return &cfg, nil
}
