package cmd

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:     "markdown [PATH...]",
	Aliases: []string{"md"},
	Short:   "Generate Markdown of inputs and outputs",
	PreRun:  commandsPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(doPrint(args, func(docs *doc.Doc) (string, error) {
			return table.Print(docs, *settings)
		}))
	},
}

var mdTableCmd = &cobra.Command{
	Use:     "table [PATH...]",
	Aliases: []string{"tbl"},
	Short:   "Generate Markdown tables of inputs and outputs",
	PreRun:  commandsPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(doPrint(args, func(docs *doc.Doc) (string, error) {
			return table.Print(docs, *settings)
		}))
	},
}

var mdDocumentCmd = &cobra.Command{
	Use:     "document [PATH...]",
	Aliases: []string{"doc"},
	Short:   "Generate Markdown document of inputs and outputs",
	PreRun:  commandsPreRun,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(doPrint(args, func(docs *doc.Doc) (string, error) {
			return document.Print(docs, *settings)
		}))
	},
}

func init() {
	markdownCmd.AddCommand(mdTableCmd)
	markdownCmd.AddCommand(mdDocumentCmd)

	rootCmd.AddCommand(markdownCmd)
}
