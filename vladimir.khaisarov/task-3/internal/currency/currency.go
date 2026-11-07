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

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	value := strings.ReplaceAll(strings.TrimSpace(string(text)), ",", ".")

	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return fmt.Errorf("parse value %q: %w", value, err)
	}

	*r = Rate(number)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}

type valCurs struct {
	Currencies []Currency `xml:"Valute"`
}

func ReadXML(path string) ([]Currency, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var data valCurs

	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("decode input file: %w", err)
	}

	return data.Currencies, nil
}

func SortByValueDesc(currencies []Currency) {
	sort.Slice(currencies, func(first, second int) bool {
		return currencies[first].Value > currencies[second].Value
	})
}

func WriteJSON(path string, currencies []Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write output file: %w", err)
	}

	return nil
}
