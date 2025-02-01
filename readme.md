
# Zimbool

**Zimbool** is a sleek, command‑line tool that pushes environment variables from a local `.env` file to a Google Cloud Run service—or transforms them into valid HCL (Terraform) configuration. It’s built for both CI/CD pipelines and daily developer use, wrapping the gcloud CLI with advanced preflight checks and dry-run support.

---

## Features

- **Preflight Checks** – Emoji‑rich system verification for required components.
- **Gcloud Integration** – Uses gcloud CLI to update Cloud Run services.
- **Dry Run Mode** – Simulate updates without applying changes.
- **HCL Transformation** – Convert a `.env` file into a Terraform locals block.
- **Concise & Elegant** – Minimalistic design and clear output.

---

## Installation

Install Zimbool using Go (ensure your `$GOBIN` is in your PATH or it defaults to `$(HOME)/.local/bin`):

```bash
go install github.com/copyleftdev/zimbool@latest
```

---

## Usage

```bash
zimbool --env-file <path-to-.env> --service <service-name> --service-account <path-to-sa.json> [--project <project-id>] [--region <region>] [--dry-run] [--to-hcl]
```

### Flags

| **Flag**             | **Description**                                                           | **Default**      | **Required**                   |
|----------------------|---------------------------------------------------------------------------|------------------|--------------------------------|
| `--env-file`         | Path to the `.env` file                                                   | `.env`           | No                             |
| `--service`          | Cloud Run service name (for updating the service)                         | —                | Yes (for update mode)          |
| `--service-account`  | Path to the service account JSON key file                                 | —                | Yes (for update mode)          |
| `--project`          | GCP Project ID (if omitted, retrieved from gcloud configuration)          | _auto-detected_  | No                             |
| `--region`           | GCP region (e.g., `us-central1`)                                            | `us-central1`    | No                             |
| `--dry-run`          | Simulate the update and print the gcloud command without making changes     | `false`          | No                             |
| `--to-hcl`           | Convert the `.env` file into a valid HCL (Terraform) locals block           | `false`          | No                             |

---

## Examples

### 1. Transform `.env` to HCL

Convert your `.env` file into a Terraform locals block:

```bash
zimbool --env-file .env --to-hcl
```

*Example output:*

```hcl
locals {
  env_vars = {
    "BAZ" = "qux"
    "FOO" = "bar"
  }
}
```

---

### 2. Dry Run Update

Simulate updating your Cloud Run service without applying changes:

```bash
zimbool --env-file .env \
        --service my-service \
        --service-account ~/keys/sa.json \
        --dry-run
```

*The tool prints the gcloud command that would be executed.*

---

### 3. Update Cloud Run Environment Variables

Push environment variables to your Cloud Run service:

```bash
zimbool --env-file .env \
        --service my-service \
        --service-account ~/keys/sa.json \
        --project my-project \
        --region us-central1
```

If the `--project` flag is omitted, Zimbool retrieves the active project from your gcloud configuration.

---

## Makefile

Zimbool includes a Makefile for common tasks:

| **Target** | **Description**                                |
|------------|------------------------------------------------|
| `make build`   | Build the binary                           |
| `make install` | Install the binary in your local bin       |
| `make test`    | Run the test suite                         |
| `make clean`   | Remove the built binary                    |
| `make fmt`     | Format the code                            |
| `make vet`     | Run static analysis (vet)                  |

---

## License

MIT License

---

**Maintained by [copyleftdev](https://github.com/copyleftdev)**

Enjoy Zimbool – where environment variables meet the cloud with style!
