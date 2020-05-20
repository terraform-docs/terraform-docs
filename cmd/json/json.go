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
	}

	cmd.PreRunE = cli.PreRunEFunc(config)
	cmd.RunE = cli.RunEFunc(config)

	cmd.PersistentFlags().BoolVar(&config.Settings.Escape, "escape", true, "escape special characters")

	// deprecation
	cmd.PersistentFlags().BoolVar(new(bool), "no-escape", false, "do not escape special characters")
	cmd.PersistentFlags().MarkDeprecated("no-escape", "use '--escape=false' instead") //nolint:errcheck

	return cmd
}
