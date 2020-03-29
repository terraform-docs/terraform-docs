package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var prettyCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "pretty [PATH]",
	Short: "Generate colorized pretty of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewPretty(settings))
	},
}

func init() {
	prettyCmd.PersistentFlags().BoolVar(new(bool), "no-color", false, "do not colorize printed result")

	rootCmd.AddCommand(prettyCmd)
}
