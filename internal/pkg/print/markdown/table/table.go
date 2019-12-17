package table

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a document as Markdown tables.
func Print(document *doc.Doc, settings *settings.Settings) (string, error) {
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

		buffer.WriteString("| Name | Alias | Version |\n")
		buffer.WriteString("|------|-------|---------|\n")

		for _, provider := range providers {
			buffer.WriteString(fmt.Sprintf("| %s | %s | %s |\n",
				strings.ReplaceAll(provider.Name, "_", "\\_"),
				strings.ReplaceAll(provider.Alias, "_", "\\_"),
				provider.Version))
		}

		buffer.WriteString("\n")
	}
}

func getInputDefaultValue(input *doc.Input) string {
	var result = "n/a"

	if input.HasDefault() {
		result = markdown.PrintCode(input.Default, "json")
	}

	return result
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, printSettings settings.Settings) {
	buffer.WriteString("## Inputs\n\n")

	if len(inputs) == 0 {
		buffer.WriteString("None\n\n")
	} else {
		buffer.WriteString("<table>\n")
		buffer.WriteString("<tr><th>Name</th><th>Description</th><th>Type</th><th>Default</th>")

		if printSettings.Has(settings.WithRequired) {
			buffer.WriteString(" <th>Required</th></tr>\n")
		} else {
			buffer.WriteString("</tr>\n")
		}

		for _, input := range inputs {
			buffer.WriteString("<tr>\n")
			buffer.WriteString(fmt.Sprintf("<td>%s</td>\n", input.Name))
			buffer.WriteString(fmt.Sprintf("<td>%s</td>\n", markdown.ConvertMultiLineText(input.Description)))
			buffer.WriteString(fmt.Sprintf("<td>\n\n%s</td>\n", markdown.PrintCode(input.Type, "hcl")))
			buffer.WriteString(fmt.Sprintf("<td>\n\n%s</td>\n", getInputDefaultValue(&input)))
			if printSettings.Has(settings.WithRequired) {
				buffer.WriteString(fmt.Sprintf("<td>%s</td>\n", printIsInputRequired(&input)))
			}
			buffer.WriteString("</tr>\n")
		}
		buffer.WriteString("</table>\n\n")
	}
}

func printIsInputRequired(input *doc.Input) string {
	if !input.HasDefault() {
		return "yes"
	}

	return "no"
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("## Outputs\n\n")

	if len(outputs) == 0 {
		buffer.WriteString("None\n\n")
	} else {
		buffer.WriteString("| Name | Description |\n")
		buffer.WriteString("|------|-------------|\n")

		for _, output := range outputs {
			buffer.WriteString(
				fmt.Sprintf("| %s | %s |\n",
					strings.Replace(output.Name, "_", "\\_", -1),
					markdown.ConvertMultiLineText(output.Description)))
		}
		buffer.WriteString("\n")
	}
}
