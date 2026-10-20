package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

const (
	minTemperature = 15
	maxTemperature = 30
	noTemperature  = -1

	minCount = 1
	maxCount = 1000

	lessOrEqual    = "<="
	greaterOrEqual = ">="
)

var (
	errInvalidOperator = errors.New("invalid operator")
	errInvalidCount    = errors.New("invalid count")
)

type temperatureBounds struct {
	min int
	max int
}

func newTemperatureBounds() temperatureBounds {
	return temperatureBounds{
		min: minTemperature,
		max: maxTemperature,
	}
}

func (b *temperatureBounds) update(operator string, temperature int) error {
	switch operator {
	case greaterOrEqual:
		if temperature > b.min {
			b.min = temperature
		}

	case lessOrEqual:
		if temperature < b.max {
			b.max = temperature
		}

	default:
		return fmt.Errorf("%w: %s", errInvalidOperator, operator)
	}

	return nil
}

func (b temperatureBounds) optimalTemperature() int {
	if b.min > b.max {
		return noTemperature
	}

	return b.min
}

func readCount(reader io.Reader, name string) (int, error) {
	var count int

	if _, err := fmt.Fscan(reader, &count); err != nil {
		return 0, fmt.Errorf("failed to read %s count: %w", name, err)
	}

	if count < minCount || count > maxCount {
		return 0, fmt.Errorf(
			"%w: %s count must be between %d and %d",
			errInvalidCount,
			name,
			minCount,
			maxCount,
		)
	}

	return count, nil
}

func processDepartment(reader io.Reader, writer io.Writer) error {
	workerCount, err := readCount(reader, "workers")
	if err != nil {
		return err
	}

	bounds := newTemperatureBounds()

	for range workerCount {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Fscan(reader, &operator, &temperature); err != nil {
			return fmt.Errorf("failed to read temperature condition: %w", err)
		}

		if err := bounds.update(operator, temperature); err != nil {
			return err
		}

		fmt.Fprintln(writer, bounds.optimalTemperature())
	}

	return nil
}

func run(reader io.Reader, writer io.Writer) error {
	departmentCount, err := readCount(reader, "departments")
	if err != nil {
		return err
	}

	for range departmentCount {
		if err := processDepartment(reader, writer); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
