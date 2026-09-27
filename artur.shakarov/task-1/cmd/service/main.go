package main

import "fmt"

func main() {
	var firstNum, secNum int
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scan(&firstNum)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		fmt.Println("Invalid operator")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scan(&secNum)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	switch operator {
	case "+":
		fmt.Println(firstNum + secNum)
	case "-":
		fmt.Println(firstNum - secNum)
	case "*":
		fmt.Println(firstNum * secNum)
	case "/":
		if secNum == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(firstNum / secNum)
	default:
		fmt.Println("Invalid operator")
	}
}

