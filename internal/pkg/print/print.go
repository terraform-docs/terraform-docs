package print

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
)

// GetPrintableValue returns a printable representation of a Terraform value.
func GetPrintableValue(value *doc.Value, settings settings.Settings, pretty bool) string {
	var result string

	if value == nil {
		return ""
	}

	if value.IsAggregateType() {
		if settings.AggregateTypeDefaults {
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
