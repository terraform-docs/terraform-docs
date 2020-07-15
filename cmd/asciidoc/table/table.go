package table

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'asciidoc table' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "table [PATH]",
		Aliases:     []string{"tbl"},
		Short:       "Generate AsciiDoc tables of inputs and outputs",
		Annotations: cli.Annotations("asciidoc table"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}
	return cmd
}
