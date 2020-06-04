package cli

// Annotations returns set of annotations for cobra.Commands,
// specifically the command 'name' and command 'kind'
func Annotations(cmd string) map[string]string {
	return map[string]string{
		"command": cmd,
		"kind":    "formatter",
	}
}
