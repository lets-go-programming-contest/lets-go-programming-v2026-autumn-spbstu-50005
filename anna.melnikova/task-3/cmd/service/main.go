package main

import (
	"flag"

	"github.com/g0disd3ad/task-3/internal/config"
	"github.com/g0disd3ad/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ParseXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(currencies)

	if err := currency.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
