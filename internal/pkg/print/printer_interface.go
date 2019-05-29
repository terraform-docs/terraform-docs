package print

import (
	"bytes"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

type PrinterInterface interface {
	PrintInputs(buffer *bytes.Buffer, inputs []doc.Input, settings settings.Settings)
	PrintOutputs(buffer *bytes.Buffer, outputs []doc.Output, settings settings.Settings)
	PrintModules(buffer *bytes.Buffer, modules []doc.Module, settings settings.Settings)
	PrintResources(buffer *bytes.Buffer, resources []doc.Resource, settings settings.Settings)
	PrintComment(buffer *bytes.Buffer, comment string, settings settings.Settings)
	PrintSeparator(buffer *bytes.Buffer, settings settings.Settings)
	Postprocessing(buffer *bytes.Buffer) (string, error)
}
