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

	printInputs(&buffer, document.Inputs, printSettings)
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


func printInput(buffer *bytes.Buffer, input doc.Input, printSettings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("#### %s\n\n", strings.ReplaceAll(input.Name, "_", "\\_")))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.ConvertMultiLineText(input.Description)))
	buffer.WriteString(fmt.Sprintf("Type:\n%s\n\n", markdown.PrintCode(input.Type, "hcl")))

	// Don't print defaults for required variables when we're already explicit about it being required
	if input.HasDefault() {
		buffer.WriteString(fmt.Sprintf("Default:\n%s\n\n", markdown.PrintCode(input.Default, "json")))
	} else if !(printSettings.Has(settings.WithRequired)) {
		buffer.WriteString("Default: n/a\n\n")
	}
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, printSettings settings.Settings) {
	buffer.WriteString("## Inputs\n\n")

	if len(inputs) == 0 {
		buffer.WriteString("None\n\n")
	}
	if printSettings.Has(settings.WithRequired) {
		buffer.WriteString("### Required Inputs\n\n")
		buffer.WriteString("The following input variables are required:\n\n")

		for _, input := range inputs {
			if !input.HasDefault() {
				printInput(buffer, input, printSettings)
			}
		}

		buffer.WriteString("### Optional Inputs\n\n")
		buffer.WriteString("The following input variables are optional (have default values):\n\n")

		for _, input := range inputs {
			if input.HasDefault() {
				printInput(buffer, input, printSettings)
			}
		}
	} else {
		buffer.WriteString("The following input variables are supported:\n\n")

		for _, input := range inputs {
			printInput(buffer, input, printSettings)
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
