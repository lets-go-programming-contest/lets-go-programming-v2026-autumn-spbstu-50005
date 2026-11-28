package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Haist-PM/task-3/internal/config"
	"github.com/Haist-PM/task-3/internal/filedata"
	"github.com/Haist-PM/task-3/internal/rates"
)

const (
	directoryMode = 0o755
	outputMode    = 0o644
)

func run(args []string) error {
	flags := flag.NewFlagSet("service", flag.ContinueOnError)
	configPath := flags.String("config", "", "path to the YAML configuration")

	if err := flags.Parse(args); err != nil {
		return fmt.Errorf("parse command line: %w", err)
	}

	settings, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("load settings: %w", err)
	}

	currencies, err := rates.Load(settings.InputFile)
	if err != nil {
		return fmt.Errorf("load exchange rates: %w", err)
	}

	rates.SortDescending(currencies)

	if err := filedata.WriteJSON(settings.OutputFile, currencies, directoryMode, outputMode); err != nil {
		return fmt.Errorf("save exchange rates: %w", err)
	}

	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil && !errors.Is(err, flag.ErrHelp) {
		panic(err)
	}
}
