package table

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

type MarkdownTable struct{}

func (printer MarkdownTable) Postprocessing(buffer *bytes.Buffer) (string, error) {
	return markdown.Sanitize(buffer.String()), nil
}

func (printer MarkdownTable) PrintSeparator(buffer *bytes.Buffer, settings settings.Settings) {
	buffer.WriteString("\n")
}

func getInputDefaultValue(input *doc.Input, settings settings.Settings) string {
	var result = "n/a"

	if input.HasDefault() {
		result = fmt.Sprintf("`%s`", print.GetPrintableValue(input.Default, settings, false))
	}

	return result
}

func (printer MarkdownTable) PrintComment(buffer *bytes.Buffer, comment string, settings settings.Settings) {
	buffer.WriteString(fmt.Sprintf("%s\n", comment))
}

func (printer MarkdownTable) PrintInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings) {
	buffer.WriteString("## Inputs\n\n")
	buffer.WriteString("| Name | Description | Type | Default |")

	if settings.Has(print.WithRequired) {
		buffer.WriteString(" Required |\n")
	} else {
		buffer.WriteString("\n")
	}

	buffer.WriteString("|------|-------------|:----:|:-----:|")

	if settings.Has(print.WithRequired) {
		buffer.WriteString(":-----:|\n")
	} else {
		buffer.WriteString("\n")
	}

	for _, input := range inputs {
		buffer.WriteString(
			fmt.Sprintf("| %s | %s | %s | %s |",
				strings.Replace(input.Name, "_", "\\_", -1),
				markdown.ConvertMultiLineText(input.Description),
				input.Type,
				getInputDefaultValue(&input, settings)))

		if settings.Has(print.WithRequired) {
			buffer.WriteString(fmt.Sprintf(" %v |\n", printIsInputRequired(&input)))
		} else {
			buffer.WriteString("\n")
		}
	}
}

func printIsInputRequired(input *doc.Input) string {
	if input.IsRequired() {
		return "yes"
	}

	return "no"
}

func (printer MarkdownTable) PrintOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings) {
	buffer.WriteString("## Outputs\n\n")
	buffer.WriteString("| Name | Description |\n")
	buffer.WriteString("|------|-------------|\n")

	for _, output := range outputs {
		buffer.WriteString(
			fmt.Sprintf("| %s | %s |\n",
				strings.Replace(output.Name, "_", "\\_", -1),
				markdown.ConvertMultiLineText(output.Description)))
	}
}

func (printer MarkdownTable) PrintModules(buffer *bytes.Buffer, modules []doc.Module, settings settings.Settings) {
	buffer.WriteString("## Modules\n\n")
	buffer.WriteString("| Name | Source | Description |\n")
	buffer.WriteString("|------|--------|-------------|\n")

	for _, module := range modules {
		buffer.WriteString(fmt.Sprintf("| %s ", strings.Replace(module.Name, "_", "\\_", -1)))
		if settings.Has(print.WithLinksToModules) {
			buffer.WriteString(fmt.Sprintf("| [%s](%s/%s.md) ", module.Source, module.Source, settings.Get(print.ModuleDocumentationFileName)))
		} else {
			buffer.WriteString(fmt.Sprintf("| %s ", module.Source))
		}
		buffer.WriteString(fmt.Sprintf("| %s ", module.Description))
		buffer.WriteString("|\n")
	}
}

func (printer MarkdownTable) PrintResources(buffer *bytes.Buffer, resources []doc.Resource, settings settings.Settings) {
	buffer.WriteString("## Resources\n\n")
	buffer.WriteString("| Name | Type | Description |\n")
	buffer.WriteString("|------|------|-------------|\n")

	for _, resource := range resources {
		buffer.WriteString(fmt.Sprintf("| %s ", strings.Replace(resource.Name, "_", "\\_", -1)))
		buffer.WriteString(fmt.Sprintf("| [%s](https://www.terraform.io/docs/providers/%s/r/%s.html) ", resource.Type, resource.Type.Provider(), resource.Type.Name()))
		buffer.WriteString(fmt.Sprintf("| %s ", resource.Description))
		buffer.WriteString("|\n")
	}
}
