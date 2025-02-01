package config

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// TestGetActiveProject verifies that GetActiveProject returns a non‐empty project.
// If gcloud is not available or not configured, the test is skipped.
func TestGetActiveProject(t *testing.T) {
	project, err := GetActiveProject()
	if err != nil {
		t.Skipf("Skipping test: %v", err)
	}
	if project == "" {
		t.Error("Expected non-empty project")
	}
}

// TestDryRunCloudRunEnv verifies that DryRunCloudRunEnv prints the expected gcloud command.
// This test uses dummy values and captures stdout.
func TestDryRunCloudRunEnv(t *testing.T) {
	// Dummy parameters for testing.
	project := "dummy-project"
	region := "us-central1"
	service := "dummy-service"
	envVars := map[string]string{
		"FOO": "bar",
		"BAZ": "qux",
	}
	// Dummy service account path; won't be used in a dry run.
	serviceAccountPath := "dummy-sa.json"

	// Redirect stdout to capture output.
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("Failed to create pipe: %v", err)
	}
	os.Stdout = w

	// Execute DryRunCloudRunEnv.
	err = DryRunCloudRunEnv(project, region, service, envVars, serviceAccountPath)
	// Close writer and restore stdout.
	w.Close()
	os.Stdout = oldStdout
	if err != nil {
		t.Fatalf("DryRunCloudRunEnv returned error: %v", err)
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(r)
	if err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// Check that the output includes the expected gcloud command substring.
	expectedSubstring := "gcloud run services update"
	if !strings.Contains(output, expectedSubstring) {
		t.Errorf("Output does not contain expected substring %q: %s", expectedSubstring, output)
	}
}
