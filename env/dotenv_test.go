package env

import (
	"os"
	"testing"
)

func TestLoadDotEnv(t *testing.T) {
	content := "FOO=bar\nBAZ=qux"
	tmpFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(content)); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	envMap, err := LoadDotEnv(tmpFile.Name())
	if err != nil {
		t.Fatalf("LoadDotEnv returned error: %v", err)
	}
	if envMap["FOO"] != "bar" {
		t.Errorf("Expected FOO=bar, got FOO=%s", envMap["FOO"])
	}
	if envMap["BAZ"] != "qux" {
		t.Errorf("Expected BAZ=qux, got BAZ=%s", envMap["BAZ"])
	}
}
