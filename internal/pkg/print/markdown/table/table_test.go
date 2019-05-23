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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithAggregateTypeDefaults)
	settings.Add(print.WithLinksToModules)

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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithRequired)
	settings.Add(print.WithLinksToModules)

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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithSortByName)
	settings.Add(print.WithLinksToModules)

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

	settings := settings.Settings{Values: map[settings.Setting]string{print.ModuleDocumentationFileName: "readme"}}
	settings.Add(print.WithSortByName)
	settings.Add(print.WithSortInputsByRequired)
	settings.Add(print.WithLinksToModules)

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
