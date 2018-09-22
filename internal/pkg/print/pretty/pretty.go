package pretty

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a pretty document.
func Print(document *doc.Doc, settings settings.Settings) (string, error) {
	var buffer bytes.Buffer

	if document.HasComment() {
		printComment(&buffer, document.Comment, settings)
	}

	if document.HasInputs() {
		printInputs(&buffer, document.Inputs, settings)
	}

	if document.HasOutputs() {
		printOutputs(&buffer, document.Outputs, settings)
	}

	return buffer.String(), nil
}

func getInputDescription(input *doc.Input) string {
	result := "-"

	if input.HasDescription() {
		result = input.Description
	}

	return result
}

func getOutputDescription(output *doc.Output) string {
	result := "-"

	if output.HasDescription() {
		result = output.Description
	}

	return result
}

func printComment(buffer *bytes.Buffer, comment string, settings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("\n%s\n", comment))
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	buffer.WriteString("\n")

	for _, input := range inputs {
		format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				input.Value(),
				getInputDescription(&input)))
	}

	buffer.WriteString("\n")
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("\n")

	for _, output := range outputs {
		format := "  \033[36moutput.%s\033[0m\n  \033[90m%s\033[0m\n\n"

		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				getOutputDescription(&output)))
	}

	buffer.WriteString("\n")
}
