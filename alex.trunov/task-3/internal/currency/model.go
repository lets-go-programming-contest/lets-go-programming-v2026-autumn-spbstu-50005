package currency

import (
	"encoding/xml"
	"strconv"
	"strings"
)

type Rate float64

func (r *Rate) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var raw string
	if err := decoder.DecodeElement(&raw, &start); err != nil {
		return err
	}

	raw = strings.TrimSpace(raw)
	raw = strings.Replace(raw, ",", ".", 1)
	val, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return err
	}

	*r = Rate(val)
	return nil
}

type Currency struct {
	NumCode  int    `xml:"NumCode" json:"num_code"`
	CharCode string `xml:"CharCode" json:"char_code"`
	Value    Rate   `xml:"Value" json:"value"`
}
