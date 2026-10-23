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
	invalidOperator         = errors.New("invalid operator")
	invalidWorkersCount     = errors.New("invalid workers count")
	invalidDepartmentsCount = errors.New("invalid departments count")
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
		return fmt.Errorf("%w: %q", invalidOperator, operator)
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
		return invalidWorkersCount
	}

	bounds := newTempBounds()

	for i := 0; i < workerCount; i++ {
		var op string
		var temp int

		if _, err := fmt.Scan(&op, &temp); err != nil {
			return fmt.Errorf("failed to read condition: %w", err)
		}

		if err := bounds.update(op, temp); err != nil {
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
		fmt.Fprintln(os.Stderr, invalidDepartmentsCount)
		return
	}

	for i := 0; i < departmentCount; i++ {
		if err := processDepartment(); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			return
		}
	}
}
