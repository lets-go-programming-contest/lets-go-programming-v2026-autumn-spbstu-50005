package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minTemperature = 15
	maxTemperature = 30
	invalidResult  = -1

	minCount = 1
	maxCount = 1000
)

var (
	errInvalidOperator = errors.New("invalid operator")
	errInvalidCount    = errors.New("invalid count")
)

type preference struct {
	operator    string
	temperature int
}

type temperatureRange struct {
	min int
	max int
}

func newTemperatureRange() *temperatureRange {
	return &temperatureRange{
		min: minTemperature,
		max: maxTemperature,
	}
}

func (r *temperatureRange) update(operator string, temperature int) error {
	switch operator {
	case ">=":
		if temperature >= r.min {
			r.min = temperature
		}

	case "<=":
		if temperature <= r.max {
			r.max = temperature
		}

	default:
		return fmt.Errorf("%w: %s", errInvalidOperator, operator)
	}

	return nil
}

func (r *temperatureRange) optimalTemperature() int {
	if r.min > r.max {
		return invalidResult
	}

	return r.min
}

func readPreferences() ([]preference, error) {
	var workerCount int

	if _, err := fmt.Scan(&workerCount); err != nil {
		return nil, fmt.Errorf("failed to read worker count: %w", err)
	}

	if workerCount < minCount || workerCount > maxCount {
		return nil, fmt.Errorf("%w: %d", errInvalidCount, workerCount)
	}

	preferences := make([]preference, workerCount)

	for i := range preferences {
		if _, err := fmt.Scan(
			&preferences[i].operator,
			&preferences[i].temperature,
		); err != nil {
			return nil, fmt.Errorf("failed to read preference: %w", err)
		}
	}

	return preferences, nil
}

func calculateTemperatures(preferences []preference) ([]int, error) {
	bounds := newTemperatureRange()
	results := make([]int, len(preferences))

	for i, pref := range preferences {
		if err := bounds.update(pref.operator, pref.temperature); err != nil {
			return nil, err
		}

		results[i] = bounds.optimalTemperature()
	}

	return results, nil
}

func printTemperatures(temperatures []int) error {
	for _, temperature := range temperatures {
		if _, err := fmt.Println(temperature); err != nil {
			return fmt.Errorf("failed to print temperature: %w", err)
		}
	}

	return nil
}

func processDepartment() error {
	preferences, err := readPreferences()
	if err != nil {
		return err
	}

	temperatures, err := calculateTemperatures(preferences)
	if err != nil {
		return err
	}

	return printTemperatures(temperatures)
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		log.Print(err)

		return
	}

	if departmentCount < minCount || departmentCount > maxCount {
		log.Print("invalid department count")

		return
	}

	for range departmentCount {
		if err := processDepartment(); err != nil {
			log.Print(err)

			return
		}
	}
}
