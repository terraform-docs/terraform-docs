package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/format"
	"github.com/segmentio/terraform-docs/internal/module"
	"github.com/segmentio/terraform-docs/internal/version"
	"github.com/segmentio/terraform-docs/pkg/print"
)

var hides []string
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		oppositeBool := func(name string) bool {
			val, _ := cmd.Flags().GetBool(name)
			return !val
		}

		for _, h := range hides {
			switch h {
			case "header":
			case "inputs":
			case "outputs":
			case "providers":
			case "requirements":
			default:
				return fmt.Errorf("'%s' is not a valid section to hide, available options [header, inputs, outputs, providers, requirements]", h)
			}
		}

		if len(hides) > 0 && contains("header") {
			settings.ShowHeader = false
		} else {
			settings.ShowHeader = oppositeBool("no-header")
		}
		options.ShowHeader = settings.ShowHeader

		if len(hides) > 0 && contains("inputs") {
			settings.ShowInputs = false
		} else {
			settings.ShowInputs = oppositeBool("no-inputs")
		}
		if len(hides) > 0 && contains("outputs") {
			settings.ShowOutputs = false
		} else {
			settings.ShowOutputs = oppositeBool("no-outputs")
		}
		if len(hides) > 0 && contains("providers") {
			settings.ShowProviders = false
		} else {
			settings.ShowProviders = oppositeBool("no-providers")
		}
		if len(hides) > 0 && contains("requirements") {
			settings.ShowRequirements = false
		} else {
			settings.ShowRequirements = oppositeBool("no-requirements")
		}

		settings.OutputValues = options.OutputValues

		if !cmd.Flags().Changed("color") {
			settings.ShowColor = oppositeBool("no-color")
		}
		if !cmd.Flags().Changed("sort") {
			settings.SortByName = oppositeBool("no-sort")
		}
		if !cmd.Flags().Changed("required") {
			settings.ShowRequired = oppositeBool("no-required")
		}
		if !cmd.Flags().Changed("escape") {
			settings.EscapeCharacters = oppositeBool("no-escape")
		}
		if !cmd.Flags().Changed("sensitive") {
			settings.ShowSensitivity = oppositeBool("no-sensitive")
		}

		return nil
	},
}

func init() {
	rootCmd.PersistentFlags().StringSliceVar(&hides, "hide", []string{}, "hide section [header, inputs, outputs, providers, requirements]")

	rootCmd.PersistentFlags().BoolVar(new(bool), "no-header", false, "do not show module header")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-inputs", false, "do not show inputs")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-outputs", false, "do not show outputs")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-providers", false, "do not show providers")
	rootCmd.PersistentFlags().BoolVar(new(bool), "no-requirements", false, "do not show module requirements")

	rootCmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "do no sort items")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByName, "sort", true, "sort items")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByRequired, "sort-by-required", false, "sort items by name and print required ones first (default false)")
	rootCmd.PersistentFlags().BoolVar(&settings.SortByType, "sort-by-type", false, "sort items by type of them (default false)")

	rootCmd.PersistentFlags().StringVar(&options.HeaderFromFile, "header-from", "main.tf", "relative path of a file to read header from")

	rootCmd.PersistentFlags().BoolVar(&options.OutputValues, "output-values", false, "inject output values into outputs (default false)")
	rootCmd.PersistentFlags().StringVar(&options.OutputValuesPath, "output-values-from", "", "inject output values from file into outputs")

	rootCmd.PersistentFlags().MarkDeprecated("no-header", "use '--hide header' instead")             //nolint:errcheck
	rootCmd.PersistentFlags().MarkDeprecated("no-inputs", "use '--hide inputs' instead")             //nolint:errcheck
	rootCmd.PersistentFlags().MarkDeprecated("no-outputs", "use '--hide outputs' instead")           //nolint:errcheck
	rootCmd.PersistentFlags().MarkDeprecated("no-providers", "use '--hide providers' instead")       //nolint:errcheck
	rootCmd.PersistentFlags().MarkDeprecated("no-requirements", "use '--hide requirements' instead") //nolint:errcheck
	rootCmd.PersistentFlags().MarkDeprecated("no-sort", "use '--sort=false' instead")                //nolint:errcheck
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

func contains(section string) bool {
	for _, h := range hides {
		if h == section {
			return true
		}
	}
	return false
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
