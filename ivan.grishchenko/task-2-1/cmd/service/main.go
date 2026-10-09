package main

import (
	"fmt"
)

const (
	defaultMinTemp = 15
	defaultMaxTemp = 30
	notFoundValue  = -1
	opGte          = ">="
	opLte          = "<="
)

func updateRange(currMin, currMax, val int, operator string) (int, int) {
	newMin, newMax := currMin, currMax

	if operator == opGte && val > newMin {
		newMin = val
	}

	if operator == opLte && val < newMax {
		newMax = val
	}

	return newMin, newMax
}

func main() {
	var totalDepts int

	if _, err := fmt.Scan(&totalDepts); err != nil {
		fmt.Println("Error reading departments:", err)

		return
	}

	for range totalDepts {
		var totalEmps int

		if _, err := fmt.Scan(&totalEmps); err != nil {
			fmt.Println("Error reading employees:", err)

			return
		}

		minTemp := defaultMinTemp
		maxTemp := defaultMaxTemp

		for range totalEmps {
			var (
				cond string
				val  int
			)

			if _, err := fmt.Scan(&cond, &val); err != nil {
				fmt.Println("Error reading condition:", err)

				return
			}

			minTemp, maxTemp = updateRange(minTemp, maxTemp, val, cond)

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(notFoundValue)
			}
		}
	}
}
