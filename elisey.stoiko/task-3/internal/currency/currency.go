package currency

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Decimal float64

func (d *Decimal) UnmarshalText(text []byte) error {
	s := strings.Replace(strings.TrimSpace(string(text)), ",", ".", 1)

	val, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return fmt.Errorf("parse decimal %q: %w", text, err)
	}

	*d = Decimal(val)

	return nil
}

type Currency struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value"     xml:"Value"`
}

func SortByValue(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}
