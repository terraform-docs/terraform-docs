package cmd

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/spf13/cobra"
)

var jsonCmd = &cobra.Command{
	Use:    "json [PATH...]",
	Short:  "Generate a JSON of inputs and outputs",
	PreRun: commandsPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(doPrint(args, func(docs *doc.Doc) (string, error) {
			return json.Print(docs, *settings)
		}))
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)
}
