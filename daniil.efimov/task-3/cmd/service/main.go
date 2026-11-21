package main

import (
	"flag"
	"fmt"

	"github.com/EfimovDD06/task-3/internal/config"
	"github.com/EfimovDD06/task-3/internal/money"
)

func main() {
	configPtr := flag.String("config", "", "path to YAML configuration file")
	flag.Parse()

	if *configPtr == "" {
		panic(fmt.Errorf("flag -config is required"))
	}

	cfg, err := config.LoadAny[config.AppConfig](*configPtr)
	if err != nil {
		panic(err)
	}

	rates, err := money.ExtractRates(cfg.InputFile)
	if err != nil {
		panic(err)
	}

	money.OrderByRateDesc(rates)

	if err := money.SaveAsJSON(cfg.OutputFile, rates); err != nil {
		panic(err)
	}
}
