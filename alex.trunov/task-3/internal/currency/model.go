package currency

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type Rate float64

func (r *Rate) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode xml element: %w", err)
	}

	raw = strings.TrimSpace(raw)
	raw = strings.Replace(raw, ",", ".", 1)

	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return fmt.Errorf("parse float rate: %w", err)
	}

	*r = Rate(val)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}