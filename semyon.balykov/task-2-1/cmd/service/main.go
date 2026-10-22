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

type temperatureCondition struct {
	operator    string
	temperature int
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
		if temperature >= b.min {
			b.min = temperature
		}

	case lessOrEqual:
		if temperature <= b.max {
			b.max = temperature
		}

	default:
		return fmt.Errorf("%w: %q", errInvalidOperator, operator)
	}

	return nil
}

func (b *temperatureBounds) optimalTemperature() int {
	if b.min > b.max {
		return noTemperature
	}

	return b.min
}

func readCondition() (temperatureCondition, error) {
	var condition temperatureCondition

	if _, err := fmt.Scan(&condition.operator, &condition.temperature); err != nil {
		return temperatureCondition{},
			fmt.Errorf("failed to read temperature condition: %w", err)
	}

	return condition, nil
}

func printTemperature(temperature int) {
	fmt.Println(temperature)
}

func processDepartment() error {
	var workerCount int

	if _, err := fmt.Scan(&workerCount); err != nil {
		return fmt.Errorf("failed to read workers count: %w", err)
	}

	if workerCount <= 0 {
		return fmt.Errorf(
			"%w: workers count must be greater than 0",
			errInvalidCount,
		)
	}

	bounds := newTemperatureBounds()

	for range workerCount {
		condition, err := readCondition()
		if err != nil {
			return err
		}

		if err := bounds.update(
			condition.operator,
			condition.temperature,
		); err != nil {
			return err
		}

		printTemperature(bounds.optimalTemperature())
	}

	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Fprintln(
			os.Stderr,
			"failed to read departments count:",
			err,
		)

		return
	}

	if departmentCount <= 0 {
		fmt.Fprintln(
			os.Stderr,
			"departments count must be greater than 0",
		)

		return
	}

	for range departmentCount {
		if err := processDepartment(); err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}
	}
}
