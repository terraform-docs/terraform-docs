package cmd

import (
	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/spf13/cobra"
)

var asciidocCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "asciidoc [PATH]",
	Aliases: []string{"ad"},
	Short:   "Generate AsciiDoc of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewAsciidocTable(settings))
	},
}

var asciidocTableCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "table [PATH]",
	Aliases: []string{"tbl"},
	Short:   "Generate AsciiDoc tables of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewAsciidocTable(settings))
	},
}

var asciidocDocumentCmd = &cobra.Command{
	Args:    cobra.ExactArgs(1),
	Use:     "document [PATH]",
	Aliases: []string{"doc"},
	Short:   "Generate AsciiDoc document of inputs and outputs",
	Annotations: map[string]string{
		"kind": "formatter",
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return doPrint(args[0], format.NewAsciidocDocument(settings))
	},
}

func init() {
	asciidocCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "do not show \"Required\" column or section")
	asciidocCmd.PersistentFlags().BoolVar(new(bool), "no-sensitive", false, "do not show \"Sensitive\" column or section")
	asciidocCmd.PersistentFlags().IntVar(&settings.AsciidocIndent, "indent", 2, "indention level of AsciiDoc sections [1, 2, 3, 4, 5]")

	asciidocCmd.AddCommand(asciidocTableCmd)
	asciidocCmd.AddCommand(asciidocDocumentCmd)

	rootCmd.AddCommand(asciidocCmd)
}
