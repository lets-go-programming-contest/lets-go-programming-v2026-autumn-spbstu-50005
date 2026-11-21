package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func GetConfig(configPath string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, fmt.Errorf("Error of reading configuration file: %w", err)
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return cfg, fmt.Errorf("Wrong format of yaml: %w", err)
	}

	return cfg, nil
}

func ValidateConfig(cfg *Config) error {
	cfg.InputFile = strings.TrimSpace(cfg.InputFile)
	cfg.OutputFile = strings.TrimSpace(cfg.OutputFile)

	if cfg.InputFile == "" || cfg.OutputFile == "" {
		return errors.New("Wrong config structure")
	}
	return nil
}
