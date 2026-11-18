package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

var (
	ErrInvalidNumCode     = errors.New("invalid num code format")
	ErrInvalidValueFormat = errors.New("invalid value format")
)

type valCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

func ParseXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read xml file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var raw valCurs

	if err := decoder.Decode(&raw); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	currencies := make([]Currency, 0, len(raw.Valutes))

	for _, item := range raw.Valutes {
		numCode := 0

		numStr := strings.TrimSpace(item.NumCode)
		if numStr != "" {
			numCode, err = strconv.Atoi(numStr)
			if err != nil {
				return nil, fmt.Errorf("%w: %w", ErrInvalidNumCode, err)
			}
		}

		valStr := strings.ReplaceAll(strings.TrimSpace(item.Value), ",", ".")

		valFloat, err := strconv.ParseFloat(valStr, 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrInvalidValueFormat, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  numCode,
			CharCode: strings.TrimSpace(item.CharCode),
			Value:    valFloat,
		})
	}

	return currencies, nil
}

func SortDesc(currencies []Currency) {
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})
}

func SaveJSON(path string, currencies []Currency) error {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPermissions); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
	}

	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
