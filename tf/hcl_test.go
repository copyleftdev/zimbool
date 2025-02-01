package tf

import (
	"strings"
	"testing"
)

func TestEnvVarsToHCL(t *testing.T) {
	envVars := map[string]string{
		"FOO": "bar",
		"BAZ": "qux",
	}
	hcl := EnvVarsToHCL(envVars)
	if !strings.Contains(hcl, "locals") || !strings.Contains(hcl, "env_vars") {
		t.Errorf("HCL output does not contain expected content:\n%s", hcl)
	}
	if !strings.Contains(hcl, `"FOO" = "bar"`) || !strings.Contains(hcl, `"BAZ" = "qux"`) {
		t.Errorf("HCL output missing key/value pairs:\n%s", hcl)
	}
}
