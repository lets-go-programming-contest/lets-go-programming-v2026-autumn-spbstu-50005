package main

import (
	"flag"
	"strings"

	"github.com/sbalyk/task-3/internal/config"
	"github.com/sbalyk/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to YAML configuration file")
	flag.Parse()

	if strings.TrimSpace(*configPath) == "" {
		panic("required flag -config is missing")
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

	if err := currency.WriteJSON(cfg.OutputFilePath, currencies); err != nil {
		panic(err)
	}
}
