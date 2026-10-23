package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
)

var errUnknownOperator = errors.New("unknown operator")

type Temperature struct {
	low  int
	high int
}

func NewTemperature() *Temperature {
	return &Temperature{low: minTemperature, high: maxTemperature}
}

func (t *Temperature) Apply(operator string, value int) error {
	switch operator {
	case ">=":
		t.low = max(t.low, value)
	case "<=":
		t.high = min(t.high, value)
	default:
		return fmt.Errorf("%w: %q", errUnknownOperator, operator)
	}

	return nil
}

func (t *Temperature) Optimal() int {
	if t.low > t.high {
		return -1
	}

	return t.low
}

func processDepartment() error {
	var employees int

	if _, err := fmt.Scan(&employees); err != nil {
		return fmt.Errorf("read employees count: %w", err)
	}

	temp := NewTemperature()

	for range employees {
		var (
			operator string
			value    int
		)

		if _, err := fmt.Scan(&operator, &value); err != nil {
			return fmt.Errorf("read employee wish: %w", err)
		}

		if err := temp.Apply(operator, value); err != nil {
			return err
		}

		fmt.Println(temp.Optimal())
	}

	return nil
}

func run() error {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		return fmt.Errorf("read departments count: %w", err)
	}

	for range departments {
		if err := processDepartment(); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
