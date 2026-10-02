package main

import "fmt"

func main() {
	var a, b int
	var sign string

	if _, err := fmt.Scan(&a); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scan(&b); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scan(&sign); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch sign {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}