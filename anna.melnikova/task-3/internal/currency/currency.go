package currency

import (
	"cmp"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	ErrInvalidValueFormat = errors.New("invalid value format")
	ErrUnsupportedCharset = errors.New("unsupported charset")
)

type Rate float64

func (r *Rate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var raw string

	if err := d.DecodeElement(&raw, &start); err != nil {
		return fmt.Errorf("decode rate element: %w", err)
	}

	trimmed := strings.TrimSpace(raw)
	if trimmed == "" || strings.Contains(trimmed, ".") || strings.Count(trimmed, ",") > 1 {
		return fmt.Errorf("%w: %q", ErrInvalidValueFormat, trimmed)
	}

	normalized := strings.Replace(trimmed, ",", ".", 1)

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidValueFormat, err)
	}

	*r = Rate(parsed)

	return nil
}

type Currency struct {
	NumCode  int    `json:"num_code"  xml:"NumCode"`
	CharCode string `json:"char_code" xml:"CharCode"`
	Value    Rate   `json:"value"     xml:"Value"`
}

type ValCurs struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

func ParseXML[T any](path string) (T, error) {
	var result T

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return result, fmt.Errorf("open xml file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charsetReader

	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("decode xml: %w", err)
	}

	return result, nil
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	case "utf-8", "utf8":
		return input, nil
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}
}

func SortDesc(currencies []Currency) {
	slices.SortFunc(currencies, func(a, b Currency) int {
		return cmp.Compare(b.Value, a.Value)
	})
}

func SaveJSON[T any](path string, data T, dirPerm, filePerm os.FileMode) error {
	dir := filepath.Dir(path)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, dirPerm); err != nil {
			return fmt.Errorf("create output directory: %w", err)
		}
	}

	encoded, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	if err := os.WriteFile(path, encoded, filePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
