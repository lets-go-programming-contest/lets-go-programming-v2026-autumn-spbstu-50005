package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	MinAllowedTemp = 15
	MaxAllowedTemp = 30
)

type Employee struct {
	Operand string
	Temp    int
}

type Department struct {
	MinTemp int
	MaxTemp int
}

func NewDepartment() *Department {
	return &Department{
		MinTemp: MinAllowedTemp,
		MaxTemp: MaxAllowedTemp,
	}
}

func readEmployee(in *bufio.Reader) *Employee {
	var emp Employee
	if _, err := fmt.Fscan(in, &emp.Operand, &emp.Temp); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input of employee:", err)
		os.Exit(1)
	}

	return &emp
}

func (d *Department) ProcessEmployee(emp *Employee) int {

	if emp.Temp < MinAllowedTemp || emp.Temp > MaxAllowedTemp {
		fmt.Fprintln(os.Stderr, "Invalid input of employee temp")
		os.Exit(1)
	}

	if emp.Operand == ">=" {
		if emp.Temp > d.MinTemp {
			d.MinTemp = emp.Temp
		}
	} else if emp.Operand == "<=" {
		if emp.Temp < d.MaxTemp {
			d.MaxTemp = emp.Temp
		}
	} else {
		fmt.Fprintln(os.Stderr, "Wrong temperature operand")
		os.Exit(1)
	}

	if d.MinTemp <= d.MaxTemp {
		return d.MinTemp
	}

	return -1
}

func getCountOfEmployees(in *bufio.Reader) {
	var employees int

	if _, err := fmt.Fscan(in, &employees); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input in employees count:", err)
		os.Exit(1)
	}

	if employees <= 0 {
		fmt.Fprintln(os.Stderr, "Error: wrong employees count")
		os.Exit(1)
	}

	dept := NewDepartment()

	for range employees {
		emp := readEmployee(in)
		optimalTemp := dept.ProcessEmployee(emp)
		fmt.Println(optimalTemp)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var departments int

	if _, err := fmt.Fscan(in, &departments); err != nil {
		fmt.Fprintln(os.Stderr, "Invalid input in departments count:", err)
		os.Exit(1)
	}

	if departments <= 0 {
		fmt.Fprintln(os.Stderr, "Error: departments count must be grater than 0")
		os.Exit(1)
	}

	for range departments {
		getCountOfEmployees(in)
	}
}
