package hcl

import (
	"github.com/spf13/cobra"

	"github.com/segmentio/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'tfvars hcl' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "hcl [PATH]",
		Short:       "Generate HCL format of terraform.tfvars of inputs",
		Annotations: cli.Annotations("tfvars hcl"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}
	return cmd
}
