package main

import (
	"flag"
	"sort"
)

func main() {
	configPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg, err := readConfig(*configPath)
	if err != nil {
		panic(err)
	}

	currencies, err := readValutes(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	if err := writeJSON(cfg.OutputFile, currencies); err != nil {
		panic(err)
	}
}
