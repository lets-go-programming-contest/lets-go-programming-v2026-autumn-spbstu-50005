package currency

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EncodeJSON(writer io.Writer, currencies []Currency) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(currencies); err != nil {
		return fmt.Errorf("encode currencies JSON: %w", err)
	}

	return nil
}

func WriteJSON(path string, currencies []Currency) error {
	outputDir := filepath.Dir(path)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("create output directory %q: %w", outputDir, err)
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file %q: %w", path, err)
	}
	defer outputFile.Close()

	return EncodeJSON(outputFile, currencies)
}
