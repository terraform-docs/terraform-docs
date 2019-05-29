package print

import (
	"bytes"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

type Printer struct {
	PrinterInterface
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
	}

	return printer.Postprocessing(&buffer)
}
