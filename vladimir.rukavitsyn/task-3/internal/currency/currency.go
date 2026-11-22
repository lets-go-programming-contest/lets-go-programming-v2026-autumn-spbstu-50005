package currency

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"
)

const (
	DirMode  os.FileMode = 0o755
	FileMode os.FileMode = 0o644
)

var ErrBadNumber = errors.New("invalid number format")

// Rate хранит курс. В XML ЦБ число записано с запятой: "43,8254".
type Rate float64

func (r *Rate) UnmarshalText(text []byte) error {
	raw := strings.TrimSpace(string(text))

	if strings.Count(raw, ",")+strings.Count(raw, ".") > 1 {
		return fmt.Errorf("%w: %q", ErrBadNumber, raw)
	}

	parsed, err := strconv.ParseFloat(strings.ReplaceAll(raw, ",", "."), 64)
	if err != nil {
		return fmt.Errorf("parse rate %q: %w", raw, err)
	}

	*r = Rate(parsed)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}

type document struct {
	Items []Currency `xml:"Valute"`
}

func Load(path string) ([]Currency, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return nil, fmt.Errorf("open xml: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charset.NewReaderLabel

	var doc document

	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return doc.Items, nil
}

func SortDesc(items []Currency) {
	slices.SortStableFunc(items, func(a, b Currency) int {
		switch {
		case a.Value > b.Value:
			return -1
		case a.Value < b.Value:
			return 1
		default:
			return 0
		}
	})
}

func Save(path string, items []Currency) error {
	if err := os.MkdirAll(filepath.Dir(path), DirMode); err != nil {
		return fmt.Errorf("create directory: %w", err)
	}

	data, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, FileMode); err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}