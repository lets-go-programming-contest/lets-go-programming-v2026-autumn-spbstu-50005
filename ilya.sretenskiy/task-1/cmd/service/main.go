package main

import "fmt"

func main() {
	var x, y int
	var operator string

	if _, err := fmt.Scan(&x); err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	if _, err := fmt.Scan(&y); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	if _, err := fmt.Scan(&operator); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(x + y)
	case "-":
		fmt.Println(x - y)
	case "*":
		fmt.Println(x * y)
	case "/":
		if y == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(x / y)
	default:
		fmt.Println("Invalid operation")
	}
}
