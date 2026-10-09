package main

import (
	"fmt"
)

func main() {
	var depts, emps int

	for {
		if _, err := fmt.Scan(&depts, &emps); err != nil {
			break
		}

		for range depts {
			minTemp := 15
			maxTemp := 30

			for range emps {
				var (
					cond string
					temp int
				)

				if _, err := fmt.Scan(&cond, &temp); err != nil {
					break
				}

				if cond == ">=" && temp > minTemp {
					minTemp = temp
				}

				if cond == "<=" && temp < maxTemp {
					maxTemp = temp
				}

				if minTemp <= maxTemp {
					_, _ = fmt.Println(minTemp)
				} else {
					_, _ = fmt.Println(-1)
				}
			}
		}
	}
}
