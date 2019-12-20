package cmd

import (
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/spf13/cobra"
)

var prettyCmd = &cobra.Command{
	Args:  cobra.ExactArgs(1),
	Use:   "pretty [PATH...]",
	Short: "Generate a colorized pretty of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(doc *doc.Doc) (string, error) {
			return pretty.Print(doc, settings)
		})
	},
}

func init() {
	rootCmd.AddCommand(prettyCmd)
}
