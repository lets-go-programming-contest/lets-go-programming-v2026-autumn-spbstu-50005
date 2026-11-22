package main

import (
	"flag"
	"sort"

	"github.com/Manicen/task-3/internal/config"
	"github.com/Manicen/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.ReadValutes(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := currency.WriteJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
