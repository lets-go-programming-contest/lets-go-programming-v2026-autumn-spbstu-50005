package main

import "fmt"

const (
	minTemp = 15
	maxTemp = 30
)

func ChangeRange(minTemperture, maxTemperture, val int, op string) (int, int) {
	newMin := minTemperture
	newMax := maxTemperture
	if op == ">=" && val > newMin {
		newMin = val
	}
	if op == "<=" && val < newMax {
		newMax = val
	}
	return newMin, newMax
}
func main() {
	var countDep int
	if _, err := fmt.Scan(&countDep); err != nil {
		return
	}
	for range countDep {
		var countWor int
		if _, err := fmt.Scan(&countWor); err != nil {
			return
		}
		currMin := minTemp
		currMax := maxTemp
		for range countWor {
			var operator string
			var value int
			if _, err := fmt.Scan(&operator, &value); err != nil {
				return
			}
			currMin, currMax = ChangeRange(currMin, currMax, value, operator)
			if currMin <= currMax {
				fmt.Println(currMin)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
