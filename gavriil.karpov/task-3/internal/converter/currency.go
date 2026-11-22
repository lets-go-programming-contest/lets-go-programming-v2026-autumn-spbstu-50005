package converter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

type DecimalValue float64

type Currency struct {
	NumCode  string       `xml:"NumCode" json:"num_code"`
	CharCode string       `xml:"CharCode" json:"char_code"`
	Value    DecimalValue `xml:"Value" json:"value"`
}

func (d *DecimalValue) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var str string
	if err := decoder.DecodeElement(&str, &start); err != nil {
		return fmt.Errorf("decode element: %w", err)
	}

	normalized := strings.Replace(str, ",", ".", 1)
	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("parse float: %w", err)
	}

	*d = DecimalValue(parsed)
	return nil
}

func (d DecimalValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(float64(d))
}
