package main

import (
	"github.com/AlEx-Truno/task-3/cli"
	"github.com/AlEx-Truno/task-3/config"
	"github.com/AlEx-Truno/task-3/internal/currency"
)

const dirPerm = 0o755

func main() {
	configPath := cli.MustGetConfigPath()

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := currency.LoadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(currencies)

	if err := currency.SaveJSON(cfg.OutputFile, currencies, dirPerm); err != nil {
		panic(err)
	}
}