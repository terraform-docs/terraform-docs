package document_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	doc := doc.TestDoc(t, "../..")
	var settings settings.Settings

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestWithAggregateTypeDefaults(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings settings.Settings
	settings.Add(print.WithAggregateTypeDefaults)

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithAggregateTypeDefaults")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithRequired(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings settings.Settings
	settings.Add(print.WithRequired)

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortByName(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings settings.Settings
	settings.Add(print.WithSortByName)

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortInputsByRequired(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings settings.Settings
	settings.Add(print.WithSortByName)
	settings.Add(print.WithSortInputsByRequired)

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
