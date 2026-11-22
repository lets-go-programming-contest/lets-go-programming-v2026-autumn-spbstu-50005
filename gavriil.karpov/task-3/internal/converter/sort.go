package converter

import "slices"

func SortByValue(currencies []Currency) {
	slices.SortFunc(currencies, func(a, b Currency) int {
		valA := float64(a.Value)
		valB := float64(b.Value)

		if valA > valB {
			return -1
		}
		if valA < valB {
			return 1
		}
		return 0
	})
}
