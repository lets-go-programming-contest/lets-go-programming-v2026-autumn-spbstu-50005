package main

import "fmt"

func main() {
	var firstNum, secondNum int
	var operator string
	var err error

	if _, err = fmt.Scan(&firstNum); err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if _, err = fmt.Scan(&secondNum); err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if _, err = fmt.Scan(&operator); err != nil {
		fmt.Println("Invalid operation")
		return
	}

	switch operator {
	case "+":
		fmt.Println(firstNum + secondNum)
	case "-":
		fmt.Println(firstNum - secondNum)
	case "*":
		fmt.Println(firstNum * secondNum)
	case "/":
		if secondNum == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(firstNum / secondNum)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
