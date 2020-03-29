package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var xmlCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "xml [PATH]",
	Short: "Generate XML of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewXML(settings))
	},
}

func init() {
	rootCmd.AddCommand(xmlCmd)
}
