package pretty

import (
	"bytes"
	"fmt"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

type Pretty struct{}

func (printer Pretty) Postprocessing(buffer *bytes.Buffer) (string, error) {
	return buffer.String(), nil
}

func (printer Pretty) PrintSeparator(buffer *bytes.Buffer, settings settings.Settings) {
	buffer.WriteString("\n")
}

func getInputDefaultValue(input *doc.Input, settings settings.Settings) string {
	var result = "required"

	if input.HasDefault() {
		result = print.GetPrintableValue(input.Default, settings, false)
	}

	return result
}

func (printer Pretty) PrintComment(buffer *bytes.Buffer, comment string, settings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("\n%s\n", comment))
}

func (printer Pretty) PrintInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	buffer.WriteString("\n")

	for _, input := range inputs {
		format := "  \033[36mvar.%s\033[0m (%s)\n  \033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(&input, settings),
				input.Description))
	}

	buffer.WriteString("\n")
}

func (printer Pretty) PrintOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("\n")

	for _, output := range outputs {
		format := "  \033[36moutput.%s\033[0m\n  \033[90m%s\033[0m\n\n"

		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				output.Description))
	}

	buffer.WriteString("\n")
}

func (printer Pretty) PrintModules(buffer *bytes.Buffer, modules []doc.Module, settings settings.Settings) {
	for _, module := range modules {
		format := "  \033[36mmodule.%s\033[0m%s\n  \033[90m%s\033[0m\n\n"
		description := ""
		if module.HasDescription() {
			description = fmt.Sprintf(" (%s)", module.Description)
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				module.Name,
				description,
				module.Source,
			))
	}
	buffer.WriteString("\n")
}

func (printer Pretty) PrintResources(buffer *bytes.Buffer, resources []doc.Resource, settings settings.Settings) {
	for _, resource := range resources {
		format := "  \033[36mresource.%s.%s\033[0m%s\n\n"
		description := ""
		if resource.HasDescription() {
			description = fmt.Sprintf(" (%s)", resource.Description)
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				resource.Type,
				resource.Name,
				description,
			))
	}
	buffer.WriteString("\n")
}
