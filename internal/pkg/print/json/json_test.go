package json_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings settings.Settings

	actual, err := json.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortByName(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings settings.Settings
	settings.Add(print.WithSortByName)

	actual, err := json.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortInputsByRequired(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings settings.Settings
	settings.Add(print.WithSortByName)
	settings.Add(print.WithSortInputsByRequired)

	actual, err := json.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
