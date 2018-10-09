package print

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

const (
	_ settings.Setting = iota
	// WithAggregateTypeDefaults prints defaults of aggregate type inputs
	WithAggregateTypeDefaults
	// WithRequired prints if inputs are required
	WithRequired
	// WithSortByName prints sorted inputs and outputs
	WithSortByName
	// WithSortInputsByRequired sorts inputs by name and prints required inputs first
	WithSortInputsByRequired
)

// GetPrintableValue returns a printable representation of a Terraform value.
func GetPrintableValue(value *doc.Value, settings settings.Settings) string {
	var result string

	if value == nil {
		return ""
	}

	switch value.Type {
	case "list":
		if settings.Has(WithAggregateTypeDefaults) {
			if value.Value != nil {
				// Convert the Go array into a JSON array
				json, err := json.MarshalIndent(value.Value, "", "")
				if err != nil {
					log.Fatal(err)
				}

				// Convert the JSON array into a string
				result = strings.Replace(string(json), "\n", " ", -1)
			} else {
				result = "[]"
			}
		} else {
			result = "<list>"
		}
	case "map":
		if settings.Has(WithAggregateTypeDefaults) {
			if value.Value != nil {
				// Convert the Go map into a JSON map
				json, err := json.MarshalIndent(value.Value, "", "")
				if err != nil {
					log.Fatal(err)
				}

				// Convert the JSON map into a string
				result = strings.Replace(string(json), "\n", " ", -1)
			} else {
				result = "{}"
			}
		} else {
			result = "<map>"
		}
	case "string":
		result = value.Value.(string)
	}

	return result
}
