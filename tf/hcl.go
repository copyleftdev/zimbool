package tf

import (
	"bytes"
	"fmt"
	"sort"
)

// EnvVarsToHCL transforms a map of environment variables into a Terraform locals block.
func EnvVarsToHCL(envVars map[string]string) string {
	var buf bytes.Buffer
	buf.WriteString("locals {\n")
	buf.WriteString("  env_vars = {\n")
	keys := make([]string, 0, len(envVars))
	for k := range envVars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := envVars[k]
		buf.WriteString(fmt.Sprintf("    %q = %q\n", k, v))
	}
	buf.WriteString("  }\n")
	buf.WriteString("}\n")
	return buf.String()
}
