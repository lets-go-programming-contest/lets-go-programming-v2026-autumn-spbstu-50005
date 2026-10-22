package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	minTemp = 15
	maxTemp = 30
	noTemp  = -1

	maxOffices = 1000
	maxWorkers = 1000

	opAtLeast = ">="
	opAtMost  = "<="
	opEqual   = "="
)

var errInvalidCount = errors.New("invalid count")

type condition struct {
	sign  string
	value int
}

type comfortRange struct {
	min int
	max int
}

func newComfortRange() *comfortRange {
	return &comfortRange{
		min: minTemp,
		max: maxTemp,
	}
}

func (r *comfortRange) apply(sign string, value int) {
	switch sign {
	case opAtLeast:
		if value > r.min {
			r.min = value
		}
	case opAtMost:
		if value < r.max {
			r.max = value
		}
	case opEqual:
		if value >= r.min && value <= r.max {
			r.min = value
			r.max = value
		} else {
			r.min = 1
			r.max = 0
		}
	}
}

func (r *comfortRange) best() int {
	if r.min > r.max {
		return noTemp
	}

	return r.min
}

func readCount(name string, limit int) (int, error) {
	var count int

	if _, err := fmt.Scan(&count); err != nil {
		return 0, fmt.Errorf("read %s count: %w", name, err)
	}

	if count <= 0 || count > limit {
		return 0, fmt.Errorf("%w: %s count must be between 1 and %d, got %d", errInvalidCount, name, limit, count)
	}

	return count, nil
}

func readConditions() ([]condition, error) {
	workerCount, err := readCount("workers", maxWorkers)
	if err != nil {
		return nil, err
	}

	conds := make([]condition, workerCount)
	for workerIndex := range workerCount {
		var (
			sign  string
			value int
		)

		if _, err := fmt.Scan(&sign, &value); err != nil {
			return nil, fmt.Errorf("read condition for worker %d: %w", workerIndex+1, err)
		}

		conds[workerIndex] = condition{
			sign:  sign,
			value: value,
		}
	}

	return conds, nil
}

func solve(conds []condition) []int {
	r := newComfortRange()
	results := make([]int, len(conds))

	for i, c := range conds {
		r.apply(c.sign, c.value)
		results[i] = r.best()
	}

	return results
}

func printResults(res []int) {
	for _, v := range res {
		fmt.Println(v)
	}
}

func processOffice() error {
	conds, err := readConditions()
	if err != nil {
		return err
	}

	results := solve(conds)
	printResults(results)

	return nil
}

func run() error {
	officeCount, err := readCount("offices", maxOffices)
	if err != nil {
		return err
	}

	for officeIndex := range officeCount {
		if err := processOffice(); err != nil {
			return fmt.Errorf("office %d: %w", officeIndex+1, err)
		}
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
