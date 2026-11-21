package currency

import (
	"cmp"
	"encoding/xml"
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
	NumCode  int           `xml:"NumCode" json:"num_code"`
	CharCode string        `xml:"CharCode" json:"char_code"`
	Value    CurrencyValue `xml:"Value" json:"value"`
}

type CurrencyValue float64

func (r *CurrencyValue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := d.DecodeElement(&raw, &start); err != nil {
		return err
	}

	normalized := strings.Replace(strings.TrimSpace(raw), ",", ".", 1)
	val, err := strconv.ParseFloat(normalized, 64)

	if err != nil {
		return err
	}

	*r = CurrencyValue(val)
	return nil
}
