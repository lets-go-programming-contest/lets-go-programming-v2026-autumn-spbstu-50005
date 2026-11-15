package main

import (
	"flag"
	"os"
	"strings"

	"github.com/estoiko/task-3/internal/config"
	"github.com/estoiko/task-3/internal/currency"
)

func main() {
	flags := flag.NewFlagSet("service", flag.ContinueOnError)
	configPath := flags.String("config", "", "path to YAML configuration file")

	if err := flags.Parse(os.Args[1:]); err != nil {
		panic(err)
	}

	if strings.TrimSpace(*configPath) == "" {
		panic("required flag -config is missing or empty")
	}

	if flags.NArg() != 0 {
		panic("unexpected positional arguments")
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ReadXML(cfg.InputFilePath)
	if err != nil {
		panic(err)
	}

	currency.SortByValue(currencies)

	err = currency.WriteJSON(cfg.OutputFilePath, currencies)
	if err != nil {
		panic(err)
	}
}
