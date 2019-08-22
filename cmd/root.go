package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/segmentio/terraform-docs/internal/pkg/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
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

var settings = _settings.NewSettings()

func init() {
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "omit sorted rendering of inputs and outputs")
	rootCmd.PersistentFlags().BoolVar(&settings.SortInputsByRequired, "sort-inputs-by-required", false, "sort inputs by name and prints required inputs first")
	rootCmd.PersistentFlags().BoolVar(&settings.AggregateTypeDefaults, "with-aggregate-type-defaults", false, "print default values of aggregate types")

	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "omit \"Required\" column when generating Markdown")
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func commandsPreRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: No path provided!")
		os.Exit(1)
	}
}

func doPrint(paths []string, fn func(*doc.Doc) (string, error)) string {
	docs, err := doc.CreateFromPaths(paths)

	if err != nil {
		log.Fatal(err)
	}

	output, err := fn(docs)

	if err != nil {
		log.Fatal(err)
	}

	return output
}
