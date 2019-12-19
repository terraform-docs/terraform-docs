package document

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
)

// Print prints a document as Markdown document.
func Print(document *doc.Doc, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	if settings.SortByName {
		if settings.SortInputsByRequired {
			doc.SortInputsByRequired(document.Inputs)
			doc.SortInputsByRequired(document.RequiredInputs)
			doc.SortInputsByRequired(document.OptionalInputs)
		} else {
			doc.SortInputsByName(document.Inputs)
			doc.SortInputsByName(document.RequiredInputs)
			doc.SortInputsByName(document.OptionalInputs)
		}
	}

	if settings.SortByName {
		doc.SortOutputsByName(document.Outputs)
	}

	printInputs(&buffer, document, settings)
	printOutputs(&buffer, document.Outputs, settings)

	out := strings.Replace(buffer.String(), "<br>```<br>", "\n```\n", -1)

	// the left over <br> or either inside or outside a code block:
	segments := strings.Split(out, "\n```\n")
	buf := bytes.NewBufferString("")
	nextIsInCodeBlock := strings.HasPrefix(out, "```\n")
	for i, segment := range segments {
		if !nextIsInCodeBlock {
			if i > 0 && len(segment) > 0 {
				buf.WriteString("\n```\n")
			}
			segment = markdown.Sanitize(segment)
			segment = strings.Replace(segment, "<br><br>", "\n\n", -1)
			segment = strings.Replace(segment, "<br>", "  \n", -1)
			buf.WriteString(segment)
			nextIsInCodeBlock = true
		} else {
			buf.WriteString("\n```\n")
			buf.WriteString(strings.Replace(segment, "<br>", "\n", -1))
			nextIsInCodeBlock = false
		}
	}
	return strings.Replace(buf.String(), " \n\n", "\n\n", -1), nil
}

func getInputDefaultValue(input *doc.Input, settings *print.Settings) string {
	var result = "n/a"

	if input.HasDefault() {
		result = markdown.PrintFencedCodeBlock(input.Default, "json")
	}

	return result
}

func printInput(buffer *bytes.Buffer, input doc.Input, settings *print.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(input.Name, settings)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.SanitizeDescriptionForDocument(input.Description, settings)))
	buffer.WriteString(fmt.Sprintf("Type: %s\n", markdown.PrintFencedCodeBlock(input.Type, "hcl")))

	// Don't print defaults for required inputs when we're already explicit about it being required
	if input.HasDefault() || !settings.ShowRequired {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s\n", getInputDefaultValue(&input, settings)))
	}
}

func printInputsRequired(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Required Inputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(inputs) == 0 {
		buffer.WriteString("No required input.\n\n")
	} else {
		buffer.WriteString("The following input variables are required:\n")

		for _, input := range inputs {
			printInput(buffer, input, settings)
		}
	}
}

func printInputsOptional(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s Optional Inputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(inputs) == 0 {
		buffer.WriteString("No optional input.\n\n")
	} else {
		buffer.WriteString("The following input variables are optional (have default values):\n")

		for _, input := range inputs {
			printInput(buffer, input, settings)
		}
	}
}

func printInputsAll(buffer *bytes.Buffer, inputs []doc.Input, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Inputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(inputs) == 0 {
		buffer.WriteString("No input.\n\n")
		return
	}

	buffer.WriteString("The following input variables are supported:\n")

	for _, input := range inputs {
		printInput(buffer, input, settings)
	}
}

func printInputs(buffer *bytes.Buffer, document *doc.Doc, settings *print.Settings) {
	if settings.ShowRequired {
		printInputsRequired(buffer, document.RequiredInputs, settings)
		printInputsOptional(buffer, document.OptionalInputs, settings)
	} else {
		printInputsAll(buffer, document.Inputs, settings)
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("\n%s Outputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(outputs) == 0 {
		buffer.WriteString("No output.\n\n")
		return
	}

	buffer.WriteString("The following outputs are exported:\n")

	for _, output := range outputs {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(output.Name, settings)))
		buffer.WriteString(fmt.Sprintf("Description: %s\n", markdown.SanitizeDescriptionForDocument(output.Description, settings)))
	}
}
