package document

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

// Print prints a document as Markdown document.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	module.Sort(settings)

	printVariables(&buffer, module, settings)
	printOutputs(&buffer, module.Outputs, settings)

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

func getVariableType(variable *tfconf.Variable) string {
	var result = ""
	var extraline = false

	if result, extraline = markdown.PrintFencedCodeBlock(variable.Type, "hcl"); !extraline {
		result += "\n"
	}
	return result
}

func getVariableValue(variable *tfconf.Variable) string {
	var result = "n/a\n"
	var extraline = false

	if variable.HasDefault() {
		if result, extraline = markdown.PrintFencedCodeBlock(variable.Default, "json"); !extraline {
			result += "\n"
		}
	}
	return result
}

func printVariable(buffer *bytes.Buffer, variable *tfconf.Variable, settings *print.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(variable.Name, settings)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.SanitizeItemForDocument(variable.Description, settings)))
	buffer.WriteString(fmt.Sprintf("Type: %s", getVariableType(variable)))

	// Don't print defaults for required variables when we're already explicit about it being required
	if variable.HasDefault() || !settings.ShowRequired {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s", getVariableValue(variable)))
	}
}

func printVariablesRequired(buffer *bytes.Buffer, variables []*tfconf.Variable, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Required Variables\n\n", markdown.GenerateIndentation(0, settings)))

	if len(variables) == 0 {
		buffer.WriteString("No required variable.\n\n")
	} else {
		buffer.WriteString("The following variables are required:\n")

		for _, variable := range variables {
			printVariable(buffer, variable, settings)
		}
	}
}

func printVariablesOptional(buffer *bytes.Buffer, variables []*tfconf.Variable, settings *print.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s Optional Variables\n\n", markdown.GenerateIndentation(0, settings)))

	if len(variables) == 0 {
		buffer.WriteString("No optional variable.\n\n")
	} else {
		buffer.WriteString("The following variables are optional (have default values):\n")

		for _, variable := range variables {
			printVariable(buffer, variable, settings)
		}
	}
}

func printVariablesAll(buffer *bytes.Buffer, variables []*tfconf.Variable, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Variables\n\n", markdown.GenerateIndentation(0, settings)))

	if len(variables) == 0 {
		buffer.WriteString("No variable.\n\n")
		return
	}

	buffer.WriteString("The following variables are supported:\n")

	for _, variable := range variables {
		printVariable(buffer, variable, settings)
	}
}

func printVariables(buffer *bytes.Buffer, module *tfconf.Module, settings *print.Settings) {
	if settings.ShowRequired {
		printVariablesRequired(buffer, module.RequiredVariables, settings)
		printVariablesOptional(buffer, module.OptionalVariables, settings)
	} else {
		printVariablesAll(buffer, module.Variables, settings)
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []*tfconf.Output, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("\n%s Outputs\n\n", markdown.GenerateIndentation(0, settings)))

	if len(outputs) == 0 {
		buffer.WriteString("No output.\n\n")
		return
	}

	buffer.WriteString("The following outputs are exported:\n")

	for _, output := range outputs {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(output.Name, settings)))
		buffer.WriteString(fmt.Sprintf("Description: %s\n", markdown.SanitizeItemForDocument(output.Description, settings)))
	}
}
