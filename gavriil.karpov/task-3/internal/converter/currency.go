package converter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type DecimalValue float64

func (d *DecimalValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	normalized := strings.Replace(raw, ",", ".", 1)

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*d = DecimalValue(parsed)

	return nil
}

func (d DecimalValue) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(float64(d))
	if err != nil {
		return nil, fmt.Errorf("marshal decimal value: %w", err)
	}

	return data, nil
}

type Currency struct {
	NumCode  int          `json:"num_code"  xml:"NumCode"`
	CharCode string       `json:"char_code" xml:"CharCode"`
	Value    DecimalValue `json:"value"     xml:"Value"`
}
