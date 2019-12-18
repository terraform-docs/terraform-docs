package cmd

import (
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "markdown [PATH...]",
	Aliases: []string{"md"},
	Short:   "Generate Markdown of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(doc *doc.Doc) (string, error) {
			return table.Print(doc, settings)
		})
	},
}

var mdTableCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "table [PATH...]",
	Aliases: []string{"tbl"},
	Short:   "Generate Markdown tables of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(doc *doc.Doc) (string, error) {
			return table.Print(doc, settings)
		})
	},
}

var mdDocumentCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "document [PATH...]",
	Aliases: []string{"doc"},
	Short:   "Generate Markdown document of inputs and outputs",
	Run: func(cmd *cobra.Command, args []string) {
		doPrint(args, func(doc *doc.Doc) (string, error) {
			return document.Print(doc, settings)
		})
	},
}

func init() {
	markdownCmd.AddCommand(mdTableCmd)
	markdownCmd.AddCommand(mdDocumentCmd)

	rootCmd.AddCommand(markdownCmd)
}
