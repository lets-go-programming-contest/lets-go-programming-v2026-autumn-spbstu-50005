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

var (
	ErrUnsupportedCharset = errors.New("unsupported charset")
	ErrInvalidValueFormat = errors.New("invalid value format")
)

type Decimal float64

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

func (d *Decimal) UnmarshalText(text []byte) error {
	s := strings.TrimSpace(string(text))

	wholeStr, afterStr, found := strings.Cut(s, ",")

	if !isDigits(wholeStr) || (found && !isDigits(afterStr)) {
		return fmt.Errorf("%w: %q", ErrInvalidValueFormat, text)
	}

	if found {
		wholeStr += "." + afterStr
	}

	val, err := strconv.ParseFloat(wholeStr, 64)

	if err != nil {
		return fmt.Errorf("parse decimal %q: %w", text, err)
	}

	*d = Decimal(val)

	return nil
}

type xmlValCurs struct {
	Valutes []Valute `xml:"Valute"`
}

type Valute struct {
	NumCode  int     `json:"num_code"  xml:"NumCode"`
	CharCode string  `json:"char_code" xml:"CharCode"`
	Value    Decimal `json:"value"     xml:"Value"`
}

func LoadFromXML(path string) ([]Valute, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, fmt.Errorf("open xml: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	var curs xmlValCurs

	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return curs.Valutes, nil
}

func SortByValueDesc(valutes []Valute) {
	sort.Slice(valutes, func(i, j int) bool {
		return valutes[i].Value > valutes[j].Value
	})
}

func SaveToJSON(path string, valutes []Valute, dirPerm, filePerm os.FileMode) error {
	dir := filepath.Dir(path)

	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	data, err := json.MarshalIndent(valutes, "", "    ")

	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write json: %w", err)
	}

	return nil
}
