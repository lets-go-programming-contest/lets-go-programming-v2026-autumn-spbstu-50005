package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var departmentCount, employeeCount int

	if _, err := fmt.Fscan(reader, &departmentCount, &employeeCount); err != nil {
		return
	}

	for departmentIdx := 0; departmentIdx < departmentCount; departmentIdx++ {
		minTemp := 15
		maxTemp := 30

		for employeeIdx := 0; employeeIdx < employeeCount; employeeIdx++ {
			var operator string
			var temperature int

			if _, err := fmt.Fscan(reader, &operator, &temperature); err != nil {
				return
			}

			if operator == ">=" && temperature > minTemp {
				minTemp = temperature
			} else if operator == "<=" && temperature < maxTemp {
				maxTemp = temperature
			}

			if minTemp <= maxTemp {
				fmt.Fprintln(writer, minTemp)
			} else {
				fmt.Fprintln(writer, -1)
			}
		}
	}
}
