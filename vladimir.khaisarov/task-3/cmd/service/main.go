package main

import (
	"flag"

	"github.com/xayse/task-3/internal/config"
	"github.com/xayse/task-3/internal/currency"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to YAML config file")
	flag.Parse()

	cfg, err := config.Load(*configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := currency.ReadXML[currency.ValCurs](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(valCurs.Currencies)

	if err := currency.WriteJSON(cfg.OutputFile, valCurs.Currencies, dirPermissions, filePermissions); err != nil {
		panic(err)
	}
}
