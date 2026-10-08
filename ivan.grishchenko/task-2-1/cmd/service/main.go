package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	for i := 0; i < n; i++ {
		minTemp := 15
		maxTemp := 30

		for j := 0; j < k; j++ {
			var op string
			var t int
			fmt.Scan(&op, &t)

			if op == ">=" && t > minTemp {
				minTemp = t
			} else if op == "<=" && t < maxTemp {
				maxTemp = t
			}

			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
