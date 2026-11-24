package currency

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func WriteJSON(
	path string,
	currencies []Currency,
	directoryPermissions os.FileMode,
	filePermissions os.FileMode,
) error {
	outputDirectory := filepath.Dir(path)

	if err := os.MkdirAll(outputDirectory, directoryPermissions); err != nil {
		return fmt.Errorf(
			"create output directory %q: %w",
			outputDirectory,
			err,
		)
	}

	data, err := json.MarshalIndent(currencies, "", "    ")
	if err != nil {
		return fmt.Errorf("encode currencies to JSON: %w", err)
	}

	data = append(data, '\n')

	if err := os.WriteFile(path, data, filePermissions); err != nil {
		return fmt.Errorf("write JSON %q: %w", path, err)
	}

	return nil
}
