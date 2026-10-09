package main
import (
	"fmt"
)
func main() {
	var depts int
	var emps int
	if _, err := fmt.Scan(&depts, &emps); err != nil {
		return
	}
	for i := 0; i < depts; i++ {
		minTemp := 15
		maxTemp := 30
		for j := 0; j < emps; j++ {
			var cond string
			var temp int
			if _, err := fmt.Scan(&cond, &temp); err != nil {
				return
			}
			if cond == ">=" && temp > minTemp {
				minTemp = temp
			}
			if cond == "<=" && temp < maxTemp {
				maxTemp = temp
			}
			if minTemp <= maxTemp {
				fmt.Println(minTemp)
			} else {
				fmt.Println(-1)
			}
		}
	}
}
