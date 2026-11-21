package currency

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

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

		switch token := token.(type) {
		case xml.StartElement:
			if seenRoot {
				return xmlRates{}, fmt.Errorf("%w: XML must contain one ValCurs root", ErrInvalidXML)
			}

			if err := decoder.DecodeElement(&rates, &token); err != nil {
				return xmlRates{}, fmt.Errorf("decode XML: %w", err)
			}

			seenRoot = true
		default:
			if err := validateOutsideRoot(token, seenRoot); err != nil {
				return xmlRates{}, err
			}
		}
	}

	if !seenRoot {
		return xmlRates{}, fmt.Errorf("%w: XML has no ValCurs root", ErrInvalidXML)
	}

	return rates, nil
}

func validateOutsideRoot(token xml.Token, seenRoot bool) error {
	switch token := token.(type) {
	case xml.CharData:
		if strings.TrimSpace(string(token)) != "" {
			return fmt.Errorf("%w: unexpected text outside ValCurs", ErrInvalidXML)
		}
	case xml.Comment, xml.ProcInst:
	case xml.Directive:
		if seenRoot {
			return fmt.Errorf("%w: unexpected XML directive after ValCurs", ErrInvalidXML)
		}
	default:
		return fmt.Errorf("%w: unexpected XML content outside ValCurs", ErrInvalidXML)
	}

	return nil
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
