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
func GetPrintableValue(value *doc.Value, settings settings.Settings, pretty bool) string {
	var result string

	if value == nil {
		return ""
	}

	switch value.Type {
	case "list":
		if settings.Has(WithAggregateTypeDefaults) {
			if value.Value != nil {
				result = getFormattedJSONString(value.Value, pretty)
			} else {
				result = "[]"
			}
		} else {
			result = "<list>"
		}
	case "map":
		if settings.Has(WithAggregateTypeDefaults) {
			if value.Value != nil {
				result = getFormattedJSONString(value.Value, pretty)
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
