package main

import (
	"flag"
	"os"
	"sort"

	"github.com/Manicen/task-3/internal/config"
	"github.com/Manicen/task-3/internal/currency"
)

const (
	folderMode os.FileMode = 0o755
	fileMode   os.FileMode = 0o644
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

	if err := currency.WriteJSON(cfg.OutputFile, currencies, folderMode, fileMode); err != nil {
		panic(err)
	}
}
