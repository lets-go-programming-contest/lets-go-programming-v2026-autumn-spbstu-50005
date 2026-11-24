package currency

import (
	"cmp"
	"encoding/xml"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func (v *ValCurs) SortValueByDesc() {
	slices.SortFunc(v.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

type Valute struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    CurrencyValue `json:"value"     xml:"Value"`
}

type CurrencyValue float64

func (r *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode currency value: %w", err)
	}

	normalized := strings.Replace(strings.TrimSpace(raw), ",", ".", 1)

	val, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse currency float %q: %w", normalized, err)
	}

	*r = CurrencyValue(val)

	return nil
}
