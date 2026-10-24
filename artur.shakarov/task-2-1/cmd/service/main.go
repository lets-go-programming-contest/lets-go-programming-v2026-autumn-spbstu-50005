package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minTemperature = 15
	maxTemperature = 30
	noTemperature  = -1
	maxCount       = 1000
)

var (
	errInvalidCount = errors.New("count must be between 1 and 1000")
	errOperator     = errors.New("operator must be >= or <=")
)

type condition struct {
	operator    string
	temperature int
}

type temperatureRange struct {
	lower int
	upper int
}

func (bounds *temperatureRange) update(item condition) {
	if item.operator == ">=" && item.temperature > bounds.lower {
		bounds.lower = item.temperature
	}

	if item.operator == "<=" && item.temperature < bounds.upper {
		bounds.upper = item.temperature
	}
}

func (bounds *temperatureRange) result() int {
	if bounds.lower > bounds.upper {
		return noTemperature
	}

	return bounds.lower
}

func readCount() (int, error) {
	var count int

	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("read count: %w", err)
	}

	if count < 1 || count > maxCount {
		return 0, errInvalidCount
	}

	return count, nil
}

func readConditions() ([]condition, error) {
	count, err := readCount()
	if err != nil {
		return nil, err
	}

	conditions := make([]condition, count)

	for employeeIndex := range conditions {
		if _, err := fmt.Scan(
			&conditions[employeeIndex].operator,
			&conditions[employeeIndex].temperature,
		); err != nil {
			return nil, fmt.Errorf("read condition: %w", err)
		}

		if conditions[employeeIndex].operator != ">=" &&
			conditions[employeeIndex].operator != "<=" {
			return nil, errOperator
		}
	}

	return conditions, nil
}

func calculate(conditions []condition) []int {
	bounds := &temperatureRange{
		lower: minTemperature,
		upper: maxTemperature,
	}

	answers := make([]int, len(conditions))

	for employeeIndex, item := range conditions {
		bounds.update(item)
		answers[employeeIndex] = bounds.result()
	}

	return answers
}

func printAnswers(answers []int) {
	for _, answer := range answers {
		fmt.Println(answer)
	}
}

func run() error {
	departments, err := readCount()
	if err != nil {
		return err
	}

	for range departments {
		conditions, err := readConditions()
		if err != nil {
			return err
		}

		answers := calculate(conditions)
		printAnswers(answers)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Print(err)
	}
}
