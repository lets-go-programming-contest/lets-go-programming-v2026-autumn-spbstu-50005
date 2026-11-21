package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Matvei-Babushkin/task-3/internal/config"
	"github.com/Matvei-Babushkin/task-3/internal/currency"
)

const (
	dirPermissions  os.FileMode = 0o755
	filePermissions os.FileMode = 0o600
)

var (
	errMissingConfig  = errors.New("flag -config is required")
	errUnexpectedArgs = errors.New("unexpected positional arguments")
)

func run(args []string) error {
	flags := flag.NewFlagSet("service", flag.ContinueOnError)
	configPath := flags.String("config", "", "path to YAML configuration")

	if err := flags.Parse(args); err != nil {
		return fmt.Errorf("run service: %w", err)
	}

	if strings.TrimSpace(*configPath) == "" {
		return errMissingConfig
	}

	if flags.NArg() != 0 {
		return fmt.Errorf("%w: %v", errUnexpectedArgs, flags.Args())
	}

	cfg, err := config.Load(*configPath)
	if err != nil {
		return fmt.Errorf("run service: %w", err)
	}

	currencies, err := currency.ReadXML(cfg.InputFile)
	if err != nil {
		return fmt.Errorf("run service: %w", err)
	}

	currency.SortByValue(currencies)

	if err := currency.WriteJSON(cfg.OutputFile, currencies, dirPermissions, filePermissions); err != nil {
		return fmt.Errorf("run service: %w", err)
	}

	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		panic(err)
	}
}
