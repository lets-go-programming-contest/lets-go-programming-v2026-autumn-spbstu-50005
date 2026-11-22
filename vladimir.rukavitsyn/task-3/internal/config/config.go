package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

var ErrMissingField = errors.New("input-file and output-file must be set")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (*Config, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("open config: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	var cfg Config

	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("decode config: %w", err)
	}

	if strings.TrimSpace(cfg.InputFile) == "" || strings.TrimSpace(cfg.OutputFile) == "" {
		return nil, ErrMissingField
	}

	return &cfg, nil
}