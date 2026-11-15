package currency

import (
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
	XMLName    xml.Name   `xml:"ValCurs"`
	Currencies []Currency `xml:"Valute"`
}

func DecodeXML(reader io.Reader) ([]Currency, error) {
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	var rates xmlRates
	if err := decoder.Decode(&rates); err != nil {
		return nil, fmt.Errorf("decode currencies XML: %w", err)
	}

	if rates.Currencies == nil {
		return []Currency{}, nil
	}

	return rates.Currencies, nil
}

func ReadXML(path string) ([]Currency, error) {
	inputFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file %q: %w", path, err)
	}

	defer func() {
		if closeErr := inputFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close input file %q: %w", path, closeErr)
		}
	}()

	return DecodeXML(inputFile)
}
