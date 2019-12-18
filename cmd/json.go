package cmd

import (
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "json [PATH...]",
	Short: "Generate a JSON of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(doc *doc.Doc) (string, error) {
			return json.Print(doc, settings)
		})
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}
