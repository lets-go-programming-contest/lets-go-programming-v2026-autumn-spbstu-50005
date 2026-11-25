package currency

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	value := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("parse rate %q: %w", value, err)
	}

	*r = Rate(number)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}

type ValCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func ReadXML[T any](path string) (T, error) {
	var data T

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return data, fmt.Errorf("open input file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	if err := decoder.Decode(&data); err != nil {
		return data, fmt.Errorf("decode input file: %w", err)
	}

	return data, nil
}

func SortByValueDesc(currencies []Currency) {
	sort.Slice(currencies, func(first, second int) bool {
		return currencies[first].Value > currencies[second].Value
	})
}

func WriteJSON[T any](path string, data T, dirPerm, filePerm os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPerm); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	encoded, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, encoded, filePerm); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}