package main

import "fmt"

func main() {
	var (
		first, second int
		op string
		err error
	)

	if _, err = fmt.Scan(&first); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err = fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}
	if _, err = fmt.Scan(&second); err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	switch op {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
