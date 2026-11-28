package currency

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const partsCount = 2

var ErrNoCurrencies = errors.New("no currencies in file")

type Currency struct {
	NumCode  int     `json:"num_code"`
	CharCode string  `json:"char_code"`
	Value    float64 `json:"value"`
}

type valCurs struct {
	Valutes []valute `xml:"Valute"`
}

type valute struct {
	NumCode  int    `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Value    string `xml:"Value"`
}

func ReadXML(path string) ([]Currency, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charsetReader

	var data valCurs

	err = decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	if len(data.Valutes) == 0 {
		return nil, ErrNoCurrencies
	}

	currencies := make([]Currency, 0, len(data.Valutes))

	for _, item := range data.Valutes {
		value, err := parseValue(item.Value)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", item.Value, err)
		}

		currencies = append(currencies, Currency{
			NumCode:  item.NumCode,
			CharCode: item.CharCode,
			Value:    value,
		})
	}

	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].Value > currencies[j].Value
	})

	return currencies, nil
}

func charsetReader(label string, input io.Reader) (io.Reader, error) {
	if strings.ToLower(label) == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}

	return input, nil
}

func parseValue(text string) (float64, error) {
	parts := strings.Split(text, ",")
	if len(parts) == partsCount {
		text = parts[0] + "." + parts[1]
	}

	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float: %w", err)
	}

	return value, nil
}
