package print

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

const (
	_ settings.Flag = iota
	// WithAggregateTypeDefaults prints defaults of aggregate type inputs
	WithAggregateTypeDefaults
	// WithRequired prints if inputs are required
	WithRequired
	// WithSortByName prints sorted inputs and outputs
	WithSortByName
	// WithSortInputsByRequired sorts inputs by name and prints required inputs first
	WithSortInputsByRequired
	// WithModules will add used modules to output
	WithModules
	// WithResources will add created resources to output
	WithResources
	// WithLinksToModules will add links to nested modules documentation into markdown output
	WithLinksToModules

	// ModuleDocumentationFileName holds filename to which should links to modules point
	ModuleDocumentationFileName = "documentation_file_name"
)

// GetPrintableValue returns a printable representation of a Terraform value.
func GetPrintableValue(value *doc.Value, settings settings.Settings, pretty bool) string {
	var result string

	if value == nil {
		return ""
	}

	if value.IsAggregateType() {
		if settings.Has(WithAggregateTypeDefaults) {
			if value.Value == nil {
				if value.Type == "list" {
					result = "[]"
				} else if value.Type == "map" {
					result = "{}"
				}
			} else {
				result = getFormattedJSONString(value.Value, pretty)
			}
		} else {
			result = "<" + value.Type + ">"
		}
	} else {
		result = getFormattedJSONString(value.Value, pretty)
	}

	return result
}

func getFormattedJSONString(value interface{}, pretty bool) string {
	if pretty {
		return getMultiLineJSONString(value)
	}

	return getSingleLineJSONString(value)
}

func getMultiLineJSONString(value interface{}) string {
	buffer, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(buffer)
}

func getSingleLineJSONString(value interface{}) string {
	buffer, err := json.MarshalIndent(value, "", "")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(string(buffer), "\n", " ", -1)
}

func (printer Printer) Print(document *doc.Doc, settings settings.Settings) (string, error) {
	var buffer bytes.Buffer

	separationNeeded := false

	if document.HasComment() {
		printer.PrintComment(&buffer, document.Comment, settings)
	}

	if document.HasInputs() {
		if settings.Has(WithSortByName) {
			if settings.Has(WithSortInputsByRequired) {
				doc.SortInputsByRequired(document.Inputs)
			} else {
				doc.SortInputsByName(document.Inputs)
			}
		}

		printer.PrintInputs(&buffer, document.Inputs, settings)
		separationNeeded = true
	}

	if document.HasModules() && settings.Has(WithModules) {
		if settings.Has(WithSortByName) {
			doc.SortModulesByName(document.Modules)
		}

		if separationNeeded {
			printer.PrintSeparator(&buffer, settings)
		}

		printer.PrintModules(&buffer, document.Modules, settings)
		separationNeeded = true
	}

	if document.HasResources() && settings.Has(WithResources) {
		if settings.Has(WithSortByName) {
			doc.SortResourcesByName(document.Resources)
		}

		if separationNeeded {
			printer.PrintSeparator(&buffer, settings)
		}

		printer.PrintResources(&buffer, document.Resources, settings)
		separationNeeded = true
	}

	if document.HasOutputs() {
		if settings.Has(WithSortByName) {
			doc.SortOutputsByName(document.Outputs)
		}

		if separationNeeded {
			printer.PrintSeparator(&buffer, settings)
		}

		printer.PrintOutputs(&buffer, document.Outputs, settings)
		separationNeeded = true
	}

	return printer.Postprocessing(&buffer)
}
