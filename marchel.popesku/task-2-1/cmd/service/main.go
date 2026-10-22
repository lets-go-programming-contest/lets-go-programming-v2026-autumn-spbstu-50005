package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minimumTemperature     = 15
	maximumTemperature     = 30
	unavailableTemperature = -1
	maximumCount           = 1000

	atLeast = ">="
	atMost  = "<="
)

var (
	errUnknownOperator = errors.New("unknown operator")
	errInvalidCount    = errors.New("invalid count")
)

type requirement struct {
	operator  string
	threshold int
}

type temperatureInterval struct {
	lower int
	upper int
}

func newTemperatureInterval() *temperatureInterval {
	return &temperatureInterval{
		lower: minimumTemperature,
		upper: maximumTemperature,
	}
}

func (interval *temperatureInterval) apply(operator string, threshold int) error {
	switch operator {
	case atLeast:
		if threshold >= interval.lower {
			interval.lower = threshold
		}
	case atMost:
		if threshold <= interval.upper {
			interval.upper = threshold
		}
	default:
		return fmt.Errorf("%w: %q", errUnknownOperator, operator)
	}

	return nil
}

func (interval *temperatureInterval) optimal() int {
	if interval.lower > interval.upper {
		return unavailableTemperature
	}

	return interval.lower
}

func readCount(name string) (int, error) {
	var count int

	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("read %s count: %w", name, err)
	}

	if count <= 0 || count > maximumCount {
		return 0, fmt.Errorf("%w: %s count must be between 1 and %d", errInvalidCount, name, maximumCount)
	}

	return count, nil
}

func readRequirements() ([]requirement, error) {
	employeeCount, err := readCount("employees")
	if err != nil {
		return nil, err
	}

	requirements := make([]requirement, employeeCount)

	for employeeIndex := range employeeCount {
		var (
			operator  string
			threshold int
		)

		if _, err := fmt.Scan(&operator, &threshold); err != nil {
			return nil, fmt.Errorf("read requirement for employee %d: %w", employeeIndex+1, err)
		}

		requirements[employeeIndex] = requirement{
			operator:  operator,
			threshold: threshold,
		}
	}

	return requirements, nil
}

func calculateTemperatures(requirements []requirement) ([]int, error) {
	interval := newTemperatureInterval()
	temperatures := make([]int, len(requirements))

	for i, current := range requirements {
		if err := interval.apply(current.operator, current.threshold); err != nil {
			return nil, fmt.Errorf("employee %d: %w", i+1, err)
		}

		temperatures[i] = interval.optimal()
	}

	return temperatures, nil
}

func printTemperatures(temperatures []int) {
	for _, temperature := range temperatures {
		fmt.Println(temperature)
	}
}

func run() error {
	departmentCount, err := readCount("departments")
	if err != nil {
		return err
	}

	for department := range departmentCount {
		requirements, err := readRequirements()
		if err != nil {
			return fmt.Errorf("department %d: %w", department+1, err)
		}

		temperatures, err := calculateTemperatures(requirements)
		if err != nil {
			return fmt.Errorf("department %d: %w", department+1, err)
		}

		printTemperatures(temperatures)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
