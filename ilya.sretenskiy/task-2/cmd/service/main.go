package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

const (
	low        = 15
	high       = 30
	noSolution = -1

	atLeast = ">="
	atMost  = "<="
)

var (
	errUnknownOperator = errors.New("unknown operator")
	errInvalidCount    = errors.New("invalid count")
)

func updateTemperature(minTemp, maxTemp *int, op string, t int) error {
	switch op {
	case atLeast:
		if t > *minTemp {
			*minTemp = t
		}
	case atMost:
		if t < *maxTemp {
			*maxTemp = t
		}
	default:
		return fmt.Errorf("%w: %q", errUnknownOperator, op)
	}
	return nil
}

func readCount(reader *bufio.Reader, name string, max int) (int, error) {
	var count int
	if _, err := fmt.Fscan(reader, &count); err != nil {
		return 0, fmt.Errorf("read %s count: %w", name, err)
	}
	if count <= 0 || count > max {
		return 0, fmt.Errorf("%w: %s count must be between 1 and %d", errInvalidCount, name, max)
	}
	return count, nil
}

func processDepartment(reader *bufio.Reader, K int) error {
	minTemp := low
	maxTemp := high

	for i := 0; i < K; i++ {
		var (
			op string
			t  int
		)

		if _, err := fmt.Fscan(reader, &op, &t); err != nil {
			return fmt.Errorf("read requirement for employee %d: %w", i+1, err)
		}

		if err := updateTemperature(&minTemp, &maxTemp, op, t); err != nil {
			return fmt.Errorf("employee %d: %w", i+1, err)
		}

		if minTemp > maxTemp {
			fmt.Println(noSolution)
		} else {
			fmt.Println(minTemp)
		}
	}

	return nil
}

func run() error {
	reader := bufio.NewReader(os.Stdin)

	departmentCount, err := readCount(reader, "departments", 100)
	if err != nil {
		return err
	}

	for departmentIndex := 0; departmentIndex < departmentCount; departmentIndex++ {
		employeeCount, err := readCount(reader, "employees", 1000)
		if err != nil {
			return fmt.Errorf("department %d: %w", departmentIndex+1, err)
		}

		if err := processDepartment(reader, employeeCount); err != nil {
			return fmt.Errorf("department %d: %w", departmentIndex+1, err)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
