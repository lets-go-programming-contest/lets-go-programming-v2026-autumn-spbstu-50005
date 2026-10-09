package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var depts, emps int

	for {
		if _, err := fmt.Fscan(reader, &depts, &emps); err != nil {
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

				if _, err := fmt.Fscan(reader, &cond, &temp); err != nil {
					return
				}

				if cond == ">=" && temp > minTemp {
					minTemp = temp
				}

				if cond == "<=" && temp < maxTemp {
					maxTemp = temp
				}

				if minTemp <= maxTemp {
					fmt.Fprintln(writer, minTemp)
				} else {
					fmt.Fprintln(writer, -1)
				}
			}
		}
	}
}
