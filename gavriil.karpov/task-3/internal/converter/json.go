package converter

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJSON[T any](path string, items []T, dirPerm os.FileMode, filePerm os.FileMode) error {
	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, dirPerm); err != nil {
		return fmt.Errorf("create directory %s: %w", dir, err)
	}

	if err := os.WriteFile(path, data, filePerm); err != nil {
		return fmt.Errorf("write json file: %w", err)
	}

	return nil
}
