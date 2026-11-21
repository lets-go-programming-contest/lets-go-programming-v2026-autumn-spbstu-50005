package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Zahar-Z-V/task-3/internal/config"
	"github.com/Zahar-Z-V/task-3/internal/currency"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	fs := flag.NewFlagSet("currency-converter", flag.ContinueOnError)
	cfgPath := fs.String("config", "config.yaml", "path to configuration YAML")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if fs.NArg() != 0 {
		return fmt.Errorf("unexpected positional arguments: %v", fs.Args())
	}

	cfg, err := config.ReadSettings(*cfgPath)
	if err != nil {
		return err
	}

	rates, err := currency.LoadFromXML(cfg.InputFile)
	if err != nil {
		return err
	}

	currency.OrderByValueDesc(rates)

	if err := currency.ExportJSON(cfg.OutputFile, rates); err != nil {
		return err
	}

	return nil
}
