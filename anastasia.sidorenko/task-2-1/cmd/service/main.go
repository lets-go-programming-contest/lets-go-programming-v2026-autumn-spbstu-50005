package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const expectedPartsCount = 2

var (
	ErrUnexpectedEOF      = errors.New("unexpected end of input")
	ErrInvalidPreference  = errors.New("invalid preference format")
	ErrUnknownOperator    = errors.New("unknown operator")
	ErrInvalidInteger     = errors.New("invalid integer")
	ErrInvalidTemperature = errors.New("invalid temperature value")
)

type Preference struct {
	Operand string
	Degree  int
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
		return 0, fmt.Errorf("%w: %q", ErrInvalidInteger, line)
	}

	return integer, nil
}

func parsePreference(line string) (Preference, error) {
	parts := strings.Fields(line)

	if len(parts) != expectedPartsCount {
		return Preference{}, fmt.Errorf("%w: %q", ErrInvalidPreference, line)
	}

	operand := parts[0]

	if operand != "<=" && operand != ">=" {
		return Preference{}, fmt.Errorf("%w: %q", ErrUnknownOperator, operand)
	}

	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return Preference{}, fmt.Errorf("%w: %q", ErrInvalidTemperature, parts[1])
	}

	return Preference{Operand: operand, Degree: val}, nil
}

func processDepartment(scanner *bufio.Scanner) error {
	low, high := 15, 30

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

		switch pref.Operand {
		case "<=":
			if pref.Degree < high {
				high = pref.Degree
			}
		case ">=":
			if pref.Degree > low {
				low = pref.Degree
			}
		}

		if low > high {
			fmt.Println(-1)
		} else {
			fmt.Println(low)
		}
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
