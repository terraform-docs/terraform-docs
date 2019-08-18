package document_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_settings.Settings{}

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

	var settings = &_settings.Settings{
		AggregateTypeDefaults: true,
	}

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

	var settings = &_settings.Settings{
		ShowRequired: true,
	}

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

	var settings = &_settings.Settings{
		SortByName: true,
	}

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

	var settings = &_settings.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

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

func TestPrintWithEscapeName(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_settings.Settings{
		EscapeMarkdown: true,
	}

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithEscapeName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithIndentationBellowAllowed(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_settings.Settings{
		MarkdownIndent: 0,
	}

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithIndentationBellowAllowed")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithIndentationAboveAllowed(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_settings.Settings{
		MarkdownIndent: 10,
	}

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithIndentationAboveAllowed")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithIndentationOfFour(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_settings.Settings{
		MarkdownIndent: 4,
	}

	actual, err := document.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithIndentationOfFour")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
