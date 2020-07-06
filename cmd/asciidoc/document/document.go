package document

import (
	"github.com/spf13/cobra"

	"github.com/terraform-docs/terraform-docs/internal/cli"
)

// NewCommand returns a new cobra.Command for 'asciidoc document' formatter
func NewCommand(config *cli.Config) *cobra.Command {
	cmd := &cobra.Command{
		Args:        cobra.ExactArgs(1),
		Use:         "document [PATH]",
		Aliases:     []string{"doc"},
		Short:       "Generate AsciiDoc document of inputs and outputs",
		Annotations: cli.Annotations("asciidoc document"),
		PreRunE:     cli.PreRunEFunc(config),
		RunE:        cli.RunEFunc(config),
	}
	return cmd
}
