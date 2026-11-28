package config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

var errMissingPath = errors.New("input-file and output-file must not be empty")

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func Load(path string) (Config, error) {
	var settings Config

	data, err := os.ReadFile(path)
	if err != nil {
		return settings, fmt.Errorf("read config %q: %w", path, err)
	}

	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)

	if err := decoder.Decode(&settings); err != nil {
		return settings, fmt.Errorf("decode config %q: %w", path, err)
	}

	if strings.TrimSpace(settings.InputFile) == "" ||
		strings.TrimSpace(settings.OutputFile) == "" {
		return settings, errMissingPath
	}

	return settings, nil
}
