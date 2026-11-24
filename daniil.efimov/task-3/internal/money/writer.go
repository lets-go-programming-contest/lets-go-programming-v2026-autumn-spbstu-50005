package money

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const folderPermission = 0o755

func SaveAsJSON(destPath string, data []ValuteInfo) error {
	dir := filepath.Dir(destPath)
	if err := os.MkdirAll(dir, folderPermission); err != nil {
		return fmt.Errorf("create directory %q: %w", dir, err)
	}

	outputFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("create file %q: %w", destPath, err)
	}

	defer func() {
		if closeErr := outputFile.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close file %q: %w", destPath, closeErr)
		}
	}()

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	return nil
}
