package main

import "fmt"

func main() {
	var firstNumber, secondNumber int
	var op string
	var err error
	if _, err = fmt.Scan(&firstNumber); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err = fmt.Scan(&secondNumber); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err = fmt.Scan(&op); err != nil {
		fmt.Println("Invalid operation")
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
