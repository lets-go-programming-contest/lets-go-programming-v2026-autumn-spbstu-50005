package currency

import (
	"bytes"
	"cmp"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/htmlindex"
)

type Rate struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  string `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ParseFile(path string) ([]Rate, error) {
	content, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("cannot read input file: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(content))
	decoder.CharsetReader = charsetReader

	var doc valCurs

	err = decoder.Decode(&doc)
	if err != nil {
		return nil, fmt.Errorf("cannot decode xml: %w", err)
	}

	rates := make([]Rate, 0, len(doc.Valutes))

	for _, item := range doc.Valutes {
		rate, err := item.toRate()
		if err != nil {
			return nil, err
		}

		rates = append(rates, rate)
	}

	return rates, nil
}

func SortByValueDesc(rates []Rate) {
	slices.SortStableFunc(rates, func(left, right Rate) int {
		return cmp.Compare(right.Value, left.Value)
	})
}

func (v valute) toRate() (Rate, error) {
	rate := Rate{
		NumCode:  0,
		CharCode: strings.TrimSpace(v.CharCode),
		Value:    0,
	}

	if numCode := strings.TrimSpace(v.NumCode); numCode != "" {
		parsed, err := strconv.Atoi(numCode)
		if err != nil {
			return Rate{}, fmt.Errorf("invalid NumCode %q: %w", v.NumCode, err)
		}

		rate.NumCode = parsed
	}

	if value := strings.TrimSpace(v.Value); value != "" {
		parsed, err := strconv.ParseFloat(strings.ReplaceAll(value, ",", "."), 64)
		if err != nil {
			return Rate{}, fmt.Errorf("invalid Value %q: %w", v.Value, err)
		}

		rate.Value = parsed
	}

	return rate, nil
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	encoding, err := htmlindex.Get(label)
	if err != nil {
		return nil, fmt.Errorf("unsupported charset %q: %w", label, err)
	}

	return encoding.NewDecoder().Reader(input), nil
}
