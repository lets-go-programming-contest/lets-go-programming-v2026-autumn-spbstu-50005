package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJSON(path string, currencies []Currency, dirMode, fileMode os.FileMode) error {
	data, err := json.MarshalIndent(currencies, "", "  ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(path), dirMode); err != nil {
		return fmt.Errorf("create output directories: %w", err)
	}

	if err := os.WriteFile(path, append(data, '\n'), fileMode); err != nil {
		return fmt.Errorf("write JSON: %w", err)
	}

	return nil
}
