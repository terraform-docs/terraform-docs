package cmd

import (
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/version"
	"github.com/segmentio/terraform-docs/pkg/print"
	"github.com/spf13/cobra"
)

var settings = print.NewSettings()
var options = module.NewOptions()

var rootCmd = &cobra.Command{
	Args:          cobra.NoArgs,
	Use:           "terraform-docs",
	Short:         "A utility to generate documentation from Terraform modules in various output formats",
	Long:          "A utility to generate documentation from Terraform modules in various output formats",
	Version:       version.Version(),
	SilenceUsage:  true,
	SilenceErrors: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		oppositeBool := func(name string) bool {
			val, _ := cmd.Flags().GetBool(name)
			return !val
		}
		settings.ShowHeader = oppositeBool("no-header")
		options.ShowHeader = settings.ShowHeader

		settings.ShowInputs = oppositeBool("no-inputs")
		settings.ShowOutputs = oppositeBool("no-outputs")
		settings.ShowProviders = oppositeBool("no-providers")
		settings.ShowRequirements = oppositeBool("no-requirements")

		settings.OutputValues = options.OutputValues

		settings.ShowColor = oppositeBool("no-color")
		settings.SortByName = oppositeBool("no-sort")
		settings.ShowRequired = oppositeBool("no-required")
		settings.EscapeCharacters = oppositeBool("no-escape")
		settings.ShowSensitivity = oppositeBool("no-sensitive")
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-header", false, "do not show module header")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-inputs", false, "do not show inputs")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-outputs", false, "do not show outputs")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-providers", false, "do not show providers")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-requirements", false, "do not show module requirements")

	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "do no sort items")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-by-required", false, "sort items by name and print required ones first")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByType, "sort-by-type", false, "sort items by type of them")

	rootCmd.PersistentFlags().StringVar(&options.HeaderFromFile, "header-from", "main.tf", "relative path of a file to read header from")

	rootCmd.PersistentFlags().BoolVar(&options.OutputValues, "output-values", false, "inject output values into outputs")
	rootCmd.PersistentFlags().StringVar(&options.OutputValuesPath, "output-values-from", "", "inject output values from file into outputs")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	return nil
}

// RootCmd represents the base command when called without any subcommands
func RootCmd() *cobra.Command {
	return rootCmd
}

var formatRunE = func(cmd *cobra.Command, args []string) error {
	name := strings.Replace(cmd.CommandPath(), "terraform-docs ", "", -1)
	printer, err := format.Factory(name, settings)
	if err != nil {
		return err
	}
	_, err = options.With(&module.Options{
		Path: args[0],
		SortBy: &module.SortBy{
			Name:     settings.SortByName,
			Required: settings.SortByRequired,
			Type:     settings.SortByType,
		},
	})
	if err != nil {
		return err
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

var formatAnnotations = func(cmd string) map[string]string {
	annotations := make(map[string]string)
	for _, s := range strings.Split(cmd, " ") {
		annotations["command"] = s
	}
	annotations["kind"] = "formatter"
	return annotations
}
