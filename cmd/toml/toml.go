package toml

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'toml' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "toml [PATH]",
		Short:       "Generate TOML of inputs and outputs",
		Annotations: cli.Annotations("toml"),
	}

	cmd.PreRunE = cli.PreRunEFunc(config)
	cmd.RunE = cli.RunEFunc(config)

	return cmd
}
