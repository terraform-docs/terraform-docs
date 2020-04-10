package cmd

import (
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "yaml [PATH]",
	Short:       "Generate YAML of inputs and outputs",
	Annotations: formatAnnotations("yaml"),
	RunE:        formatRunE,
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
