package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	dirPermissions  = 0o755
	filePermissions = 0o644
)

func SaveJSON(path string, data any) error {
	err := os.MkdirAll(filepath.Dir(path), dirPermissions)
	if err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	content, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("cannot encode json: %w", err)
	}

	err = os.WriteFile(path, content, filePermissions)
	if err != nil {
		return fmt.Errorf("cannot write output file: %w", err)
	}

	return nil
}
