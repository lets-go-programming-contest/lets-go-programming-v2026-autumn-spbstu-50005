package main

import (
	"errors"
	"fmt"
)

var (
	ErrDivisionByZero  = errors.New("division by zero")
	ErrInvalidOperator = errors.New("invalid operation")
)

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperator
	}
}

func main() {
	var a, b int
	var op string
	var err error

	if _, err = fmt.Scan(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err = fmt.Scan(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err = fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operator")
		return
	}

	result, err := calculate(a, b, op)
	if err != nil {
		switch {
		case errors.Is(err, ErrDivisionByZero):
			fmt.Println("Division by zero")
		case errors.Is(err, ErrInvalidOperator):
			fmt.Println("Invalid operation")
		}
		return
	}
	fmt.Println(result)
}
