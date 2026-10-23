package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemp     = 15
	maxTemp     = 30
	invalidTemp = -1

	greaterOrEqual = ">="
	lessOrEqual    = "<="
)

var (
	errInvalidOperator         = errors.New("invalid operator")
	errInvalidWorkersCount     = errors.New("invalid workers count")
	errInvalidDepartmentsCount = errors.New("invalid departments count")
)

type tempBounds struct {
	min int
	max int
}

func newTempBounds() *tempBounds {
	return &tempBounds{
		min: minTemp,
		max: maxTemp,
	}
}

func (tempb *tempBounds) update(operator string, temperature int) error {
	switch operator {
	case greaterOrEqual:
		if temperature >= tempb.min {
			tempb.min = temperature
		}

	case lessOrEqual:
		if temperature <= tempb.max {
			tempb.max = temperature
		}

	default:
		return fmt.Errorf("%w: %q", errInvalidOperator, operator)
	}

	return nil
}

func (tempb *tempBounds) optimalTemperature() int {
	if tempb.min > tempb.max {
		return invalidTemp
	}

	return tempb.min
}

func processDepartment() error {
	var workerCount int

	if _, err := fmt.Scan(&workerCount); err != nil {
		return fmt.Errorf("failed to read workers count: %w", err)
	}

	if workerCount <= 0 {
		return errInvalidWorkersCount
	}

	bounds := newTempBounds()

	for range workerCount {
		var (
			operator string
			temp     int
		)

		if _, err := fmt.Scan(&operator, &temp); err != nil {
			return fmt.Errorf("failed to read condition: %w", err)
		}

		if err := bounds.update(operator, temp); err != nil {
			return err
		}

		fmt.Println(bounds.optimalTemperature())
	}

	return nil
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		fmt.Fprintln(os.Stderr, "failed to read departments count:", err)

		return
	}

	if departmentCount <= 0 {
		fmt.Fprintln(os.Stderr, errInvalidDepartmentsCount)

		return
	}

	for range departmentCount {
		if err := processDepartment(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)

			return
		}
	}
}
