package currency

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

var ErrUnsupportedCharset = errors.New("unsupported charset")

type xmlRates struct {
	Currencies []Currency `xml:"Valute"`
}

func DecodeXML(reader io.Reader) ([]Currency, error) {
	decoder := xml.NewDecoder(reader)

	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch strings.ToLower(strings.TrimSpace(charset)) {
		case "windows-1251", "cp1251":
			return charmap.Windows1251.NewDecoder().Reader(input), nil

		case "utf-8", "utf8":
			return input, nil

		default:
			return nil, fmt.Errorf(
				"%w: %s",
				ErrUnsupportedCharset,
				charset,
			)
		}
	}

	var rates xmlRates

	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode currencies XML: %w", err)
	}

	return rates.Currencies, nil
}

func ReadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read input file %q: %w", path, err)
	}

	return DecodeXML(bytes.NewReader(data))
}
