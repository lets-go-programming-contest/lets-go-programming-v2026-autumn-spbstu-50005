package main

import (
	"flag"
	"fmt"

	"github.com/murderggg/task-3/internal/config"
	"github.com/murderggg/task-3/internal/currency"
)

func run(configPath string) error {
	settings, err := config.Load(configPath)
	if err != nil {
		return fmt.Errorf("load configuration: %w", err)
	}

	currencies, err := currency.Read(settings.InputFile)
	if err != nil {
		return fmt.Errorf("read currencies: %w", err)
	}

	currency.Sort(currencies)

	if err := currency.Write(settings.OutputFile, currencies); err != nil {
		return fmt.Errorf("save currencies: %w", err)
	}

	return nil
}

func main() {
	configPath := flag.String("config", "config.yaml", "path to the YAML configuration")
	flag.Parse()

	if err := run(*configPath); err != nil {
		panic(err)
	}
}
