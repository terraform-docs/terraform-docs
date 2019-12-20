package table

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
)

// Print prints a document as Markdown tables.
func Print(document *doc.Doc, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	if settings.SortByName {
		if settings.SortInputsByRequired {
			doc.SortInputsByRequired(document.Inputs)
			// doc.SortInputsByRequired(document.RequiredInputs)
			// doc.SortInputsByRequired(document.OptionalInputs)
		} else {
			doc.SortInputsByName(document.Inputs)
			// doc.SortInputsByName(document.RequiredInputs)
			// doc.SortInputsByName(document.OptionalInputs)
		}
	}

	if settings.SortByName {
		doc.SortOutputsByName(document.Outputs)
	}

	printInputs(&buffer, document.Inputs, settings)
	printOutputs(&buffer, document.Outputs, settings)

	return markdown.Sanitize(buffer.String()), nil
}

func getInputType(input *doc.Input) string {
	inputType, _ := markdown.PrintFencedCodeBlock(input.Type, "")
	return inputType
}

func getInputValue(input *doc.Input) string {
	var result = "n/a"

	if input.HasDefault() {
		result, _ = markdown.PrintFencedCodeBlock(input.Default, "")
	}
	return result
}

func printIsInputRequired(input *doc.Input) string {
	if !input.HasDefault() {
		return "yes"
	}
	return "no"
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Inputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(inputs) == 0 {
		buffer.WriteString("No input.\n\n")
		return
	}

	buffer.WriteString("| Name | Description | Type | Default |")

	if settings.ShowRequired {
		buffer.WriteString(" Required |\n")
	} else {
		buffer.WriteString("\n")
	}

	buffer.WriteString("|------|-------------|------|---------|")

	if settings.ShowRequired {
		buffer.WriteString(":-----:|\n")
	} else {
		buffer.WriteString("\n")
	}

	for _, input := range inputs {
		buffer.WriteString(
			fmt.Sprintf(
				"| %s | %s | %s | %s |",
				markdown.SanitizeName(input.Name, settings),
				markdown.SanitizeItemForTable(input.Description, settings),
				markdown.SanitizeItemForTable(getInputType(&input), settings),
				markdown.SanitizeItemForTable(getInputValue(&input), settings),
			),
		)

		if settings.ShowRequired {
			buffer.WriteString(fmt.Sprintf(" %v |\n", printIsInputRequired(&input)))
		} else {
			buffer.WriteString("\n")
		}
	}

}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("\n%s Outputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(outputs) == 0 {
		buffer.WriteString("No output.\n\n")
		return
	}

	buffer.WriteString("| Name | Description |\n")
	buffer.WriteString("|------|-------------|\n")

	for _, output := range outputs {
		buffer.WriteString(
			fmt.Sprintf(
				"| %s | %s |\n",
				markdown.SanitizeName(output.Name, settings),
				markdown.SanitizeItemForTable(output.Description, settings),
			),
		)
	}
}
