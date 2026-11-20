package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tarasovea/task-3/internal/currency"
)

const dirPerm = 0o755

func Save(path string, currencies []currency.Currency) error {
	err := os.MkdirAll(filepath.Dir(path), dirPerm)
	if err != nil {
		return fmt.Errorf("create output directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}

	defer func() {
		_ = file.Close()
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")

	err = encoder.Encode(currencies)
	if err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}
