package currency

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/net/html/charset"
)

var errMissingValCurs = errors.New("incorrect XML signature: ValCurs is missing")

type valCursDocument struct {
	XMLName xml.Name   `xml:"ValCurs"`
	Valutes []Currency `xml:"Valute"`
}

func LoadXML(path string) ([]Currency, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("could not read XML: %w", err)
	}

	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel

	var doc valCursDocument
	if err := decoder.Decode(&doc); err != nil {
		return nil, fmt.Errorf("error of decoding XML: %w", err)
	}

	if doc.XMLName.Local != "ValCurs" {
		return nil, errMissingValCurs
	}

	return doc.Valutes, nil
}

func SortDesc(items []Currency) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Value > items[j].Value
	})
}

func SaveJSON(path string, items []Currency, dirPerm os.FileMode) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("could not to make dir %q: %w", dir, err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not to make file %q: %w", path, err)
	}

	defer func() {
		if closeErr := file.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close output file: %w", closeErr)
		}
	}()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	if err = enc.Encode(items); err != nil {
		return fmt.Errorf("error of coding to JSON: %w", err)
	}

	return nil
}