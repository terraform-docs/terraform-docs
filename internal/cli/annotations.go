package cli

import (
	"strings"
)

// Annotations returns set of annotations for cobra.Commands,
// specifically the 'command' namd and command 'kind'
func Annotations(cmd string) map[string]string {
	annotations := make(map[string]string)
	for _, s := range strings.Split(cmd, " ") {
		annotations["command"] = s
	}
	annotations["kind"] = "formatter"
	return annotations
}
