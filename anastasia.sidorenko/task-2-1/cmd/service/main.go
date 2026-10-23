package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	minTemperature = 15
	maxTemperature = 30
	expectedParts  = 2
	invalidTemp    = -1

	operandGreaterOrEqual = ">="
	operandLessOrEqual    = "<="
)

var (
	ErrUnexpectedEOF      = errors.New("unexpected end of input")
	ErrInvalidPreference  = errors.New("invalid preference format")
	ErrUnknownOperator    = errors.New("unknown operator")
	ErrInvalidInteger     = errors.New("invalid integer")
	ErrInvalidTemperature = errors.New("invalid temperature value")
	ErrOutOfRange         = errors.New("temperature out of range")
)

type Preference struct {
	Operand string
	Degree  int
}

type TemperatureRange struct {
	low  int
	high int
}

func NewTemperatureRange() *TemperatureRange {
	return &TemperatureRange{
		low:  minTemperature,
		high: maxTemperature,
	}
}

func (tr *TemperatureRange) Update(operand string, degree int) error {
	if degree < minTemperature || degree > maxTemperature {
		return fmt.Errorf("%w: %d", ErrOutOfRange, degree)
	}

	switch operand {
	case operandGreaterOrEqual:
		if degree > tr.low {
			tr.low = degree
		}
	case operandLessOrEqual:
		if degree < tr.high {
			tr.high = degree
		}
	default:
		return fmt.Errorf("%w: %q", ErrUnknownOperator, operand)
	}

	return nil
}

func (tr *TemperatureRange) GetOptimal() int {
	if tr.low > tr.high {
		return invalidTemp
	}

	return tr.low
}

func readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("failed to read line: %w", err)
		}

		return "", ErrUnexpectedEOF
	}

	return strings.TrimSpace(scanner.Text()), nil
}

func readInt(scanner *bufio.Scanner) (int, error) {
	line, err := readLine(scanner)
	if err != nil {
		return 0, err
	}

	integer, err := strconv.Atoi(line)
	if err != nil {
		return 0, fmt.Errorf("%w: %q: %w", ErrInvalidInteger, line, err)
	}

	return integer, nil
}

func parsePreference(line string) (Preference, error) {
	parts := strings.Fields(line)

	if len(parts) != expectedParts {
		return Preference{}, fmt.Errorf("%w: %q", ErrInvalidPreference, line)
	}

	operand := parts[0]

	if operand != operandGreaterOrEqual && operand != operandLessOrEqual {
		return Preference{}, fmt.Errorf("%w: %q", ErrUnknownOperator, operand)
	}

	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return Preference{}, fmt.Errorf("%w: %q: %w", ErrInvalidTemperature, parts[1], err)
	}

	return Preference{Operand: operand, Degree: val}, nil
}

func processDepartment(scanner *bufio.Scanner) error {
	tempRange := NewTemperatureRange()

	workers, err := readInt(scanner)
	if err != nil {
		return err
	}

	for range workers {
		line, err := readLine(scanner)
		if err != nil {
			return err
		}

		pref, err := parsePreference(line)
		if err != nil {
			return err
		}

		if err := tempRange.Update(pref.Operand, pref.Degree); err != nil {
			return err
		}

		fmt.Println(tempRange.GetOptimal())
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	departments, err := readInt(scanner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return
	}

	for range departments {
		if err := processDepartment(scanner); err != nil {
			fmt.Fprintln(os.Stderr, err)

			return
		}
	}
}
