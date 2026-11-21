package money

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type RateValue float64

func (r *RateValue) UnmarshalText(text []byte) error {
	clean := strings.Replace(string(text), ",", ".", 1)
	val, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", text, err)
	}
	*r = RateValue(val)
	return nil
}

type ValuteInfo struct {
	CodeNum  int       `json:"num_code" xml:"NumCode"`
	CodeChar string    `json:"char_code" xml:"CharCode"`
	Rate     RateValue `json:"value" xml:"Value"`
}

func OrderByRateDesc(list []ValuteInfo) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Rate > list[j].Rate
	})
}
