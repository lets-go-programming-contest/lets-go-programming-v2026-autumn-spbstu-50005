package main

import (
	"fmt"
	"os"
)

func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		os.Exit(0)
	}

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
