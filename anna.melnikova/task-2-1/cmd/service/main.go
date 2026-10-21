package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
	noSolution     = -1

	opGreaterOrEqual = ">="
	opLessOrEqual    = "<="
)

var (
	ErrUnexpectedEOF      = errors.New("unexpected end of input")
	ErrInvalidInteger     = errors.New("invalid integer")
	ErrInvalidCondition   = errors.New("invalid condition format")
	ErrUnknownOperator    = errors.New("unknown operator")
	ErrInvalidTemperature = errors.New("invalid temperature value")
)

type DepartmentTracker struct {
	minTemp int
	maxTemp int
}

func NewDepartmentTracker() *DepartmentTracker {
	return &DepartmentTracker{
		minTemp: defaultMinTemp,
		maxTemp: defaultMaxTemp,
	}
}

func (d *DepartmentTracker) Apply(operator string, temp int) (int, error) {
	switch operator {
	case opGreaterOrEqual:
		if temp > d.minTemp {
			d.minTemp = temp
		}
	case opLessOrEqual:
		if temp < d.maxTemp {
			d.maxTemp = temp
		}
	default:
		return 0, fmt.Errorf("%w: %q", ErrUnknownOperator, operator)
	}

	if d.minTemp <= d.maxTemp {
		return d.minTemp, nil
	}

	return noSolution, nil
}

func nextWord(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("scanner error: %w", err)
		}

		return "", ErrUnexpectedEOF
	}

	return scanner.Text(), nil
}

func parseCondition(scanner *bufio.Scanner) (string, int, error) {
	token, err := nextWord(scanner)
	if err != nil {
		return "", 0, fmt.Errorf("%w: %w", ErrInvalidCondition, err)
	}

	var (
		operator string
		numStr   string
	)

	switch {
	case strings.HasPrefix(token, opGreaterOrEqual):
		operator = opGreaterOrEqual
		numStr = strings.TrimPrefix(token, opGreaterOrEqual)
	case strings.HasPrefix(token, opLessOrEqual):
		operator = opLessOrEqual
		numStr = strings.TrimPrefix(token, opLessOrEqual)
	default:
		return "", 0, fmt.Errorf("%w: %q", ErrInvalidCondition, token)
	}

	numStr = strings.TrimSpace(numStr)
	if numStr == "" {
		numStr, err = nextWord(scanner)
		if err != nil {
			return "", 0, fmt.Errorf("%w: expected number after %s: %w", ErrInvalidCondition, operator, err)
		}
	}

	temp, err := strconv.Atoi(numStr)
	if err != nil {
		return "", 0, fmt.Errorf("%w: %q: %w", ErrInvalidTemperature, numStr, err)
	}

	return operator, temp, nil
}

func processDepartment(scanner *bufio.Scanner, writer *bufio.Writer, deptIdx int) error {
	tokenK, err := nextWord(scanner)
	if err != nil {
		return fmt.Errorf("department #%d: failed to read K: %w", deptIdx+1, err)
	}

	numEmployees, err := strconv.Atoi(tokenK)
	if err != nil || numEmployees < 0 {
		return fmt.Errorf("%w: K=%q: %w", ErrInvalidInteger, tokenK, err)
	}

	tracker := NewDepartmentTracker()

	for range numEmployees {
		operator, temp, err := parseCondition(scanner)
		if err != nil {
			return fmt.Errorf("department #%d: %w", deptIdx+1, err)
		}

		optimalTemp, err := tracker.Apply(operator, temp)
		if err != nil {
			return err
		}

		if _, err := writer.WriteString(strconv.Itoa(optimalTemp)); err != nil {
			return fmt.Errorf("write error: %w", err)
		}

		if err := writer.WriteByte('\n'); err != nil {
			return fmt.Errorf("write error: %w", err)
		}
	}

	return nil
}

func run(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(w)

	tokenN, err := nextWord(scanner)
	if err != nil {
		if errors.Is(err, ErrUnexpectedEOF) {
			return nil
		}

		return fmt.Errorf("failed to read N: %w", err)
	}

	numDepartments, err := strconv.Atoi(tokenN)
	if err != nil || numDepartments < 0 {
		return fmt.Errorf("%w: N=%q: %w", ErrInvalidInteger, tokenN, err)
	}

	for i := range numDepartments {
		if err := processDepartment(scanner, writer, i); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return fmt.Errorf("flush error: %w", err)
	}

	return nil
}

func main() {
	if err := run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
