package main

import "fmt"

func main() {
	var (
		first, second int
		operator      string
	)

	n, _ := fmt.Scan(&first, &second, &operator)
	switch n {
	case 0:
		fmt.Println("Invalid first operand")
		return
	case 1:
		fmt.Println("Invalid second operand")
		return
	}

	switch operator {
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
