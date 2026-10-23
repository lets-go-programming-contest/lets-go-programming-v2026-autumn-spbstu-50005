package main

import "fmt"

func main() {
	var iNumber1, iNumber2 int
	var operation string
	_, err := fmt.Scanln(&iNumber1)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scanln(&iNumber2)
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
		fmt.Println(iNumber1 + iNumber2)
	case "-":
		fmt.Println(iNumber1 - iNumber2)

	case "*":
		fmt.Println(iNumber1 * iNumber2)

	case "/":
		if iNumber2 == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(iNumber1 / iNumber2)
	default:
		fmt.Println("Invalid operation")
	}
}
