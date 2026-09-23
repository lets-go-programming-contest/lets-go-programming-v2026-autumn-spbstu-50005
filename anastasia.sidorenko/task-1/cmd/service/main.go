package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstOp, secondOp, op string
	if _, err := fmt.Scanln(&firstOp); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err := fmt.Scanln(&secondOp); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err := fmt.Scanln(&op); err != nil {
		fmt.Println("Invalid operation")
		return
	}
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
