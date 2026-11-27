package currency

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var (
	ErrInput  = errors.New("input file error")
	ErrOutput = errors.New("output file error")
	ErrValue  = errors.New("invalid currency value")
	ErrConfig = errors.New("config error")
)

type RateValue float64

func (rv *RateValue) UnmarshalText(data []byte) error {
	normalized := strings.ReplaceAll(strings.TrimSpace(string(data)), ",", ".")
	if normalized == "" {
		return fmt.Errorf("%w: empty value", ErrValue)
	}

	parsed, err := strconv.ParseFloat(normalized, 64)
	if err != nil {
		return fmt.Errorf("%w %q: %w", ErrValue, string(data), err)
	}

	*rv = RateValue(parsed)

	return nil
}

type Rate struct {
	NumCode  int       `xml:"NumCode" json:"num_code"`
	CharCode string    `xml:"CharCode" json:"char_code"`
	Value    RateValue `xml:"Value" json:"value"`
}

type valCursDocument struct {
	XMLName xml.Name `xml:"ValCurs"`
	Items   []Rate   `xml:"Valute"`
}

func charsetReader(charset string, input io.Reader) (io.Reader, error) {
	switch strings.ToLower(charset) {
	case "windows-1251", "cp1251":
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	case "utf-8", "utf8":
		return input, nil
	}

	return nil, fmt.Errorf("%w: unsupported charset %q", ErrInput, charset)
}

func ReadValutes(path string) ([]Rate, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("%w: open %q: %w", ErrInput, path, err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = charsetReader

	var curs valCursDocument
	if err := decoder.Decode(&curs); err != nil {
		return nil, fmt.Errorf("%w: parse XML %q: %w", ErrInput, path, err)
	}

	return curs.Items, nil
}

func WriteJSON(path string, data []Rate, dirMode, fileMode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
		return fmt.Errorf("%w: create dir for %q: %w", ErrOutput, path, err)
	}

	payload, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return fmt.Errorf("%w: encode JSON: %w", ErrOutput, err)
	}

	if err := os.WriteFile(path, payload, fileMode); err != nil {
		return fmt.Errorf("%w: write JSON %q: %w", ErrOutput, path, err)
	}

	return nil
}
