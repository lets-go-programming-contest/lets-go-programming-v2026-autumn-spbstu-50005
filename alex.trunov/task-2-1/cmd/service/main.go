package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemp       = 15
	maxTemp       = 30
	lessThan      = "<="
	moreThan      = ">="
	noConcessions = -1
)

var errInvalidOperator = errors.New("invalid operator")

func readInt() (int, error) {
	var value int

	if _, err := fmt.Scan(&value); err != nil {
		return 0, fmt.Errorf("scan int: %w", err)
	}

	return value, nil
}

func readWish() (string, int, error) {
	var (
		operator    string
		temperature int
	)

	if _, err := fmt.Scan(&operator, &temperature); err != nil {
		return "", 0, fmt.Errorf("scan wish: %w", err)
	}

	if operator != moreThan && operator != lessThan {
		return "", 0, errInvalidOperator
	}

	return operator, temperature, nil
}

func updateTemp(operator string, lowTemp, highTemp, currentTemp int) (int, int) {
	switch operator {
	case moreThan:
		if currentTemp > lowTemp {
			lowTemp = currentTemp
		}
	case lessThan:
		if currentTemp < highTemp {
			highTemp = currentTemp
		}
	}

	return lowTemp, highTemp
}

func printTemp(lowTemp, highTemp int) {
	if lowTemp <= highTemp {
		fmt.Println(lowTemp)

		return
	}

	fmt.Println(noConcessions)
}

func fail(err error) {
	if err != nil {
		os.Exit(1)
	}
}

func main() {
	totalCases, err := readInt()
	fail(err)

	for range totalCases {
		requestsCount, readErr := readInt()
		fail(readErr)

		lowTemp, highTemp := minTemp, maxTemp

		for range requestsCount {
			operator, currentTemp, wishErr := readWish()
			fail(wishErr)

			lowTemp, highTemp = updateTemp(operator, lowTemp, highTemp, currentTemp)
			printTemp(lowTemp, highTemp)
		}
	}
}
