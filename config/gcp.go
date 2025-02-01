package config

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetActiveProject returns the active GCP project from gcloud.
func GetActiveProject() (string, error) {
	if _, err := exec.LookPath("gcloud"); err != nil {
		return "", fmt.Errorf("gcloud not found in PATH")
	}
	cmd := exec.Command("gcloud", "config", "get-value", "project", "--quiet")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("gcloud command failed: %v", err)
	}
	project := strings.TrimSpace(string(out))
	if project == "" || project == "unset" {
		return "", fmt.Errorf("active project not set in gcloud")
	}
	return project, nil
}

// UpdateCloudRunEnv updates the Cloud Run service's environment variables using the gcloud CLI.
// It builds a command like:
//
//	gcloud run services update <service> --project <project> --region <region> --update-env-vars KEY1=VALUE1,KEY2=VALUE2 --quiet
//
// If serviceAccountPath is provided, it sets the GOOGLE_APPLICATION_CREDENTIALS environment variable.
func UpdateCloudRunEnv(project, region, service string, envVars map[string]string, serviceAccountPath string) error {
	// Build a comma-separated list of key=value pairs.
	var pairs []string
	for k, v := range envVars {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	envString := strings.Join(pairs, ",")

	// Construct the gcloud command arguments.
	args := []string{
		"run", "services", "update", service,
		"--project", project,
		"--region", region,
		"--update-env-vars", envString,
		"--quiet",
	}

	cmd := exec.Command("gcloud", args...)
	// If a service account file is provided, set GOOGLE_APPLICATION_CREDENTIALS.
	if serviceAccountPath != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("GOOGLE_APPLICATION_CREDENTIALS=%s", serviceAccountPath))
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("gcloud command failed: %v, output: %s", err, string(output))
	}

	fmt.Printf("gcloud output:\n%s\n", string(output))
	return nil
}

// DryRunCloudRunEnv simulates the update by printing the gcloud command that would be executed.
func DryRunCloudRunEnv(project, region, service string, envVars map[string]string, serviceAccountPath string) error {
	var pairs []string
	for k, v := range envVars {
		pairs = append(pairs, fmt.Sprintf("%s=%s", k, v))
	}
	envString := strings.Join(pairs, ",")
	args := []string{
		"run", "services", "update", service,
		"--project", project,
		"--region", region,
		"--update-env-vars", envString,
		"--quiet",
	}

	fmt.Println("Dry Run - The following gcloud command would be executed:")
	fmt.Printf("gcloud %s\n", strings.Join(args, " "))
	if serviceAccountPath != "" {
		fmt.Printf("with GOOGLE_APPLICATION_CREDENTIALS=%s\n", serviceAccountPath)
	}
	return nil
}
