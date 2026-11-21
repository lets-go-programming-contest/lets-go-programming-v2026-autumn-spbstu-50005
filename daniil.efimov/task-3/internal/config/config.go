package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/goccy/go-yaml"
)

var ErrEmptyFilePaths = errors.New("input or output file path is empty")

type AppConfig struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (c *AppConfig) Validate() error {
	c.InputFile = strings.TrimSpace(c.InputFile)
	c.OutputFile = strings.TrimSpace(c.OutputFile)

	if c.InputFile == "" || c.OutputFile == "" {
		return ErrEmptyFilePaths
	}
	return nil
}

func LoadAny[T any](path string) (T, error) {
	var result T

	raw, err := os.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("read config %q: %w", path, err)
	}

	if err := yaml.Unmarshal(raw, &result); err != nil {
		return result, fmt.Errorf("decode config %q: %w", path, err)
	}

	if v, ok := any(&result).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return result, fmt.Errorf("validate config %q: %w", path, err)
		}
	}

	return result, nil
}
