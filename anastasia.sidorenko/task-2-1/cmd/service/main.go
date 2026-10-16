package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Preference struct {
	Operand string
	Degree  int
}

func readLine(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("Error: failed to read line: %w", err)
		}
		return "", fmt.Errorf("Error: unexpected end of input")
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
		return 0, fmt.Errorf("Error: invalid integer %q: %w", line, err)
	}
	return integer, nil
}

func parsePreference(line string) (Preference, error) {
	parts := strings.Fields(line)
	if len(parts) != 2 {
		return Preference{}, fmt.Errorf("Error: invalid preference format %q", line)
	}
	op := parts[0]
	if op != "<=" && op != ">=" {
		return Preference{}, fmt.Errorf("Error: unknown operator %q", op)
	}
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		return Preference{}, fmt.Errorf("Error: invalid temperature value %q: %w", parts[1], err)
	}
	return Preference{Operand: op, Degree: val}, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var err error
	var departments int
	var workers int
	var line string
	var pref Preference
	var low, high int

	departments, err = readInt(scanner)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for i := 0; i < departments; i++ {
		low, high = 15, 30

		workers, err = readInt(scanner)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		for j := 0; j < workers; j++ {
			line, err = readLine(scanner)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}

			pref, err = parsePreference(line)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
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
	}
}
