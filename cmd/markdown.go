package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "markdown [PATH]",
	Aliases: []string{"md"},
	Short:   "Generate Markdown of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewTable(settings))
	},
}

var markdownTableCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "table [PATH]",
	Aliases: []string{"tbl"},
	Short:   "Generate Markdown tables of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewTable(settings))
	},
}

var markdownDocumentCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "document [PATH]",
	Aliases: []string{"doc"},
	Short:   "Generate Markdown document of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewDocument(settings))
	},
}

func init() {
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "do not show \"Required\" column or section")
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-sensitive", false, "do not show \"Sensitive\" column or section")
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
	markdownCmd.PersistentFlags().IntVar(&settings.MarkdownIndent, "indent", 2, "indention level of Markdown sections [1, 2, 3, 4, 5]")

	markdownCmd.AddCommand(markdownTableCmd)
	markdownCmd.AddCommand(markdownDocumentCmd)

	rootCmd.AddCommand(markdownCmd)
}
