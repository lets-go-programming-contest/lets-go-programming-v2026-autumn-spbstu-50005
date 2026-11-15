package main

import (
	"flag"

	"github.com/sidorenkoanastasia-ca/task-3/internal/config"
	"github.com/sidorenkoanastasia-ca/task-3/internal/currency"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := config.LoadFile(*configPath)
	if err != nil {
		panic(err)
	}

	cur, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	currency.SortByValues(cur)

	if err := currency.WriteJSON(cfg.OutputFile, cur); err != nil {
		panic(err)
	}
}
