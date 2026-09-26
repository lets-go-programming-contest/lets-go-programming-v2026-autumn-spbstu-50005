package main

import "fmt"

func main() {
	var a, b int
	var c string
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	_, err = fmt.Scan(&c)
	if err != nil || (c != "+" && c != "-" && c != "*" && c != "/") {
		fmt.Println("Invalid operation")
		return
	}
	if c == "/" && b == 0 {
		fmt.Println("Division by zero")
		return
	}
	if c == "-" {
		fmt.Println(a - b)
		return
	}
	if c == "*" {
		fmt.Println(a * b)
		return
	}
	if c == "/" {
		fmt.Println(a / b)
		return
	}
	fmt.Println(a + b)
}
