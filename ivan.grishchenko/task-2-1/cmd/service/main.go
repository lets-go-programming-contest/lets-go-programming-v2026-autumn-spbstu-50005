package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		depts, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		if !scanner.Scan() {
			break
		}

		emps, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}

		for range depts {
			minTemp := 15
			maxTemp := 30

			for range emps {
				if !scanner.Scan() {
					return
				}

				cond := scanner.Text()

				if !scanner.Scan() {
					return
				}

				temp, err := strconv.Atoi(scanner.Text())
				if err != nil {
					return
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
