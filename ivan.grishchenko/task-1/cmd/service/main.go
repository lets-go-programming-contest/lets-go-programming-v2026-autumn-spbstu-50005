package main
import "fmt"
func main() {
	var a, b int
	var op string
	fmt.Scan(&a, &op, &b)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		switch b {
		case 0:
			fmt.Println("Invalid operation")
		default:
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Invalid operation")
	}
}
