package currency

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

var (
	ErrInvalidValueFormat = errors.New("invalid value format")
	ErrUnsupportedCharset = errors.New("unsupported charset")
)

type CommaFloat float64

func isDigits(s string) bool {
	if s == "" {
		return false
	}

	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return true
}

func (r *CommaFloat) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))

	wholeStr, afterStr, found := strings.Cut(s, ",")

	if !isDigits(wholeStr) || (!isDigits(afterStr) && found) {
		return fmt.Errorf("%w: %q", ErrInvalidValueFormat, text)
	}

	if found {
		wholeStr += "." + afterStr
	}

	val, err := strconv.ParseFloat(wholeStr, 64)
	if err != nil {
		return fmt.Errorf("parse Value to float64: %w", err)
	}

	*r = CommaFloat(val)

	return nil
}

type Valute struct {
	NumCode  int        `json:"num_code"  xml:"NumCode"`
	CharCode string     `json:"char_code" xml:"CharCode"`
	Value    CommaFloat `json:"value"     xml:"Value"`
}

type valCurs struct {
	Valutes []Valute `xml:"Valute"`
}

func ReadXML(path string) ([]Valute, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(label, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %q", ErrUnsupportedCharset, label)
	}

	var data valCurs
	if err := decoder.Decode(&data); err != nil {
		return nil, fmt.Errorf("parse ValCurs: %w", err)
	}

	return data.Valutes, nil
}

func SortByValues(valutes []Valute) {
	sort.Slice(valutes, func(i int, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

func WriteJSON(path string, valutes []Valute) error {
	if err := os.MkdirAll(filepath.Dir(path), dirPermissions); err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
