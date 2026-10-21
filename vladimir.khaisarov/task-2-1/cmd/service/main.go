package main

import "fmt"

const (
	minTemp       = 15
	maxTemp       = 30
	lessThan      = "<="
	moreThan      = ">="
	noConcessions = -1
)

type TemperatureRange struct {
	minimum int
	maximum int
}

func (t *TemperatureRange) Update(value int, operator string) {
	if operator == moreThan && value > t.minimum {
		t.minimum = value
	}

	if operator == lessThan && value < t.maximum {
		t.maximum = value
	}
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
		return
	}

	for range departmentCount {
		var workerCount int

		if _, err := fmt.Scan(&workerCount); err != nil {
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
				return
			}

			temperature.Update(value, operator)

			fmt.Println(temperature.GetTemperature())
		}
	}
}
