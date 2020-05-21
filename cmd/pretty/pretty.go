package pretty

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for pretty formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "pretty [PATH]",
		Short:       "Generate colorized pretty of inputs and outputs",
		Annotations: cli.Annotations("pretty"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}

	// flags
	cmd.PersistentFlags().BoolVar(&config.Settings.Color, "color", true, "colorize printed result")

	// deprecation
	cmd.PersistentFlags().BoolVar(new(bool), "no-color", false, "do not colorize printed result")
	cmd.PersistentFlags().MarkDeprecated("no-color", "use '--color=false' instead") //nolint:errcheck

	return cmd
}
