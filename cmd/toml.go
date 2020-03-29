package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var tomlCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "toml [PATH]",
	Short: "Generate TOML of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewTOML(settings))
	},
}

func init() {
	rootCmd.AddCommand(tomlCmd)
}
