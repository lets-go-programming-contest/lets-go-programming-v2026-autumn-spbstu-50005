package main

import "fmt"

const (
	minTemp       = 15
	maxTemp       = 30
	lessThan      = "<="
	moreThan      = ">="
	noConcessions = -1
)

func ChangeRange(minTemperature, maxTemperature, value int, operator string) (int, int) {
	newMin := minTemperature
	newMax := maxTemperature

	if operator == moreThan && value > newMin {
		newMin = value
	}

	if operator == lessThan && value < newMax {
		newMax = value
	}

	return newMin, newMax
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		return
	}

	for range departmentCount {
		var workerCount int

		if _, err := fmt.Scan(&workerCount); err != nil {
			return
		}

		currentMin := minTemp
		currentMax := maxTemp

		for range workerCount {
			var (
				operator string
				value    int
			)

			if _, err := fmt.Scan(&operator, &value); err != nil {
				return
			}

			currentMin, currentMax = ChangeRange(currentMin, currentMax, value, operator)

			if currentMin <= currentMax {
				fmt.Println(currentMin)
			} else {
				fmt.Println(noConcessions)
			}
		}
	}
}
