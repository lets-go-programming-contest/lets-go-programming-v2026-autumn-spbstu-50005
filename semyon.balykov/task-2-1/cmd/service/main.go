package main

import (
	"errors"
	"fmt"
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

func (b *temperatureBounds) optimalTemperature() int {
	if b.min > b.max {
		return noTemperature
	}

	return b.min
}

func processDepartment() error {
	var workerCount int

	if _, err := fmt.Scan(&workerCount); err != nil {
		return fmt.Errorf("failed to read workers count: %w", err)
	}

	if workerCount < minCount || workerCount > maxCount {
		return fmt.Errorf(
			"%w: workers count must be between %d and %d",
			errInvalidCount,
			minCount,
			maxCount,
		)
	}

	bounds := newTemperatureBounds()

	for range workerCount {
		var (
			operator    string
			temperature int
		)

		if _, err := fmt.Scan(&operator, &temperature); err != nil {
			return fmt.Errorf("failed to read temperature condition: %w", err)
		}

		if err := bounds.update(operator, temperature); err != nil {
			return err
		}

		if _, err := fmt.Println(bounds.optimalTemperature()); err != nil {
			return fmt.Errorf("failed to print temperature: %w", err)
		}
	}

	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		if _, printErr := fmt.Fprintln(os.Stderr, err); printErr != nil {
			return
		}

		return
	}

	if departmentCount < minCount || departmentCount > maxCount {
		err := fmt.Errorf(
			"%w: departments count must be between %d and %d",
			errInvalidCount,
			minCount,
			maxCount,
		)

		if _, printErr := fmt.Fprintln(os.Stderr, err); printErr != nil {
			return
		}

		return
	}

	for range departmentCount {
		if err := processDepartment(); err != nil {
			if _, printErr := fmt.Fprintln(os.Stderr, err); printErr != nil {
				return
			}

			return
		}
	}
}
