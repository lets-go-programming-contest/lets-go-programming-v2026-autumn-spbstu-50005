package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemp            = 15
	maxTemp            = 30
	noValidTemperature = -1
	lessOrEqual        = "<="
	greaterOrEqual     = ">="
)

var ErrInvalidOperator = errors.New("invalid operator")

type bounds struct {
	left, right int
}

func (b *bounds) change(sign string, temp int) error {
	switch sign {
	case lessOrEqual:
		b.right = min(b.right, temp)
	case greaterOrEqual:
		b.left = max(b.left, temp)
	default:
		return fmt.Errorf("%w: %q", ErrInvalidOperator, sign)
	}

	return nil
}

func (b *bounds) optimal() int {
	if b.left > b.right {
		return noValidTemperature
	}

	return b.left
}

func main() {
	var departments int

	if _, err := fmt.Scan(&departments); err != nil {
		fmt.Fprintln(os.Stderr, "Error: invalid departments count:", err)

		return
	}

	for range departments {
		temperatureBounds := bounds{left: minTemp, right: maxTemp}

		var workers int

		if _, err := fmt.Scan(&workers); err != nil {
			fmt.Fprintln(os.Stderr, "Error: invalid workers count:", err)

			return
		}

		for range workers {
			var (
				sign string
				temp int
			)

			if _, err := fmt.Scan(&sign, &temp); err != nil {
				fmt.Fprintln(os.Stderr, "Error: invalid preference:", err)

				return
			}

			if err := temperatureBounds.change(sign, temp); err != nil {
				fmt.Fprintln(os.Stderr, "Error:", err)

				return
			}

			fmt.Println(temperatureBounds.optimal())
		}
	}
}
