package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstOp, secondOp, op string
	fmt.Scanln(&firstOp)
	fmt.Scanln(&secondOp)
	fmt.Scanln(&op)
	firstNumber, err := strconv.Atoi(firstOp)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	secondNumber, err := strconv.Atoi(secondOp)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	switch op {
	case "+":
		fmt.Println(firstNumber + secondNumber)
	case "-":
		fmt.Println(firstNumber - secondNumber)
	case "*":
		fmt.Println(firstNumber * secondNumber)
	case "/":
		if secondNumber == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstNumber / secondNumber)
	default:
		fmt.Println("Invalid operation")
	}
}
