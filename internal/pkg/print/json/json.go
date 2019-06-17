package json

import (
	"encoding/json"
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
)

const (
	indent string = "  "
	prefix string = ""
)

// Print prints a document as json.
func Print(document *doc.Doc) (string, error) {
	buffer, err := json.MarshalIndent(document, prefix, indent)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
