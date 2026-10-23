package main

import "fmt"

func main() {
	var firstOp, secondOp int
	var operation string

	_, err := fmt.Scanln(&firstOp)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondOp)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println(firstOp + secondOp)
	case "-":
		fmt.Println(firstOp - secondOp)
	case "*":
		fmt.Println(firstOp * secondOp)
	case "/":
		if secondOp == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstOp / secondOp)
	default:
		fmt.Println("Invalid operation")
	}
}
