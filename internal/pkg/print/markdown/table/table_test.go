package table_test

import (
	"testing"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	var settings print.Settings

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, settings)

	actual, err := table.Print(doc2, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithRequired(t *testing.T) {
	var settings print.Settings
	settings.Add(settings.WithRequired)

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, settings)

	actual, err := table.Print(doc2, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortInputsByRequired(t *testing.T) {
	var settings print.Settings
	settings.Add(settings.WithSortVariablesByRequired)

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, settings)

	actual, err := table.Print(doc2, settings)
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

	var settings = &_print.Settings{
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

func TestPrintWithIndentationBellowAllowed(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_print.Settings{
		MarkdownIndent: 0,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithIndentationBellowAllowed")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithIndentationAboveAllowed(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_print.Settings{
		MarkdownIndent: 10,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithIndentationAboveAllowed")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithIndentationOfFour(t *testing.T) {
	doc := doc.TestDoc(t, "../..")

	var settings = &_print.Settings{
		MarkdownIndent: 4,
	}

	actual, err := table.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithIndentationOfFour")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
