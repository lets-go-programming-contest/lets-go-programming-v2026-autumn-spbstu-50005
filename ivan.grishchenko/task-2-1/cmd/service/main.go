package main

import (
	"fmt"
)

func main() {
	var deptCount int
	var empCount int

	if _, err := fmt.Scan(&deptCount, &empCount); err != nil {
		return
	}

	for range deptCount {
		minTemp := 15
		maxTemp := 30

		for range empCount {
			var operator string
			var temperature int

			if _, err := fmt.Scan(&operator, &temperature); err != nil {
				return
			}

			if operator == ">=" && temperature > minTemp {
				minTemp = temperature
			}

			if operator == "<=" && temperature < maxTemp {
				maxTemp = temperature
			}

			if minTemp <= maxTemp {
				_, _ = fmt.Println(minTemp)
			} else {
				_, _ = fmt.Println(-1)
			}
		}
	}
}
