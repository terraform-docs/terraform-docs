package document

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// Print prints a document as Markdown document.
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

	return markdown.Sanitize(buffer.String()), nil
}

func getInputDefaultValue(input *doc.Input, settings settings.Settings) string {
	var result = "n/a"

	if input.HasDefault() {
		if settings.Has(print.WithAggregateTypeDefaults) && input.IsAggregateType() {
			result = printFencedCodeBlock(print.GetPrintableValue(input.Default, settings, true))
		} else {
			result = fmt.Sprintf("`%s`", print.GetPrintableValue(input.Default, settings, false))
		}
	}

	return result
}

func printComment(buffer *bytes.Buffer, comment string, settings settings.Settings) {
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

func printInput(buffer *bytes.Buffer, input doc.Input, settings settings.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("### %s\n\n", strings.Replace(input.Name, "_", "\\_", -1)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.ConvertMultiLineText(input.Description)))
	buffer.WriteString(fmt.Sprintf("Type: `%s`\n", input.Type))

	// Don't print defaults for required inputs when we're already explicit about it being required
	if !(settings.Has(print.WithRequired) && input.IsRequired()) {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s\n", getInputDefaultValue(&input, settings)))
	}
}

func printInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	if settings.Has(print.WithRequired) {
		buffer.WriteString("## Required Inputs\n\n")
		buffer.WriteString("The following input variables are required:\n")

		for _, input := range inputs {
			if input.IsRequired() {
				printInput(buffer, input, settings)
			}
		}

		buffer.WriteString("\n")
		buffer.WriteString("## Optional Inputs\n\n")
		buffer.WriteString("The following input variables are optional (have default values):\n")

		for _, input := range inputs {
			if !input.IsRequired() {
				printInput(buffer, input, settings)
			}
		}
	} else {
		buffer.WriteString("## Inputs\n\n")
		buffer.WriteString("The following input variables are supported:\n")

		for _, input := range inputs {
			printInput(buffer, input, settings)
		}
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("## Outputs\n\n")
	buffer.WriteString("The following outputs are exported:\n")

	for _, output := range outputs {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("### %s\n\n", strings.Replace(output.Name, "_", "\\_", -1)))
		buffer.WriteString(fmt.Sprintf("Description: %s\n", markdown.ConvertMultiLineText(output.Description)))
	}
}
