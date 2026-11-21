package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/magnatnapalochke/task-3/internal/config"
	"github.com/magnatnapalochke/task-3/internal/currency"
	"github.com/magnatnapalochke/task-3/internal/parser"
)

const (
	dirPerm  os.FileMode = 0o0755
	filePerm os.FileMode = 0o0644
)

func main() {
	configPath := flag.String("config", "config.yaml", "path to config file")
	flag.Parse()

	cfg, err := config.GetConfig(*configPath)
	if err != nil {
		panic(err)
	}

	if err := config.ValidateConfig(&cfg); err != nil {
		panic(err)
	}

	curs, err := parser.ReadXML[currency.ValCurs](cfg.InputFile)
	if err != nil {
		panic(err)
	}

	curs.SortValueByDesc()

	if err := parser.EnsureDir(cfg.OutputFile, dirPerm); err != nil {
		panic(err)
	}

	if err := parser.WriteJSON(cfg.OutputFile, curs.Valutes, filePerm); err != nil {
		panic(err)
	}

	fmt.Println("Data saved in:", cfg.OutputFile)
}
