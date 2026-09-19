package main

import "fmt"

func main() {
	var iNumber1 int
	var iNumber2 int
	var operation string
	fmt.Println("enter first number: ")
	_, err := fmt.Scanln(&iNumber1)
	if err != nil {
		fmt.Println("ERROR: not a number")
		return
	}
	fmt.Println("enter second number: ")
	_, err = fmt.Scanln(&iNumber2)
	if err != nil {
		fmt.Println("ERROR: not a number")
		return
	}
	fmt.Println("choose the operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	if operation == "+" {
		fmt.Println("result: ", iNumber1+iNumber2)
	}
	if operation == "-" {
		fmt.Println("result: ", iNumber1+iNumber2)
	}
	if operation == "*" {
		fmt.Println("result: ", iNumber1*iNumber2)
	}
	if operation == "/" {
		if iNumber2 == 0 {
			fmt.Println("ERROR: divizion by zero")
			return
		}
		fmt.Println("result: ", iNumber1/iNumber2)
	} else {
		fmt.Println("wrong operation")
	}

}
