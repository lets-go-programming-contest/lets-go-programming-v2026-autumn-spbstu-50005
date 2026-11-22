package main

import (
	"github.com/GabrielKarpov/task-3/internal/config"
	"github.com/GabrielKarpov/task-3/internal/converter"
	"github.com/GabrielKarpov/task-3/internal/flags"
)

const (
	dirPerm  = 0o755
	filePerm = 0o644
)

func main() {
	configPath := flags.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := converter.ReadXML[converter.Currency](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	converter.SortByValue(currencies)

	err = converter.WriteJSON(cfg.OutputFile, currencies, dirPerm, filePerm)
	if err != nil {
		panic(err)
	}
}
