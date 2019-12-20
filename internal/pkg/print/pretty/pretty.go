package pretty

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
)

// Print prints a pretty document.
func Print(document *doc.Doc, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	if settings.SortByName {
		if settings.SortInputsByRequired {
			doc.SortInputsByRequired(document.Inputs)
		} else {
			doc.SortInputsByName(document.Inputs)
		}
	}

	if settings.SortByName {
		doc.SortOutputsByName(document.Outputs)
	}

	printInputs(&buffer, document.Inputs, settings)
	printOutputs(&buffer, document.Outputs, settings)

	return buffer.String(), nil
}

func getInputDefaultValue(input *doc.Input, settings *print.Settings) string {
	var result = "required"

	if input.HasDefault() {
		result = input.Default
	}

	return result
}

func getDescription(description string) string {
	var result = "n/a"

	if description != "" {
		result = strings.TrimSuffix(description, "\n")
	}

	return result
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString("\n\n")

	for _, input := range inputs {
		format := "\033[36minput.%s\033[0m (%s)\n\033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(&input, settings),
				getDescription(input.Description),
			),
		)
	}

	buffer.WriteString("\n")
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings *print.Settings) {
	buffer.WriteString("\n")

	for _, output := range outputs {
		format := "\033[36moutput.%s\033[0m\n\033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				getDescription(output.Description),
			),
		)
	}
}
