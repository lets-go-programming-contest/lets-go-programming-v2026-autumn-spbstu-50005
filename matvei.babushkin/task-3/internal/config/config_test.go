package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Matvei-Babushkin/task-3/internal/config"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	for _, testCase := range []struct {
		name, data string
		valid      bool
	}{
		{"valid", "input-file: input.xml\noutput-file: result/out.json\n", true},
		{"syntax", "input-file: [", false},
		{"missing", "input-file: input.xml\n", false},
		{"blank", "input-file: '  '\noutput-file: out.json", false},
		{"empty", "", false},
		{"unknown key", "input-file: input.xml\noutput-file: out.json\noutput_file: typo.json", false},
		{"multiple documents", "input-file: input.xml\n" +
			"output-file: out.json\n" +
			"---\n" +
			"input-file: other.xml\n" +
			"output-file: other.json", false},
		{"duplicate key", "input-file: input.xml\ninput-file: other.xml\noutput-file: out.json", false},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			path := filepath.Join(t.TempDir(), "config.yaml")
			if err := os.WriteFile(path, []byte(testCase.data), 0o600); err != nil {
				t.Fatal(err)
			}

			cfg, err := config.Load(path)

			if (err == nil) != testCase.valid {
				t.Fatalf("config: %#v, error: %v", cfg, err)
			}

			if testCase.valid && (cfg.InputFile != "input.xml" || cfg.OutputFile != "result/out.json") {
				t.Fatalf("unexpected config: %#v", cfg)
			}
		})
	}

	if _, err := config.Load(filepath.Join(t.TempDir(), "missing.yaml")); err == nil {
		t.Fatal("expected read error")
	}
}
