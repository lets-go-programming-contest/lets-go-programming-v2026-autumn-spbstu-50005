package main

import (
	"fmt"
	"log"
)

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
	if err != nil {
		log.Fatalf("failed to read tests count: %v", err)
	}

	if testsCount <= 0 {
		log.Fatalf("invalid tests count: %d", testsCount)
	}

	for range testsCount {
		processTest()
	}
}

func processTest() {
	var limitsCount int

	_, err := fmt.Scan(&limitsCount)
	if err != nil {
		log.Fatalf("failed to read limits count: %v", err)
	}

	if limitsCount <= 0 {
		log.Fatalf("invalid limits count: %d", limitsCount)
	}

	temp := TemperatureRange{
		min: defaultMin,
		max: defaultMax,
	}

	for range limitsCount {
		sign, degree, err := readConstraint()
		if err != nil {
			log.Fatalf("failed to read constraint: %v", err)
		}

		applyConstraint(&temp, sign, degree)
	}
}

func readConstraint() (string, int, error) {
	var (
		sign   string
		degree int
	)

	_, err := fmt.Scan(&sign, &degree)
	if err != nil {
		return "", 0, fmt.Errorf("scan constraint: %w", err)
	}

	return sign, degree, nil
}

func applyConstraint(temp *TemperatureRange, sign string, degree int) {
	if sign != signGreaterOrEqual && sign != signLessOrEqual {
		return
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
