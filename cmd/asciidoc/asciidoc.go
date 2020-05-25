package asciidoc

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/cmd/asciidoc/document"
	"github.com/segmentio/terraform-docs/cmd/asciidoc/table"
	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'asciidoc' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "asciidoc [PATH]",
		Aliases:     []string{"adoc"},
		Short:       "Generate AsciiDoc of inputs and outputs",
		Annotations: cli.Annotations("asciidoc"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Required, "required", true, "show Required column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Sensitive, "sensitive", true, "show Sensitive column or section")
	cmd.PersistentFlags().IntVar(&config.Settings.Indent, "indent", 2, "indention level of AsciiDoc sections [1, 2, 3, 4, 5]")

	// deprecation
	cmd.PersistentFlags().BoolVar(&config.Settings.Deprecated.NoRequired, "no-required", false, "do not show \"Required\" column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Deprecated.NoSensitive, "no-sensitive", false, "do not show \"Sensitive\" column or section")
	cmd.PersistentFlags().MarkDeprecated("no-required", "use '--required=false' instead")   //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-sensitive", "use '--sensitive=false' instead") //nolint:errcheck

	// subcommands
	cmd.AddCommand(document.NewCommand(config))
	cmd.AddCommand(table.NewCommand(config))

	return cmd
}
