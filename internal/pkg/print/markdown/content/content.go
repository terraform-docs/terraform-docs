package content

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a document as markdown content.
func Print(document *doc.Doc, settings settings.Settings) (string, error) {
	var buffer bytes.Buffer

	if document.HasComment() {
		printComment(&buffer, document.Comment, settings)
	}

	if document.HasInputs() {
		if settings.Has(print.WithSortByName) {
			if settings.Has(print.WithSortInputsByRequired) {
				doc.SortInputsByRequired(document.Inputs)
			} else {
				doc.SortInputsByName(document.Inputs)
			}
		}

		printInputs(&buffer, document.Inputs, settings)
	}

	if document.HasOutputs() {
		if settings.Has(print.WithSortByName) {
			doc.SortOutputsByName(document.Outputs)
		}

		if document.HasInputs() {
			buffer.WriteString("\n")
		}

		printOutputs(&buffer, document.Outputs, settings)
	}

	return buffer.String(), nil
}

func printComment(buffer *bytes.Buffer, comment string, settings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("%s\n", comment))
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	if settings.Has(print.WithRequired) {
		buffer.WriteString("## Required Inputs\n\n")
		buffer.WriteString("These variables must be set:\n")

		for _, input := range inputs {
			if input.IsRequired() {
				printInputMarkdown(buffer, input, settings, false)
			}
		}

		buffer.WriteString("\n")
		buffer.WriteString("## Optional Inputs\n\n")
		buffer.WriteString("These variables are optional with default values:\n")

		for _, input := range inputs {
			if !input.IsRequired() {
				printInputMarkdown(buffer, input, settings, true)
			}
		}
	} else {
		buffer.WriteString("## Inputs\n\n")
		buffer.WriteString("These variables are defined:\n")

		for _, input := range inputs {
			printInputMarkdown(buffer, input, settings, true)
		}
	}
}

func printInputMarkdown(buffer *bytes.Buffer, input doc.Input, settings settings.Settings, showDefault bool) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("### %s\n\n", strings.Replace(input.Name, "_", "\\_", -1)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", prepareDescriptionForMarkdown(getInputDescription(&input))))
	buffer.WriteString(fmt.Sprintf("Type: `%s`\n", input.Type))

	if showDefault {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s\n", getInputDefaultValue(&input, settings)))
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("## Outputs\n\n")
	buffer.WriteString("The config outputs these values:\n")

	for _, output := range outputs {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("### %s\n\n", strings.Replace(output.Name, "_", "\\_", -1)))
		buffer.WriteString(fmt.Sprintf("Description: %s\n", prepareDescriptionForMarkdown(getOutputDescription(&output))))
	}
}

func getInputDefaultValue(input *doc.Input, settings settings.Settings) string {
	var result = "-"

	if input.HasDefault() {
		result = fmt.Sprintf("`%s`", print.GetPrintableValue(input.Default, settings))
	}

	return result
}

func getInputDescription(input *doc.Input) string {
	var result = "-"

	if input.HasDescription() {
		result = input.Description
	}

	return result
}

func getOutputDescription(output *doc.Output) string {
	var result = "-"

	if output.HasDescription() {
		result = output.Description
	}

	return result
}

func prepareDescriptionForMarkdown(s string) string {
	// Convert double newlines to <br><br>.
	s = strings.Replace(
		strings.TrimSpace(s),
		"\n\n",
		"<br><br>",
		-1)

	// Convert single newline to space.
	return strings.Replace(s, "\n", " ", -1)
}
