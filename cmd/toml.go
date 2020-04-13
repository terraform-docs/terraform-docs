package cmd

import (
	"github.com/spf13/cobra"
)

var tomlCmd = &cobra.Command{
	Args:        cobra.ExactArgs(1),
	Use:         "toml [PATH]",
	Short:       "Generate TOML of inputs and outputs",
	Annotations: formatAnnotations("toml"),
	RunE:        formatRunE,
}

func init() {
	rootCmd.AddCommand(tomlCmd)
}
