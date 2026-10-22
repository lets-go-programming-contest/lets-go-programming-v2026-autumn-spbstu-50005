package main

import (
	"fmt"
	"os"
)

const (
	minNum         = 1
	maxNum         = 1000
	minTemperature = 15
	maxTemperature = 30
)

func readCount() (int, error) {
	var count int
	_, err := fmt.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("invalid data: %v", err)
	}
	if count < minNum || count > maxNum {
		return 0,
			fmt.Errorf("number must be from %v to %v",
				minNum,
				maxNum,
			)
	}
	return count, nil
}

func readPreference() (string, int, error) {
	var sign string
	var temperature int
	n, err := fmt.Scan(&sign, &temperature)
	switch n {
	case 0:
		return "", 0, fmt.Errorf("no preference data: %v", err)
	case 1:
		return "", 0, fmt.Errorf("invalid temperature: %v", err)
	}

	if temperature < minTemperature || temperature > maxTemperature {
		return "", 0, fmt.Errorf("temperature must be from %d to %d",
			minTemperature,
			maxTemperature,
		)
	}
	return sign, temperature, nil
}

func run() error {
	numOfDepartments, err := readCount()
	if err != nil {
		return fmt.Errorf("number of departments: %v", err)
	}
	for i := 0; i < numOfDepartments; i++ {
		numOfEmployees, err := readCount()
		if err != nil {
			return fmt.Errorf("number of employees: %v", err)
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
				return fmt.Errorf("sign must be >= or <=")
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
