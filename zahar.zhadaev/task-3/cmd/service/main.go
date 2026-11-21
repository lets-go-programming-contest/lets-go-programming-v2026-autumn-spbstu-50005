package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Zahar-Z-V/task-3/internal/config"
	"github.com/Zahar-Z-V/task-3/internal/currency"
)

var errUnexpectedArgs = errors.New("unexpected positional arguments")

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	flagSet := flag.NewFlagSet("currency-converter", flag.ContinueOnError)
	cfgPath := flagSet.String("config", "config.yaml", "path to configuration YAML")

	if err := flagSet.Parse(args); err != nil {
		return fmt.Errorf("parse flags: %w", err)
	}

	if flagSet.NArg() != 0 {
		return fmt.Errorf("%w: %v", errUnexpectedArgs, flagSet.Args())
	}

	cfg, err := config.ReadSettings(*cfgPath)
	if err != nil {
		return fmt.Errorf("read settings: %w", err)
	}

	rates, err := currency.LoadFromXML(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("load rates from xml: %w", err)
	}

	currency.OrderByValueDesc(rates)

	if err := currency.ExportJSON(cfg.OutputFile, rates); err != nil {
		return fmt.Errorf("export json: %w", err)
	}

	return nil
}
