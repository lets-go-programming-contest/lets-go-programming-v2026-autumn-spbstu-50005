package main

import "fmt"

const (
	defaultMin               = 15
	defaultMax               = 30
	signGreaterOrEqual       = ">="
	signLessOrEqual          = "<="
	invalidTemperatureResult = -1
)

type TemperatureRange struct {
	min int
	max int
}

func main() {
	var testsCount int

	_, err := fmt.Scan(&testsCount)
	if err != nil || testsCount <= 0 {
		return
	}

	for range testsCount {
		processTest()
	}
}

func processTest() {
	var limitsCount int

	_, err := fmt.Scan(&limitsCount)
	if err != nil || limitsCount <= 0 {
		return
	}

	temp := TemperatureRange{
		min: defaultMin,
		max: defaultMax,
	}

	for range limitsCount {
		var (
			sign   string
			degree int
		)

		_, err := fmt.Scan(&sign, &degree)
		if err != nil {
			return
		}

		if sign != signGreaterOrEqual && sign != signLessOrEqual {
			continue
		}

		if sign == signGreaterOrEqual && degree > temp.min {
			temp.min = degree
		}

		if sign == signLessOrEqual && degree < temp.max {
			temp.max = degree
		}

		if temp.min <= temp.max {
			fmt.Println(temp.min)
		} else {
			fmt.Println(invalidTemperatureResult)
		}
	}
}
