package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Value  `json:"value"     xml:"Value"`
}

func SortByValue(currencies []Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

var (
	ErrUnsupportedEncoding = errors.New("unsupported XML encoding")
	ErrInvalidValue        = errors.New("invalid currency Value")
)

type xmlRates struct {
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func ReadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML: %w", err)
	}

	rates, err := decodeXML(data)
	if err != nil {
		return nil, err
	}

	return rates.Currencies, nil
}

func decodeXML(data []byte) (xmlRates, error) {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if !strings.EqualFold(charset, "windows-1251") {
			return nil, fmt.Errorf("%w: %s", ErrUnsupportedEncoding, charset)
		}

		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	var rates xmlRates

	rates.Currencies = make([]Currency, 0)

	if err := decoder.Decode(&rates); err != nil {
		return xmlRates{}, fmt.Errorf("decode XML: %w", err)
	}

	return rates, nil
}

type Value float64

func (value *Value) UnmarshalText(data []byte) error {
	text := strings.TrimSpace(string(data))
	whole, fraction, hasComma := strings.Cut(text, ",")

	if whole == "" || (hasComma && fraction == "") {
		return ErrInvalidValue
	}

	for _, part := range []string{whole, fraction} {
		for _, digit := range part {
			if digit < '0' || digit > '9' {
				return ErrInvalidValue
			}
		}
	}

	if hasComma {
		text = whole + "." + fraction
	}

	number, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}

	*value = Value(number)

	return nil
}

func WriteJSON(path string, currencies []Currency, dirMode, fileMode os.FileMode) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
		return fmt.Errorf("create output directories: %w", err)
	}

	if err := os.WriteFile(path, append(data, '\n'), fileMode); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
