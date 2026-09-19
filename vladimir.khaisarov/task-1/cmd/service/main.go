package main

import "fmt"

func main() {
	var iNumber1 int
	var iNumber2 int
	var operation string
	fmt.Println("enter first number: ")
	_,err := fmt.Scanln(&iNumber1)
	if err != nil {
		fmt.Println("ERROR: not a number")
		return
	}
	fmt.Println("enter second number: ")
	_,err = fmt.Scanln(&iNumber2)
	if err != nil {
		fmt.Println("ERROR: not a number")
		return
	}
	fmt.Println("choose the operation (+, -, *, /): ")
	fmt.Scanln(&operation)
	_,err = fmt.Scanln(&operation)
	if err != nil {
		fmt.Println("ERROR: invalid operation")
		return
	}

	switch operation {
	case "+":
		fmt.Println("result: ", iNumber1+iNumber2)

	case "-":
		fmt.Println("result: ", iNumber1-iNumber2)

	case "*":
		fmt.Println("result: ", iNumber1*iNumber2)

	case "/":
		if iNumber2 == 0 {
			fmt.Println("ERROR: divizion by zero")
			return
		}

		fmt.Println("result: ", iNumber1/iNumber2)
	default:
		fmt.Println("wrong operation")
	}
}