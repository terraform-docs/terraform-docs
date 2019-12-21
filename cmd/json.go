package cmd

import (
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "json [PATH...]",
	Short: "Generate a JSON of variables and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(module *tfconf.Module) (string, error) {
			return json.Print(module, settings)
		})
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}
