package table_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	"github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	doc := doc.TestDoc(t, "../..")
	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithAggregateTypeDefaults)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithRequired)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithSortByName)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithSortByName)
	settings.Add(print.WithSortInputsByRequired)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithModules(t *testing.T) {
	doc := doc.TestDoc(t, "../..")
	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithModules)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithModules")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithResources(t *testing.T) {
	doc := doc.TestDoc(t, "../..")
	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithResources)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithResources")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithLinksToModules(t *testing.T) {
	doc := doc.TestDoc(t, "../..")
	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithModules)
	settings.Add(print.WithLinksToModules)

	actual, err := print.Printer{PrinterInterface: table.MarkdownTable{}}.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithLinksToModules")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
