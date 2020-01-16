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

	if settings.ShowHeader {
		printHeader(&buffer, module.Header, settings)
	}
	if settings.ShowProviders {
		printProviders(&buffer, module.Providers, settings)
	}
	if settings.ShowInputs {
		printInputs(&buffer, module.Inputs, settings)
	}
	if settings.ShowOutputs {
		printOutputs(&buffer, module.Outputs, settings)
	}

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
		result = string(input.Default)
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

func printHeader(buffer *bytes.Buffer, header string, settings *print.Settings) {
	if len(header) == 0 {
		return
	}
	buffer.WriteString("\n\n")

	for _, line := range strings.Split(header, "\n") {
		var format string
		if settings.ShowColor {
			format = "\033[90m%s\033[0m\n"
		} else {
			format = "%s\n"
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				line,
			),
		)
	}
	buffer.WriteString("\n")
}

func printProviders(buffer *bytes.Buffer, providers []*tfconf.Provider, settings *print.Settings) {
	buffer.WriteString("\n\n")

	for _, provider := range providers {
		var format string
		if settings.ShowColor {
			format = "\033[36mprovider.%s\033[0m%s\n\n"
		} else {
			format = "provider.%s%s\n\n"
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				provider.GetName(),
				getProviderVersion(provider),
			),
		)
	}
}

func printInputs(buffer *bytes.Buffer, inputs []*tfconf.Input, settings *print.Settings) {
	buffer.WriteString("\n\n")

	for _, input := range inputs {
		var format string
		if settings.ShowColor {
			format = "\033[36minput.%s\033[0m (%s)\n\033[90m%s\033[0m\n\n"
		} else {
			format = "input.%s (%s)\n%s\n\n"
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				input.Name,
				getInputDefaultValue(input, settings),
				getDescription(string(input.Description)),
			),
		)
	}
}

func printOutputs(buffer *bytes.Buffer, outputs []*tfconf.Output, settings *print.Settings) {
	buffer.WriteString("\n\n")

	for _, output := range outputs {
		var format string
		if settings.ShowColor {
			format = "\033[36moutput.%s\033[0m\n\033[90m%s\033[0m\n\n"
		} else {
			format = "output.%s\n%s\n\n"
		}
		buffer.WriteString(
			fmt.Sprintf(
				format,
				output.Name,
				getDescription(string(output.Description)),
			),
		)
	}
}
