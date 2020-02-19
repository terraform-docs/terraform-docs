package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var yamlCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "yaml [PATH]",
	Short: "Generate YAML of inputs and outputs",
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewYAML(settings))
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
