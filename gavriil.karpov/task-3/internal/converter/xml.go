package converter

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

type Document[T any] struct {
	XMLName xml.Name `xml:"ValCurs"`
	Items   []T      `xml:"Valute"`
}

func ReadXML[T any](path string) ([]T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open xml file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	decoder := xml.NewDecoder(file)
	decoder.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(label, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}

		if strings.EqualFold(label, "utf-8") {
			return input, nil
		}

		return nil, fmt.Errorf("%w: %q", ErrUnsupportedCharset, label)
	}

	var root Document[T]
	if err := decoder.Decode(&root); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	return root.Items, nil
}
