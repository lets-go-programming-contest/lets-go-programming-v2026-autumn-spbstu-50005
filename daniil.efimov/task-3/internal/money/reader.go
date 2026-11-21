package money

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"golang.org/x/text/encoding/charmap"
)

type xmlStructure struct {
	XMLName xml.Name     `xml:"ValCurs"`
	Items   []ValuteInfo `xml:"Valute"`
}

func ExtractRates(filePath string) ([]ValuteInfo, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file %q: %w", filePath, err)
	}
	defer f.Close()

	decoder := xml.NewDecoder(f)
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if charset == "windows-1251" {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return nil, fmt.Errorf("unsupported charset: %s", charset)
	}

	var doc xmlStructure
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode XML: %w", err)
	}

	return doc.Items, nil
}
