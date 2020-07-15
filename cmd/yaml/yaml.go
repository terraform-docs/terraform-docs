package yaml

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'yaml' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "yaml [PATH]",
		Short:       "Generate YAML of inputs and outputs",
		Annotations: cli.Annotations("yaml"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}
	return cmd
}
