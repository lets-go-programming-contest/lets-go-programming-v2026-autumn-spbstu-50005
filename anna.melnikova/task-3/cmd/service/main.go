package main

import (
	"flag"
	"os"

	"github.com/g0disd3ad/task-3/internal/config"
	"github.com/g0disd3ad/task-3/internal/currency"
)

const (
	dirPermissions  os.FileMode = 0o755
	filePermissions os.FileMode = 0o644
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg, err := config.Load(configPath)
	if err != nil {
		panic(err)
	}

	curs, err := currency.ParseXML[currency.ValCurs](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortDesc(curs.Valutes)

	if err := currency.SaveJSON(cfg.OutputFile, curs.Valutes, dirPermissions, filePermissions); err != nil {
		panic(err)
	}
}
