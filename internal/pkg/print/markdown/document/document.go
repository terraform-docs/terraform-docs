package document

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a document as Markdown document.
func Print(document *doc.Doc, settings *settings.Settings) (string, error) {
	var buffer bytes.Buffer

	if document.HasComment() {
		printComment(&buffer, document.Comment, settings)
	}

	if document.HasInputs() {
		if settings.SortByName {
			if settings.SortInputsByRequired {
				doc.SortInputsByRequired(document.Inputs)
			} else {
				doc.SortInputsByName(document.Inputs)
			}
		}

		printInputs(&buffer, document.Inputs, settings)
	}

	if document.HasOutputs() {
		if settings.SortByName {
			doc.SortOutputsByName(document.Outputs)
		}

		if document.HasInputs() {
			buffer.WriteString("\n")
		}

		printOutputs(&buffer, document.Outputs, settings)
	}

	return markdown.Sanitize(buffer.String()), nil
}

func getInputDefaultValue(input *doc.Input, settings *settings.Settings) string {
	var result = "n/a"

	if input.HasDefault() {
		if settings.AggregateTypeDefaults && input.IsAggregateType() {
			result = printFencedCodeBlock(print.GetPrintableValue(input.Default, settings, true))
		} else {
			result = fmt.Sprintf("`%s`", print.GetPrintableValue(input.Default, settings, false))
		}
	}

	return result
}

func printComment(buffer *bytes.Buffer, comment string, settings *settings.Settings) {
	buffer.WriteString(fmt.Sprintf("%s\n", comment))
}

func printFencedCodeBlock(code string) string {
	var buffer bytes.Buffer

	buffer.WriteString("\n\n")
	buffer.WriteString("```json\n")
	buffer.WriteString(code)
	buffer.WriteString("\n")
	buffer.WriteString("```")

	return buffer.String()
}

func printInput(buffer *bytes.Buffer, input doc.Input, settings *settings.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(input.Name, settings)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.SanitizeDescription(input.Description, settings)))
	buffer.WriteString(fmt.Sprintf("Type: `%s`\n", input.Type))

	// Don't print defaults for required inputs when we're already explicit about it being required
	if !(settings.ShowRequired && input.IsRequired()) {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s\n", getInputDefaultValue(&input, settings)))
	}
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings *settings.Settings) {
	if settings.ShowRequired {
		buffer.WriteString(fmt.Sprintf("%s Required Inputs\n\n", markdown.GenerateIndentation(0, settings)))
		buffer.WriteString("The following input variables are required:\n")

		for _, input := range inputs {
			if input.IsRequired() {
				printInput(buffer, input, settings)
			}
		}

		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("%s Optional Inputs\n\n", markdown.GenerateIndentation(0, settings)))
		buffer.WriteString("The following input variables are optional (have default values):\n")

		for _, input := range inputs {
			if !input.IsRequired() {
				printInput(buffer, input, settings)
			}
		}
	} else {
		buffer.WriteString(fmt.Sprintf("%s Inputs\n\n", markdown.GenerateIndentation(0, settings)))
		buffer.WriteString("The following input variables are supported:\n")

		for _, input := range inputs {
			printInput(buffer, input, settings)
		}
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings *settings.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Outputs\n\n", markdown.GenerateIndentation(0, settings)))
	buffer.WriteString("The following outputs are exported:\n")

	for _, output := range outputs {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(output.Name, settings)))
		buffer.WriteString(fmt.Sprintf("Description: %s\n", markdown.SanitizeDescription(output.Description, settings)))
	}
}
