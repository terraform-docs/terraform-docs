package pretty

import (
	"bytes"
	"fmt"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a pretty document.
func Print(document *doc.Doc, printSettings settings.Settings) (string, error) {
	var buffer bytes.Buffer

	if printSettings.Has(settings.WithProviders) {
		printProviders(&buffer, document.Providers)
	}

	printVariables(&buffer, document.Variables, printSettings)
	printOutputs(&buffer, document.Outputs, printSettings)

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

func getVariableDefaultValue(variable *doc.Variable, printSettings settings.Settings) string {
	var result = "required"

	if variable.HasDefault() {
		result = variable.Default
	}

	return result
}

func printVariables(buffer *bytes.Buffer, variables []doc.Variable, printSettings settings.Settings) {
	buffer.WriteString("\n")

	for _, variable := range variables {
		format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				variable.Name,
				getVariableDefaultValue(&variable, printSettings),
				variable.Description))
	}

	buffer.WriteString("\n")
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, printSettings settings.Settings) {
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
