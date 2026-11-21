package filedata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Read[T any](path string, decode func([]byte, any) error) (T, error) {
	var result T

	data, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return result, fmt.Errorf("read %q: %w", path, err)
	}

	if err := decode(data, &result); err != nil {
		return result, fmt.Errorf("decode %q: %w", path, err)
	}

	return result, nil
}

func WriteJSON[T any](path string, result T, directoryMode, fileMode os.FileMode) error {
	data, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		return fmt.Errorf("encode JSON: %w", err)
	}

	path = filepath.Clean(path)
	directory := filepath.Dir(path)

	if err := os.MkdirAll(directory, directoryMode); err != nil {
		return fmt.Errorf("create directory %q: %w", directory, err)
	}

	if err := os.WriteFile(path, data, fileMode); err != nil {
		return fmt.Errorf("write %q: %w", path, err)
	}

	return nil
}
