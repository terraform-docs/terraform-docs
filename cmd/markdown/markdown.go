package markdown

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/cmd/markdown/document"
	"github.com/segmentio/terraform-docs/cmd/markdown/table"
	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'markdown' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "markdown [PATH]",
		Aliases:     []string{"md"},
		Short:       "Generate Markdown of inputs and outputs",
		Annotations: cli.Annotations("markdown"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Required, "required", true, "show \"Required\" column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Sensitive, "sensitive", true, "show \"Sensitive\" column or section")
	cmd.PersistentFlags().BoolVar(&config.Settings.Escape, "escape", true, "escape special characters")
	cmd.PersistentFlags().IntVar(&config.Settings.Indent, "indent", 2, "indention level of Markdown sections [1, 2, 3, 4, 5]")

	// deprecation
	cmd.PersistentFlags().BoolVar(new(bool), "no-required", false, "do not show \"Required\" column or section")
	cmd.PersistentFlags().BoolVar(new(bool), "no-sensitive", false, "do not show \"Sensitive\" column or section")
	cmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
	cmd.PersistentFlags().MarkDeprecated("no-required", "use '--required=false' instead")   //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-sensitive", "use '--sensitive=false' instead") //nolint:errcheck
	cmd.PersistentFlags().MarkDeprecated("no-escape", "use '--escape=false' instead")       //nolint:errcheck

	// subcommands
	cmd.AddCommand(document.NewCommand(config))
	cmd.AddCommand(table.NewCommand(config))

	return cmd
}
