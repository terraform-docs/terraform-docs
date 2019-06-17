package document

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a document as Markdown document.
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
	buffer.WriteString("## Providers\n\n")

	if len(providers) == 0 {
		buffer.WriteString("None\n\n")
	} else {
		buffer.WriteString("The following providers are used by this module:\n\n")
		for _, provider := range providers {
			var name= provider.Name
			if len(provider.Alias) > 0 {
				name = fmt.Sprintf("%s.%s", provider.Name, provider.Alias)
			}
			name = strings.ReplaceAll(name, "_", "\\_")
			var version = ""
			if len(provider.Version) > 0 {
				version = fmt.Sprintf(" (%s)", provider.Version)
			}

			buffer.WriteString(fmt.Sprintf("* %s%s\n", name, version))
		}

		buffer.WriteString("\n")
	}
}


func printVariable(buffer *bytes.Buffer, variable doc.Variable, printSettings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("#### %s\n\n", strings.ReplaceAll(variable.Name, "_", "\\_")))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.ConvertMultiLineText(variable.Description)))
	buffer.WriteString(fmt.Sprintf("Type:\n%s\n\n", markdown.PrintCode(variable.Type, "hcl")))

	// Don't print defaults for required variables when we're already explicit about it being required
	if variable.HasDefault() {
		buffer.WriteString(fmt.Sprintf("Default:\n%s\n\n", markdown.PrintCode(variable.Default, "json")))
	} else if !(printSettings.Has(settings.WithRequired)) {
		buffer.WriteString("Default: n/a\n\n")
	}
}

func printVariables(buffer *bytes.Buffer, variables []doc.Variable, printSettings settings.Settings) {
	buffer.WriteString("## Variables\n\n")

	if len(variables) == 0 {
		buffer.WriteString("None\n\n")
	}
	if printSettings.Has(settings.WithRequired) {
		buffer.WriteString("### Required Variables\n\n")
		buffer.WriteString("The following variables are required:\n\n")

		for _, variable := range variables {
			if !variable.HasDefault() {
				printVariable(buffer, variable, printSettings)
			}
		}

		buffer.WriteString("### Optional Variables\n\n")
		buffer.WriteString("The following variables are optional (have default values):\n\n")

		for _, variable := range variables {
			if variable.HasDefault() {
				printVariable(buffer, variable, printSettings)
			}
		}
	} else {
		buffer.WriteString("The following variables are supported:\n\n")

		for _, variable := range variables {
			printVariable(buffer, variable, printSettings)
		}
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, printSettings settings.Settings) {
	buffer.WriteString("## Outputs\n\n")

	if len(outputs) == 0 {
		buffer.WriteString("None\n\n")
	} else {
		buffer.WriteString("The following outputs are exported:\n\n")

		for _, output := range outputs {
			buffer.WriteString(fmt.Sprintf("#### %s\n\n", strings.Replace(output.Name, "_", "\\_", -1)))
			buffer.WriteString(fmt.Sprintf("Description: %s\n", markdown.ConvertMultiLineText(output.Description)))
			buffer.WriteString("\n")
		}
	}
}
