package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	InputFile  string `yaml:"input-file"`
	OutputFile string `yaml:"output-file"`
}

func MustLoad(path string) Config {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(fmt.Errorf("Error: config failed to read: %w", err))
	}

	var cfg Config
	if err = yaml.Unmarshal(data, &cfg); err != nil {
		panic(fmt.Errorf("Error: config failed to parse: %w", err))
	}

	return cfg
}
