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

type TempManager struct {
	low  int
	high int
}

func newTempManager() *TempManager {
	return &TempManager{
		low:  minTemp,
		high: maxTemp,
	}
}

func (tm *TempManager) update(operator string, currentTemp int) {
	switch operator {
	case moreThan:
		if currentTemp > tm.low {
			tm.low = currentTemp
		}
	case lessThan:
		if currentTemp < tm.high {
			tm.high = currentTemp
		}
	}
}

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

func (tm *TempManager) printTemp() {
	if tm.low > tm.high {
		fmt.Println(noConcessions)

		return
	}

	fmt.Println(tm.low)
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

		tempManager := newTempManager()

		for range requestsCount {
			operator, currentTemp, wishErr := readWish()
			fail(wishErr)

			tempManager.update(operator, currentTemp)
			tempManager.printTemp()
		}
	}
}
