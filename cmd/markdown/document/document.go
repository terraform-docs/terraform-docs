package document

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'markdown document' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "document [PATH]",
		Aliases:     []string{"doc"},
		Short:       "Generate Markdown document of inputs and outputs",
		Annotations: cli.Annotations("markdown document"),
	}

	cmd.PreRunE = cli.PreRunEFunc(config)
	cmd.RunE = cli.RunEFunc(config)

	return cmd
}
