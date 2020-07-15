package tfvars

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/cmd/tfvars/hcl"
	"github.com/terraform-docs/terraform-docs/cmd/tfvars/json"
	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'tfvars' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "tfvars [PATH]",
		Short:       "Generate terraform.tfvars of inputs",
		Annotations: cli.Annotations("tfvars"),
	}

	// subcommands
	cmd.AddCommand(hcl.NewCommand(config))
	cmd.AddCommand(json.NewCommand(config))

	return cmd
}
