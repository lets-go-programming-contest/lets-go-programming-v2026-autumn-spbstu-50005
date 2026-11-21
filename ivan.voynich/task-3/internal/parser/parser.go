package parser

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func ReadXML[T any](filepath string) (T, error) {
	var result T

	data, err := os.ReadFile(filepath)

	if err != nil {
		return result, fmt.Errorf("error reading XML file: %w", err)
	}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		if strings.EqualFold(charset, "windows-1251") {
			return charmap.Windows1251.NewDecoder().Reader(input), nil
		}
		return input, nil
	}

	if err := decoder.Decode(&result); err != nil {
		return result, fmt.Errorf("error decoding XML file: %w", err)
	}

	return result, nil
}

func EnsureDir(filePath string, perm os.FileMode) error {
	dir := filepath.Dir(filePath)

	if dir == "" || dir == "." {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("error creating directory %s: %w", dir, err)
	}

	return nil
}

func WriteJSON[T any](filepath string, data T, perm os.FileMode) error {
	jsonBytes, err := json.MarshalIndent(data, "", "	")

	if err != nil {
		return fmt.Errorf("error endcoding json: %w", err)
	}

	if err := os.WriteFile(filepath, jsonBytes, perm); err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}

	return nil
}
