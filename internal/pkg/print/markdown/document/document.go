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

	printProviders(&buffer, module.Providers, settings)
	printInputs(&buffer, module, settings)
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

func getProviderVersion(provider *tfconf.Provider) string {
	var result = ""
	if provider.Version != "" {
		result = fmt.Sprintf(" (%s)", provider.Version)
	}
	return result
}

func getInputType(input *tfconf.Input) string {
	var result = ""
	var extraline = false

	if result, extraline = markdown.PrintFencedCodeBlock(input.Type, "hcl"); !extraline {
		result += "\n"
	}
	return result
}

func getInputValue(input *tfconf.Input) string {
	var result = "n/a\n"
	var extraline = false

	if input.HasDefault() {
		if result, extraline = markdown.PrintFencedCodeBlock(input.Default, "json"); !extraline {
			result += "\n"
		}
	}
	return result
}

func printInput(buffer *bytes.Buffer, input *tfconf.Input, settings *print.Settings) {
	buffer.WriteString("\n")
	buffer.WriteString(fmt.Sprintf("%s %s\n\n", markdown.GenerateIndentation(1, settings), markdown.SanitizeName(input.Name, settings)))
	buffer.WriteString(fmt.Sprintf("Description: %s\n\n", markdown.SanitizeItemForDocument(input.Description, settings)))
	buffer.WriteString(fmt.Sprintf("Type: %s", getInputType(input)))

	// Don't print defaults for required inputs when we're already explicit about it being required
	if input.HasDefault() || !settings.ShowRequired {
		buffer.WriteString(fmt.Sprintf("\nDefault: %s", getInputValue(input)))
	}
}

func printInputsRequired(buffer *bytes.Buffer, inputs []*tfconf.Input, settings *print.Settings) {
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

func printInputsOptional(buffer *bytes.Buffer, inputs []*tfconf.Input, settings *print.Settings) {
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

func printInputsAll(buffer *bytes.Buffer, inputs []*tfconf.Input, settings *print.Settings) {
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

func printProviders(buffer *bytes.Buffer, providers []*tfconf.Provider, settings *print.Settings) {
	buffer.WriteString(fmt.Sprintf("%s Providers\n\n", markdown.GenerateIndentation(0, settings)))

	if len(providers) == 0 {
		buffer.WriteString("No provider.\n\n")
		return
	}

	buffer.WriteString("The following providers are used by this module:\n")

	for _, provider := range providers {
		buffer.WriteString("\n")
		buffer.WriteString(fmt.Sprintf("- %s%s\n", markdown.SanitizeName(provider.GetName(), settings), getProviderVersion(provider)))
	}
	buffer.WriteString("\n")
}

func printInputs(buffer *bytes.Buffer, module *tfconf.Module, settings *print.Settings) {
	if settings.ShowRequired {
		printInputsRequired(buffer, module.RequiredInputs, settings)
		printInputsOptional(buffer, module.OptionalInputs, settings)
	} else {
		printInputsAll(buffer, module.Inputs, settings)
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
