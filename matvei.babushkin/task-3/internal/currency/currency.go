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
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

// The exercise sorts the original Value, without dividing by Nominal.
func SortByValue(currencies []Currency) {
	sort.SliceStable(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

var (
	ErrInvalidXML          = errors.New("invalid currencies XML")
	ErrUnsupportedEncoding = errors.New("unsupported XML encoding")
	ErrInvalidValue        = errors.New("invalid currency Value")
)

// XML keeps Value as text because the source uses a decimal comma.
type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type xmlRates struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []xmlCurrency `xml:"Valute"`
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

	currencies := make([]Currency, 0, len(rates.Currencies))

	for _, rate := range rates.Currencies {
		value, err := parseValue(rate.Value)
		if err != nil {
			return nil, fmt.Errorf("parse Value %q for %s: %w", rate.Value, rate.CharCode, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  rate.NumCode,
			CharCode: rate.CharCode,
			Value:    value,
		})
	}

	return currencies, nil
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

	seenRoot := false

	for {
		token, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return xmlRates{}, fmt.Errorf("decode XML: %w", err)
		}

		if start, ok := token.(xml.StartElement); ok {
			if seenRoot {
				return xmlRates{}, ErrInvalidXML
			}

			if err := decoder.DecodeElement(&rates, &start); err != nil {
				return xmlRates{}, fmt.Errorf("decode XML: %w", err)
			}

			seenRoot = true
		} else if text, ok := token.(xml.CharData); ok && strings.TrimSpace(string(text)) != "" {
			return xmlRates{}, ErrInvalidXML
		}
	}

	if !seenRoot {
		return xmlRates{}, ErrInvalidXML
	}

	return rates, nil
}

// parseValue accepts decimal digits with an optional comma and fractional part.
func parseValue(text string) (float64, error) {
	text = strings.TrimSpace(text)
	whole, fraction, hasComma := strings.Cut(text, ",")

	if whole == "" || (hasComma && fraction == "") {
		return 0, ErrInvalidValue
	}

	for _, part := range []string{whole, fraction} {
		for _, digit := range part {
			if digit < '0' || digit > '9' {
				return 0, ErrInvalidValue
			}
		}
	}

	if hasComma {
		text = whole + "." + fraction
	}

	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrInvalidValue, err)
	}

	return value, nil
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
