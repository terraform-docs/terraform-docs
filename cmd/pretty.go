package cmd

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/spf13/cobra"
)

var prettyCmd = &cobra.Command{
	Use:    "pretty [PATH...]",
	Short:  "Generate a colorized pretty of inputs and outputs",
	PreRun: commandsPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(doPrint(args, func(docs *doc.Doc) (string, error) {
			return pretty.Print(docs, settings)
		}))
	},
}

func init() {
	rootCmd.AddCommand(prettyCmd)
}
