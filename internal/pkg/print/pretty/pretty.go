package pretty

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
)

// Print prints a pretty document.
func Print(module *tfconf.Module, settings *print.Settings) (string, error) {
	var buffer bytes.Buffer

	module.Sort(settings)

	printProviders(&buffer, module.Providers, settings)
	printInputs(&buffer, module.Inputs, settings)
	printOutputs(&buffer, module.Outputs, settings)

	return buffer.String(), nil
}

func getProviderVersion(provider *tfconf.Provider) string {
	var result = ""
	if provider.Version != "" {
		result = fmt.Sprintf(" (%s)", provider.Version)
	}
	return result
}

func getInputDefaultValue(input *tfconf.Input, settings *print.Settings) string {
	var result = "required"

	if input.HasDefault() {
		result = input.Default
	}

	return result
}

func getDescription(description string) string {
	var result = "n/a"

	if description != "" {
		result = strings.TrimSuffix(description, "\n")
	}

	return result
}

func printProviders(buffer *bytes.Buffer, providers []*tfconf.Provider, settings *print.Settings) {
	buffer.WriteString("\n\n")

	for _, provider := range providers {
		format := "\033[36mprovider.%s\033[0m%s\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				provider.GetName(),
				getProviderVersion(provider),
			),
		)
	}
	buffer.WriteString("\n")
}

func printInputs(buffer *bytes.Buffer, inputs []*tfconf.Input, settings *print.Settings) {
	buffer.WriteString("\n")

	for _, input := range inputs {
		format := "\033[36minput.%s\033[0m (%s)\n\033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(input, settings),
				getDescription(input.Description),
			),
		)
	}

	buffer.WriteString("\n")
}

func printOutputs(buffer *bytes.Buffer, outputs []*tfconf.Output, settings *print.Settings) {
	buffer.WriteString("\n")

	for _, output := range outputs {
		format := "\033[36moutput.%s\033[0m\n\033[90m%s\033[0m\n\n"
		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				getDescription(output.Description),
			),
		)
	}
}
