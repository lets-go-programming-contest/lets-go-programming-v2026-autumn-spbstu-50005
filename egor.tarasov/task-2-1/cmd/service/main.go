package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minNum         = 1
	maxNum         = 1000
	minTemperature = 15
	maxTemperature = 30
)

var (
	ErrInvalidData           = errors.New("invalid data")
	ErrCountOutOfRange       = errors.New("number out of range")
	ErrNoPreferenceData      = errors.New("no preference data")
	ErrInvalidTemperature    = errors.New("invalid temperature")
	ErrTemperatureOutOfRange = errors.New("temperature out of range")
	ErrInvalidSign           = errors.New("sign must be >= or <=")
)

func readCount() (int, error) {
	var count int

	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidData, err)
	}

	if count < minNum || count > maxNum {
		return 0, fmt.Errorf("%w: must be from %d to %d", ErrCountOutOfRange, minNum, maxNum)
	}

	return count, nil
}

func readPreference() (string, int, error) {
	var (
		sign        string
		temperature int
	)

	n, err := fmt.Scan(&sign, &temperature)
	switch n {
	case 0:
		return "", 0, fmt.Errorf("%w: %w", ErrNoPreferenceData, err)
	case 1:
		return "", 0, fmt.Errorf("%w: %w", ErrInvalidTemperature, err)
	}

	if temperature < minTemperature || temperature > maxTemperature {
		return "", 0, fmt.Errorf("%w: must be from %d to %d",
			ErrTemperatureOutOfRange,
			minTemperature,
			maxTemperature,
		)
	}

	return sign, temperature, nil
}

func run() error {
	numOfDepartments, err := readCount()
	if err != nil {
		return fmt.Errorf("number of departments: %w", err)
	}

	for i := 0; i < numOfDepartments; i++ {
		numOfEmployees, err := readCount()
		if err != nil {
			return fmt.Errorf("number of employees: %w", err)
		}

		curMinTemperature := minTemperature
		curMaxTemperature := maxTemperature

		for j := 0; j < numOfEmployees; j++ {
			sign, temperature, err := readPreference()
			if err != nil {
				return err
			}

			switch sign {
			case ">=":
				if curMinTemperature < temperature {
					curMinTemperature = temperature
				}
			case "<=":
				if curMaxTemperature > temperature {
					curMaxTemperature = temperature
				}
			default:
				return ErrInvalidSign
			}

			if curMinTemperature > curMaxTemperature {
				fmt.Println(-1)
			} else {
				fmt.Println(curMinTemperature)
			}
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)

		return
	}
}
