package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/cmd/asciidoc"
	"github.com/segmentio/terraform-docs/cmd/completion"
	"github.com/segmentio/terraform-docs/cmd/json"
	"github.com/segmentio/terraform-docs/cmd/markdown"
	"github.com/segmentio/terraform-docs/cmd/pretty"
	"github.com/segmentio/terraform-docs/cmd/tfvars"
	"github.com/segmentio/terraform-docs/cmd/toml"
	"github.com/segmentio/terraform-docs/cmd/version"
	"github.com/segmentio/terraform-docs/cmd/xml"
	"github.com/segmentio/terraform-docs/cmd/yaml"
	"github.com/segmentio/terraform-docs/internal/cli"
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	if err := NewCommand().Execute(); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return err
	}
	return nil
}

// NewCommand returns a new cobra.Command for 'root' command
func NewCommand() *cobra.Command {
	config := cli.DefaultConfig()
	cmd := &cobra.Command{
		Args:          cobra.NoArgs,
		Use:           "terraform-docs",
		Short:         "A utility to generate documentation from Terraform modules in various output formats",
		Long:          "A utility to generate documentation from Terraform modules in various output formats",
		Version:       version.Full(),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	// flags
	cmd.PersistentFlags().StringSliceVar(&config.Sections.Hide, "hide", []string{}, "hide section [header, inputs, outputs, providers, requirements]")

	cmd.PersistentFlags().BoolVar(&config.Sort.Enabled, "sort", true, "sort items")
	cmd.PersistentFlags().BoolVar(&config.Sort.By.Required, "sort-by-required", false, "sort items by name and print required ones first (default false)")
	cmd.PersistentFlags().BoolVar(&config.Sort.By.Type, "sort-by-type", false, "sort items by type of them (default false)")

	cmd.PersistentFlags().StringVar(&config.HeaderFrom, "header-from", "main.tf", "relative path of a file to read header from")

	cmd.PersistentFlags().BoolVar(&config.OutputValues.Enabled, "output-values", false, "inject output values into outputs (default false)")
	cmd.PersistentFlags().StringVar(&config.OutputValues.From, "output-values-from", "", "inject output values from file into outputs (default \"\")")

	// deprecation
	cmd.PersistentFlags().BoolVar(new(bool), "no-header", false, "do not show module header")
	cmd.PersistentFlags().BoolVar(new(bool), "no-inputs", false, "do not show inputs")
	cmd.PersistentFlags().BoolVar(new(bool), "no-outputs", false, "do not show outputs")
	cmd.PersistentFlags().BoolVar(new(bool), "no-providers", false, "do not show providers")
	cmd.PersistentFlags().BoolVar(new(bool), "no-requirements", false, "do not show module requirements")
	cmd.PersistentFlags().BoolVar(new(bool), "no-sort", false, "do no sort items")

	cmd.PersistentFlags().MarkDeprecated("no-header", "use '--hide header' instead")             //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-inputs", "use '--hide inputs' instead")             //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-outputs", "use '--hide outputs' instead")           //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-providers", "use '--hide providers' instead")       //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-requirements", "use '--hide requirements' instead") //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-sort", "use '--sort=false' instead")                //nolint:errcheck

	// formatter subcommands
	cmd.AddCommand(asciidoc.NewCommand(config))
	cmd.AddCommand(json.NewCommand(config))
	cmd.AddCommand(markdown.NewCommand(config))
	cmd.AddCommand(pretty.NewCommand(config))
	cmd.AddCommand(tfvars.NewCommand(config))
	cmd.AddCommand(toml.NewCommand(config))
	cmd.AddCommand(xml.NewCommand(config))
	cmd.AddCommand(yaml.NewCommand(config))

	// other subcommands
	cmd.AddCommand(completion.NewCommand())
	cmd.AddCommand(version.NewCommand())

	return cmd
}
