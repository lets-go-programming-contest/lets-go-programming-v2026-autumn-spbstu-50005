package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemp        = 15
	maxTemp        = 30
	noTemp         = -1
	lessOrEqual    = "<="
	greaterOrEqual = ">="
)

var (
	ErrUnknownOperator = errors.New("unknown comparison operator")
	ErrTempOutOfRange  = errors.New("temperature is out of allowed range")
	ErrNoSuitableTemp  = errors.New("no suitable temperature found")
	ErrInvalidValue    = errors.New("value must be greater than zero")
)

type tempRange struct {
	left, right int
}

func (tempRng *tempRange) update(sign string, temp int) error {
	if temp < minTemp || temp > maxTemp {
		return fmt.Errorf("%w: %d (allowed %d..%d)", ErrTempOutOfRange, temp, minTemp, maxTemp)
	}

	switch sign {
	case lessOrEqual:
		tempRng.right = min(tempRng.right, temp)
	case greaterOrEqual:
		tempRng.left = max(tempRng.left, temp)
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperator, sign)
	}

	return nil
}

func (tempRng *tempRange) findOptimalTemp() (int, error) {
	if tempRng.left > tempRng.right {
		return 0, ErrNoSuitableTemp
	}

	return tempRng.left, nil
}

func processDepartment(workers int) error {
	var tempRng tempRange
	tempRng.left = minTemp
	tempRng.right = maxTemp

	for range workers {
		var (
			sign string
			temp int
		)

		if _, err := fmt.Scan(&sign, &temp); err != nil {
			return fmt.Errorf("failed to read employee preference: %w", err)
		}

		if err := tempRng.update(sign, temp); err != nil {
			return err
		}

		if opt, err := tempRng.findOptimalTemp(); err != nil {
			fmt.Println(noTemp)
		} else {
			fmt.Println(opt)
		}
	}

	return nil
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "Error: failed to read departments count")

		return
	}

	if departments <= 0 {
		fmt.Fprintln(os.Stderr, "Error:", ErrInvalidValue)

		return
	}

	for range departments {
		var workers int
		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Fprintln(os.Stderr, "Error: failed to read workers count")

			return
		}

		if workers <= 0 {
			fmt.Fprintln(os.Stderr, "Error:", ErrInvalidValue)

			return
		}

		if err := processDepartment(workers); err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)

			return
		}
	}
}
