package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minTemp       = 15
	maxTemp       = 30
	lessThan      = "<="
	moreThan      = ">="
	noConcessions = -1
)

var ErrUnknownOperator = errors.New("unknown temperature operator")

type TemperatureRange struct {
	minimum int
	maximum int
}

func (t *TemperatureRange) Update(value int, operator string) error {
	if operator != moreThan && operator != lessThan {
		return fmt.Errorf("%w: %q", ErrUnknownOperator, operator)
	}

	if operator == moreThan && value > t.minimum {
		t.minimum = value
	}

	if operator == lessThan && value < t.maximum {
		t.maximum = value
	}

	return nil
}

func (t *TemperatureRange) GetTemperature() int {
	if t.minimum > t.maximum {
		return noConcessions
	}

	return t.minimum
}

func main() {
	var departmentCount int

	if _, err := fmt.Scan(&departmentCount); err != nil {
		log.Printf("read department count: %v", err)

		return
	}

	for range departmentCount {
		var workerCount int

		if _, err := fmt.Scan(&workerCount); err != nil {
			log.Printf("read worker count: %v", err)

			return
		}

		temperature := TemperatureRange{
			minimum: minTemp,
			maximum: maxTemp,
		}

		for range workerCount {
			var (
				operator string
				value    int
			)

			if _, err := fmt.Scan(&operator, &value); err != nil {
				log.Printf("read temperature requirement: %v", err)

				return
			}

			if err := temperature.Update(value, operator); err != nil {
				log.Printf("update temperature range: %v", err)

				return
			}

			fmt.Println(temperature.GetTemperature())
		}
	}
}
