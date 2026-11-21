package main

import (
	"github.com/AlEx-Truno/task-3/cli"
	"github.com/AlEx-Truno/task-3/config"
	"github.com/AlEx-Truno/task-3/internal/currency"
)

func main() {
	configPath := cli.MustGetConfigPath()
	cfg := config.MustLoad(configPath)

	currencies, err := currency.LoadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(currencies)

	if err := currency.SaveJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
