package xml

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'xml' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "xml [PATH]",
		Short:       "Generate XML of inputs and outputs",
		Annotations: cli.Annotations("xml"),
	}

	cmd.PreRunE = cli.PreRunEFunc(config)
	cmd.RunE = cli.RunEFunc(config)

	return cmd
}
