package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestRunInvalidArguments(t *testing.T) {
	t.Parallel()

	for _, args := range [][]string{nil, {"-config", ""}, {"-config"}, {"-unknown"}, {"-config", "a.yaml", "extra"}} {
		if err := run(args); err == nil {
			t.Fatalf("expected error for %v", args)
		}
	}
}

func TestRun(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()

	input, err := filepath.Abs("../../testdata/currencies.xml")
	if err != nil {
		t.Fatal(err)
	}

	output := filepath.Join(dir, "nested", "out.json")

	data, err := yaml.Marshal(map[string]string{"input-file": input, "output-file": output})
	if err != nil {
		t.Fatal(err)
	}

	config := filepath.Join(dir, "config.yaml")
	if err := os.WriteFile(config, data, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := run([]string{"-config", config}); err != nil {
		t.Fatal(err)
	}

	data, err = os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}

	var rows []struct {
		Code string `json:"char_code"`
	}

	if err := json.Unmarshal(data, &rows); err != nil {
		t.Fatal(err)
	}

	if len(rows) != 3 || rows[0].Code != "EUR" {
		t.Fatalf("unexpected output: %s", data)
	}
}
