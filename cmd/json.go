package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "json [PATH]",
	Short: "Generate JSON of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewJSON(settings))
	},
}

func init() {
	jsonCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")

	rootCmd.AddCommand(jsonCmd)
}
