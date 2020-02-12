package cmd

import (
	"fmt"

	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/version"
	"github.com/segmentio/terraform-docs/pkg/print"
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

	//-----------------------------
	// deprecated - will be removed
	//-----------------------------
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-inputs-by-required", false, "[deprecated] use '--sort-by-required' instead")
	rootCmd.PersistentFlags().BoolVar(new(bool), "with-aggregate-type-defaults", false, "[deprecated] print default values of aggregate types")
	//-----------------------------
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

func doPrint(path string, printer print.Format) error {
	options := &module.Options{
		Path: path,
		SortBy: &module.SortBy{
			Name:     settings.SortByName,
			Required: settings.SortByRequired,
		},
	}
	tfmodule, err := module.LoadWithOptions(options)
	if err != nil {
		return err
	}
	output, err := printer.Print(tfmodule, settings)
	if err != nil {
		return err
	}
	fmt.Println(output)
	return nil
}
