package main

import "fmt"

const (
	minTemperature = 15
	maxTemperature = 30
)

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

		minAllowed := minTemperature
		maxAllowed := maxTemperature

		for range workerCount {
			var (
				sign        string
				temperature int
			)

			if _, err := fmt.Scan(&sign, &temperature); err != nil {
				return
			}

			switch sign {
			case ">=":
				if temperature > minAllowed {
					minAllowed = temperature
				}
			case "<=":
				if temperature < maxAllowed {
					maxAllowed = temperature
				}
			}

			if minAllowed <= maxAllowed {
				fmt.Println(minAllowed)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
