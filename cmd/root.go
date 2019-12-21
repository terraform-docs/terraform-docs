package cmd

import (
	"fmt"
	"log"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/segmentio/terraform-docs/internal/pkg/version"
	"github.com/spf13/cobra"
)

var settings = print.NewSettings()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	Use:     "terraform-docs",
	Short:   "A utility to generate documentation from Terraform modules in various output formats",
	Long:    "A utility to generate documentation from Terraform modules in various output formats",
	Version: version.Version(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		nosort, _ := cmd.Flags().GetBool("no-sort")
		norequired, _ := cmd.Flags().GetBool("no-required")
		noescape, _ := cmd.Flags().GetBool("no-escape")

		settings.SortByName = !nosort
		settings.ShowRequired = !norequired
		settings.EscapeMarkdown = !noescape
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "omit sorted rendering of inputs and outputs")
	rootCmd.PersistentFlags().BoolVar(&settings.SortInputsByRequired, "sort-inputs-by-required", false, "sort inputs by name and prints required inputs first")
	rootCmd.PersistentFlags().BoolVar(&settings.AggregateTypeDefaults, "with-aggregate-type-defaults", false, "print default values of aggregate types")

	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "omit \"Required\" column when generating Markdown")
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
	markdownCmd.PersistentFlags().IntVar(&settings.MarkdownIndent, "indent", 2, "indention level of Markdown sections [1, 2, 3, 4, 5]")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func doPrint(paths []string, fn func(*tfconf.Module) (string, error)) {
	module, err := tfconf.CreateModule(paths[0])
	if err != nil {
		log.Fatal(err)
	}

	output, err := fn(module)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
