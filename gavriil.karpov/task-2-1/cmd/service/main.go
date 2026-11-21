package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	maximumTemperature = 30
	minimumTemperature = 15
	lessOrEqual        = "<="
	greaterOrEqual     = ">="
)

var ErrInvalidSign = errors.New("invalid sign")

type DepartmentTemperatureRange struct {
	lowestTemperature  int
	highestTemperature int
}

func New() *DepartmentTemperatureRange {
	return &DepartmentTemperatureRange{
		lowestTemperature:  minimumTemperature,
		highestTemperature: maximumTemperature,
	}
}

func (t *DepartmentTemperatureRange) ChangeTemperatures(
	inputTemperature int,
	sign string,
) error {
	switch sign {
	case lessOrEqual:
		if inputTemperature < t.highestTemperature {
			t.highestTemperature = inputTemperature
		}
	case greaterOrEqual:
		if inputTemperature > t.lowestTemperature {
			t.lowestTemperature = inputTemperature
		}
	default:
		return fmt.Errorf("invalid sign %s: %w", sign, ErrInvalidSign)
	}

	return nil
}

func (t *DepartmentTemperatureRange) calculateOptimalTemperature() int {
	if t.highestTemperature < t.lowestTemperature {
		return -1
	}

	return t.lowestTemperature
}

func main() {
	var departmentsNumber int

	if _, err := fmt.Scan(&departmentsNumber); err != nil {
		fmt.Fprintln(os.Stderr, "invalid departments number:", err)

		return
	}

	if departmentsNumber < 1 {
		fmt.Fprintln(os.Stderr, "departments number must be positive")

		return
	}

	for range departmentsNumber {
		var headcount int

		if _, err := fmt.Scan(&headcount); err != nil {
			fmt.Fprintln(os.Stderr, "invalid headcount:", err)

			return
		}

		if headcount < 1 {
			fmt.Fprintln(os.Stderr, "headcount must be positive")

			return
		}

		temperatureRange := New()

		for range headcount {
			var (
				sign        string
				temperature int
			)

			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				fmt.Fprintln(os.Stderr, "invalid Input:", err)

				return
			}

			if err := temperatureRange.ChangeTemperatures(temperature, sign); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)

				return
			}

			fmt.Println(temperatureRange.calculateOptimalTemperature())
		}
	}
}
