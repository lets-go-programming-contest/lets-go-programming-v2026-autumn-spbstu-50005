package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	ErrEmptyInputFile  = errors.New("input-file is empty")
	ErrEmptyOutputFile = errors.New("output-file is empty")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("parse config file: %w", err)
	}

	if cfg.InputFile == "" {
		return nil, ErrEmptyInputFile
	}

	if cfg.OutputFile == "" {
		return nil, ErrEmptyOutputFile
	}

	return &cfg, nil
}
