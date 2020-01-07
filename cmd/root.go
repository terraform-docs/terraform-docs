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
		noproviders, _ := cmd.Flags().GetBool("no-providers")
		noinputs, _ := cmd.Flags().GetBool("no-inputs")
		nooutputs, _ := cmd.Flags().GetBool("no-outputs")

		nocolor, _ := cmd.Flags().GetBool("no-color")
		nosort, _ := cmd.Flags().GetBool("no-sort")
		norequired, _ := cmd.Flags().GetBool("no-required")
		noescape, _ := cmd.Flags().GetBool("no-escape")

		settings.ShowProviders = !noproviders
		settings.ShowInputs = !noinputs
		settings.ShowOutputs = !nooutputs

		settings.ShowColor = !nocolor
		settings.SortByName = !nosort
		settings.ShowRequired = !norequired
		settings.EscapeCharacters = !noescape
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-providers", false, "do not show providers information")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-inputs", false, "do not show inputs information")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-outputs", false, "do not show outputs information")

	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "omit sorted rendering of inputs and outputs")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-by-required", false, "sort items by name and print required ones first")

	//-----------------------------
	// deprecated - will be removed
	//-----------------------------
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-inputs-by-required", false, "[deprecated] use '--sort-by-required' instead")
	rootCmd.PersistentFlags().BoolVar(new(bool), "with-aggregate-type-defaults", false, "[deprecated] print default values of aggregate types")
	//-----------------------------

	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "omit \"Required\" column when generating Markdown")
	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
	markdownCmd.PersistentFlags().IntVar(&settings.MarkdownIndent, "indent", 2, "indention level of Markdown sections [1, 2, 3, 4, 5]")

	prettyCmd.PersistentFlags().BoolVar(new(bool), "no-color", false, "do not colorize printed result")

	jsonCmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
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
