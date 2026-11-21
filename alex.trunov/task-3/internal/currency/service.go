package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/net/html/charset"
)

func LoadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not reat XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var currencies []Currency
	hasRoot := false

	for {
		tok, err := decoder.Token()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error of reading XML: %w", err)
		}

		startElem, ok := tok.(xml.StartElement)
		if !ok {
			continue
		}

		switch startElem.Name.Local {
		case "ValCurs":
			hasRoot = true
		case "Valute":
			var item Currency
			if err := decoder.DecodeElement(&item, &startElem); err != nil {
				return nil, fmt.Errorf("error of decoding Valute: %w", err)
			}
			currencies = append(currencies, item)
		}
	}

	if !hasRoot {
		return nil, errors.New("incorrect XML signature: ValCurs is missing")
	}

	return currencies, nil
}

func SortDesc(items []Currency) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
}

func SaveJSON(path string, items []Currency) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("could not to make dir %q: %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not to make file %q: %w", path, err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(items); err != nil {
		return fmt.Errorf("error of coding to JSON: %w", err)
	}

	return nil
}
