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
var options = tfconf.Options{}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Args:    cobra.NoArgs,
	Use:     "terraform-docs",
	Short:   "A utility to generate documentation from Terraform modules in various output formats",
	Long:    "A utility to generate documentation from Terraform modules in various output formats",
	Version: version.Version(),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		noheader, _ := cmd.Flags().GetBool("no-header")
		noproviders, _ := cmd.Flags().GetBool("no-providers")
		noinputs, _ := cmd.Flags().GetBool("no-inputs")
		nooutputs, _ := cmd.Flags().GetBool("no-outputs")

		nocolor, _ := cmd.Flags().GetBool("no-color")
		nosort, _ := cmd.Flags().GetBool("no-sort")
		norequired, _ := cmd.Flags().GetBool("no-required")
		noescape, _ := cmd.Flags().GetBool("no-escape")

		settings.ShowHeader = !noheader
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
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-header", false, "do not show module header")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-providers", false, "do not show providers")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-inputs", false, "do not show inputs")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-outputs", false, "do not show outputs")

	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "do no sort items")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-by-required", false, "sort items by name and print required ones first")

	rootCmd.PersistentFlags().BoolVar(&options.OutputValues, "output-values", false, "inject output values into outputs")
	rootCmd.PersistentFlags().StringVar(&options.OutputValuesPath, "output-values-from", "terraform-outputs.json", "inject output values from file into outputs")

	//-----------------------------
	// deprecated - will be removed
	//-----------------------------
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-inputs-by-required", false, "[deprecated] use '--sort-by-required' instead")
	rootCmd.PersistentFlags().BoolVar(new(bool), "with-aggregate-type-defaults", false, "[deprecated] print default values of aggregate types")
	//-----------------------------

	markdownCmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "do not show \"Required\" column or section")
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

// FormatterCmds returns list of available formatter
// commands (e.g. markdown, json, yaml) and ignores
// the other commands (e.g. completion, version, help)
func FormatterCmds() []*cobra.Command {
	return []*cobra.Command{
		jsonCmd,
		prettyCmd,
		mdDocumentCmd,
		mdTableCmd,
		yamlCmd,
	}
}

func doPrint(path string, fn func(*tfconf.Module) (string, error)) {
	options.Path = path
	module, err := tfconf.CreateModule(&options)
	if err != nil {
		log.Fatal(err)
	}

	output, err := fn(module)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}
