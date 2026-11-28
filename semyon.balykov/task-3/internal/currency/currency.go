package currency

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var ErrInvalidValue = errors.New("invalid currency value")

type Decimal float64

func (d *Decimal) UnmarshalText(text []byte) error {
	value := strings.TrimSpace(string(text))

	integerPart, fractionalPart, hasComma := strings.Cut(value, ",")

	if hasComma {
		if strings.Contains(fractionalPart, ",") {
			return fmt.Errorf("%w: %q", ErrInvalidValue, value)
		}

		value = integerPart + "." + fractionalPart
	}

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("%w: %q: %w", ErrInvalidValue, value, err)
	}

	*d = Decimal(number)

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
