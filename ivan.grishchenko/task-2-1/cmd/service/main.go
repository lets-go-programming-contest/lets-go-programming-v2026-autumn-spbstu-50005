package main

import (
	"fmt"
)

func main() {
	var departmentCount, employeeCount int
	if _, err := fmt.Scan(&departmentCount, &employeeCount); err != nil {
		return
	}

	for departmentIdx := 0; departmentIdx < departmentCount; departmentIdx++ {
		minTemp := 15
		maxTemp := 30

		for employeeIdx := 0; employeeIdx < employeeCount; employeeIdx++ {
			var operator string
			var temperature int
			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				return
			}

			if operator == ">=" && temperature > minTemp {
				minTemp = temperature
			} else if operator == "<=" && temperature < maxTemp {
				maxTemp = temperature
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
