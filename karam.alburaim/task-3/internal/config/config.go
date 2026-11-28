package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	ErrNoInputFile  = errors.New("input-file is not set in config")
	ErrNoOutputFile = errors.New("output-file is not set in config")
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	var cfg Config

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return cfg, fmt.Errorf("cannot read config %q: %w", path, err)
	}

	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, fmt.Errorf("cannot decode config %q: %w", path, err)
	}

	switch {
	case cfg.InputFile == "":
		return cfg, ErrNoInputFile
	case cfg.OutputFile == "":
		return cfg, ErrNoOutputFile
	}

	return cfg, nil
}
