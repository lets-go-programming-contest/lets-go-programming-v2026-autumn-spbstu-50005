package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Haist-PM/task-3/internal/filedata"
	"gopkg.in/yaml.v3"
)

var (
	errMissingConfigPath = errors.New("the -config flag requires a path")
	errMissingInputFile  = errors.New("input-file is required")
	errMissingOutputFile = errors.New("output-file is required")
)

type Settings struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func (settings Settings) validate() error {
	if strings.TrimSpace(settings.InputFile) == "" {
		return errMissingInputFile
	}

	if strings.TrimSpace(settings.OutputFile) == "" {
		return errMissingOutputFile
	}

	return nil
}

func Load(path string) (Settings, error) {
	if strings.TrimSpace(path) == "" {
		return Settings{}, errMissingConfigPath
	}

	settings, err := filedata.Read[Settings](path, yaml.Unmarshal)
	if err != nil {
		return Settings{}, fmt.Errorf("read YAML configuration: %w", err)
	}

	if err := settings.validate(); err != nil {
		return Settings{}, fmt.Errorf("validate configuration %q: %w", path, err)
	}

	return settings, nil
}
