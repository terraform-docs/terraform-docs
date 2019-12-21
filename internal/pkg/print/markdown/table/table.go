package table

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

// Print prints a document as Markdown tables.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	module.Sort(settings)

	printVariables(&buffer, module.Variables, settings)
	printOutputs(&buffer, module.Outputs, settings)

	return markdown.Sanitize(buffer.String()), nil
}

func getVariableType(variable *tfconf.Variable) string {
	varType, _ := markdown.PrintFencedCodeBlock(variable.Type, "")
	return varType
}

func getVariableValue(variable *tfconf.Variable) string {
	var result = "n/a"

	if variable.HasDefault() {
		result, _ = markdown.PrintFencedCodeBlock(variable.Default, "")
	}
	return result
}

func printIsVariableRequired(variable *tfconf.Variable) string {
	if !variable.HasDefault() {
		return "yes"
	}
	return "no"
}

func printVariables(buffer *bytes.Buffer, variables []*tfconf.Variable, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Variables\n\n", markdown.GenerateIndentation(0, settings)))

	if len(variables) == 0 {
		buffer.WriteString("No variable.\n\n")
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
		buffer.WriteString(":--------:|\n")
	} else {
		buffer.WriteString("\n")
	}

	for _, variable := range variables {
		buffer.WriteString(
			fmt.Sprintf(
				"| %s | %s | %s | %s |",
				markdown.SanitizeName(variable.Name, settings),
				markdown.SanitizeItemForTable(variable.Description, settings),
				markdown.SanitizeItemForTable(getVariableType(variable), settings),
				markdown.SanitizeItemForTable(getVariableValue(variable), settings),
			),
		)

		if settings.ShowRequired {
			buffer.WriteString(fmt.Sprintf(" %v |\n", printIsVariableRequired(variable)))
		} else {
			buffer.WriteString("\n")
		}
	}

}

func printOutputs(buffer *bytes.Buffer, outputs []*tfconf.Output, settings *print.Settings) {
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
