package json

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'json' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "json [PATH]",
		Short:       "Generate JSON of inputs and outputs",
		Annotations: cli.Annotations("json"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Escape, "escape", true, "escape special characters")

	// deprecation
	cmd.PersistentFlags().BoolVar(&config.Settings.Deprecated.NoEscape, "no-escape", false, "do not escape special characters")
	cmd.PersistentFlags().MarkDeprecated("no-escape", "use '--escape=false' instead") //nolint:errcheck

	return cmd
}
