package main

import (
	"flag"

	"github.com/RukavitsynVladimir/task-3/internal/config"
	"github.com/RukavitsynVladimir/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	items, err := currency.Load(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(items)

	if err := currency.Save(cfg.OutputFile, items); err != nil {
		panic(err)
	}
}