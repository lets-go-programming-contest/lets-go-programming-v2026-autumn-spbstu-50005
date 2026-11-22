package main

import (
	"flag"

	"github.com/karam0909/task-3/internal/config"
	"github.com/karam0909/task-3/internal/currency"
	"github.com/karam0909/task-3/internal/storage"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to the YAML configuration file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	rates, err := currency.ParseFile(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(rates)

	err = storage.SaveJSON(cfg.OutputFile, rates)
	if err != nil {
		panic(err)
	}
}
