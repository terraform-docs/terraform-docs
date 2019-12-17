package pretty

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
)

// Print prints a pretty document.
func Print(document *doc.Doc, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	if settings.Has(settings.WithProviders) {
		printProviders(&buffer, document.Providers)
	}

	printInputs(&buffer, document.Inputs, settings)
	printOutputs(&buffer, document.Outputs, settings)

	return markdown.Sanitize(buffer.String()), nil
}

func printProviders(buffer *bytes.Buffer, providers []doc.Provider) {
	buffer.WriteString("\n")

	for _, provider := range providers {
		var name = provider.Name
		if len(provider.Alias) > 0 {
			name = fmt.Sprintf("%s.%s", provider.Name, provider.Alias)
		}
		format := "  \033[36mprovider.%s\033[0m\n  \033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				name,
				provider.Version))
	}
}

func getInputDefaultValue(input *doc.Input, settings *print.Settings) string {
	var result = "required"

	if input.HasDefault() {
		result = input.Default
	}

	return result
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString("\n")

	for _, input := range inputs {
		format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(&input, settings),
				input.Description))
	}

	buffer.WriteString("\n")
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings *print.Settings) {
	buffer.WriteString("\n")

	for _, output := range outputs {
		format := "  \033[36moutput.%s\033[0m\n  \033[90m%s\033[0m\n\n"

		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				output.Description))
	}

	buffer.WriteString("\n")
}
