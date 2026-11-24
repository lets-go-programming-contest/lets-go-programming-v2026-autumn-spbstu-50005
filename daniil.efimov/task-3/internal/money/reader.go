package money

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

var ErrUnsupportedCharset = errors.New("unsupported charset")

type xmlStructure struct {
	XMLName xml.Name     `xml:"ValCurs"`
	Items   []ValuteInfo `xml:"Valute"`
}

func ExtractRates(filePath string) ([]ValuteInfo, error) {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file %q: %w", filePath, err)
	}

	defer func() {
		if closeErr := inputFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close file %q: %w", filePath, closeErr)
		}
	}()

	decoder := xml.NewDecoder(inputFile)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		return nil, fmt.Errorf("%w: %s", ErrUnsupportedCharset, charset)
	}

	var doc xmlStructure
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return doc.Items, nil
}
