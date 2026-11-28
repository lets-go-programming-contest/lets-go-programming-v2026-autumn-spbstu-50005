package rates

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/Haist-PM/task-3/internal/filedata"
	"golang.org/x/text/encoding/charmap"
)

var (
	errUnsupportedEncoding = errors.New("unsupported XML encoding")
	errNonFiniteValue      = errors.New("currency value must be finite")
)

type xmlDocument struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Entries []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode  int           `json:"num_code"  xml:"NumCode"`
	CharCode string        `json:"char_code" xml:"CharCode"`
	Value    ExchangeValue `json:"value"     xml:"Value"`
}

type ExchangeValue float64

func (value *ExchangeValue) UnmarshalText(data []byte) error {
	numberText := strings.TrimSpace(string(data))
	whole, fraction, hasComma := strings.Cut(numberText, ",")

	if hasComma {
		numberText = whole + "." + fraction
	}

	const bits = 64

	number, err := strconv.ParseFloat(numberText, bits)
	if err != nil {
		return fmt.Errorf("parse Value %q: %w", string(data), err)
	}

	if math.IsNaN(number) || math.IsInf(number, 0) {
		return fmt.Errorf("%w: %q", errNonFiniteValue, string(data))
	}

	*value = ExchangeValue(number)

	return nil
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	if strings.EqualFold(label, "windows-1251") {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return nil, fmt.Errorf("%w: %q", errUnsupportedEncoding, label)
}

func decodeXML(data []byte, target any) error {
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charsetReader

	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("parse CBR XML: %w", err)
	}

	return nil
}

func Load(path string) ([]Currency, error) {
	document, err := filedata.Read[xmlDocument](path, decodeXML)
	if err != nil {
		return nil, fmt.Errorf("read CBR XML: %w", err)
	}

	if document.Entries == nil {
		return []Currency{}, nil
	}

	return document.Entries, nil
}

func SortDescending(currencies []Currency) {
	sort.SliceStable(currencies, func(first, second int) bool {
		return currencies[first].Value > currencies[second].Value
	})
}
