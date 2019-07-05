package table_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestWithAggregateTypeDefaults(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{
		AggregateTypeDefaults: true,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithAggregateTypeDefaults")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithRequired(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{
		ShowRequired: true,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortByName(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{
		SortByName: true,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortInputsByRequired(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithEscapeName(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = _settings.Settings{
		EscapeMarkdown: true,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithEscapeName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
