package main

import (
	"flag"
	"os"

	"github.com/anastasia.kuzmicheva/task-3/internal/config"
	"github.com/anastasia.kuzmicheva/task-3/internal/currency"
)

const (
	dirPerm  os.FileMode = 0o755
	filePerm os.FileMode = 0o644
)

func main() {
	configPath := flag.String("config", "", "path to config file")

	flag.Parse()

	cfg, err := config.Load(*configPath)

	if err != nil {
		panic(err)
	}

	valutes, err := currency.LoadFromXML(cfg.InputFile)

	if err != nil {
		panic(err)
	}

	currency.SortByValueDesc(valutes)

	if err := currency.SaveToJSON(cfg.OutputFile, valutes, dirPerm, filePerm); err != nil {
		panic(err)
	}
}
