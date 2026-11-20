package main

import (
	"flag"

	"github.com/tarasovea/task-3/internal/config"
	"github.com/tarasovea/task-3/internal/currency"
	"github.com/tarasovea/task-3/internal/storage"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	err = storage.Save(cfg.OutputFile, currencies)
	if err != nil {
		panic(err)
	}
}
