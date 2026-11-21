package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	directoryPermissions = 0o755
	filePermissions      = 0o600
	floatBitSize         = 64
)

var (
	errInvalidValue = errors.New("currency value must be finite")
	errExtraXML     = errors.New("unexpected content after ValCurs")
)

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type xmlCurrency struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type exchangeRates struct {
	XMLName    xml.Name      `xml:"ValCurs"`
	Currencies []xmlCurrency `xml:"Valute"`
}

func Read(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read XML %q: %w", path, err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var rates exchangeRates

	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode XML %q: %w", path, err)
	}

	if err := checkEnd(decoder); err != nil {
		return nil, fmt.Errorf("validate XML %q: %w", path, err)
	}

	currencies := make([]Currency, 0, len(rates.Currencies))

	for _, entry := range rates.Currencies {
		item, err := convert(entry)
		if err != nil {
			return nil, fmt.Errorf("currency %q: %w", entry.CharCode, err)
		}

		currencies = append(currencies, item)
	}

	return currencies, nil
}

func convert(entry xmlCurrency) (Currency, error) {
	var item Currency

	text := strings.ReplaceAll(strings.TrimSpace(entry.Value), ",", ".")

	value, err := strconv.ParseFloat(text, floatBitSize)
	if err != nil {
		return item, fmt.Errorf("parse value %q: %w", entry.Value, err)
	}

	if math.IsNaN(value) || math.IsInf(value, 0) {
		return item, errInvalidValue
	}

	return Currency{
		NumCode:  entry.NumCode,
		CharCode: entry.CharCode,
		Value:    value,
	}, nil
}

func checkEnd(decoder *xml.Decoder) error {
	for {
		token, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			return nil
		}

		if err != nil {
			return fmt.Errorf("read trailing XML: %w", err)
		}

		switch value := token.(type) {
		case xml.CharData:
			if strings.TrimSpace(string(value)) != "" {
				return errExtraXML
			}
		case xml.Comment, xml.ProcInst:
		default:
			return errExtraXML
		}
	}
}

func Sort(currencies []Currency) {
	sort.SliceStable(currencies, func(left, right int) bool {
		return currencies[left].Value > currencies[right].Value
	})
}

func Write(path string, currencies []Currency) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), directoryPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON %q: %w", path, err)
	}

	return nil
}
