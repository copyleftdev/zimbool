package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/copyleftdev/zimbool/config"
	"github.com/copyleftdev/zimbool/env"
	"github.com/copyleftdev/zimbool/tf"
)

// preflightCheck verifies required system components and prints emoji-rich messages.
func preflightCheck(envFile, serviceAccount string) error {
	var errs []string

	// Check for gcloud CLI.
	if _, err := exec.LookPath("gcloud"); err != nil {
		fmt.Println("❌ gcloud CLI not found in PATH")
		errs = append(errs, "gcloud CLI missing")
	} else {
		fmt.Println("✅ gcloud CLI is installed")
	}

	// Check that the .env file exists.
	if _, err := os.Stat(envFile); err != nil {
		fmt.Printf("❌ .env file not found: %s\n", envFile)
		errs = append(errs, ".env file missing")
	} else {
		fmt.Printf("✅ .env file found: %s\n", envFile)
	}

	// Check that the service account file exists.
	if _, err := os.Stat(serviceAccount); err != nil {
		fmt.Printf("❌ Service account file not found: %s\n", serviceAccount)
		errs = append(errs, "service account file missing")
	} else {
		fmt.Printf("✅ Service account file found: %s\n", serviceAccount)
	}

	if len(errs) > 0 {
		return fmt.Errorf("preflight check failed: %s", errs)
	}
	return nil
}

func main() {
	// CLI flags.
	envFile := flag.String("env-file", ".env", "Path to the .env file")
	project := flag.String("project", "", "GCP Project ID (if omitted, retrieved from gcloud)")
	region := flag.String("region", "us-central1", "GCP region (e.g. us-central1)")
	service := flag.String("service", "", "Cloud Run service name (required for update)")
	serviceAccount := flag.String("service-account", "", "Path to service account JSON key file (required for update)")
	dryRun := flag.Bool("dry-run", false, "Simulate update without applying changes")
	toHCL := flag.Bool("to-hcl", false, "Transform .env into valid HCL (Terraform) output")
	flag.Parse()

	// Load .env file.
	envVars, err := env.LoadDotEnv(*envFile)
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// If --to-hcl flag is set, transform the .env file to HCL and output it.
	if *toHCL {
		hclOutput := tf.EnvVarsToHCL(envVars)
		fmt.Println(hclOutput)
		os.Exit(0)
	}

	// Validate required flags for Cloud Run update.
	if *service == "" {
		log.Fatal("Cloud Run service name is required (--service)")
	}
	if *serviceAccount == "" {
		log.Fatal("Service account JSON file path is required (--service-account)")
	}

	// Run preflight checks.
	fmt.Println("🛠️ Running preflight checks...")
	if err := preflightCheck(*envFile, *serviceAccount); err != nil {
		log.Fatalf("Preflight check failed: %v", err)
	}
	fmt.Println("✅ Preflight checks passed!")

	// If project flag is omitted, retrieve the active project from gcloud.
	if *project == "" {
		p, err := config.GetActiveProject()
		if err != nil {
			log.Fatalf("Project not provided and unable to retrieve active project from gcloud: %v", err)
		}
		project = &p
	}

	// Perform dry run or apply update.
	if *dryRun {
		if err := config.DryRunCloudRunEnv(*project, *region, *service, envVars, *serviceAccount); err != nil {
			log.Fatalf("Dry run failed: %v", err)
		}
		fmt.Println("Dry run complete. No changes were made.")
	} else {
		if err := config.UpdateCloudRunEnv(*project, *region, *service, envVars, *serviceAccount); err != nil {
			log.Fatalf("Failed to update Cloud Run service: %v", err)
		}
		fmt.Println("Environment variables updated successfully.")
	}
}
