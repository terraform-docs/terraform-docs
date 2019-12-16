package document_test

import (
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	var printSettings settings.Settings

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, printSettings)

	actual, err := document.Print(doc2, printSettings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithRequired(t *testing.T) {
	var printSettings settings.Settings
	printSettings.Add(settings.WithRequired)

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, printSettings)

	actual, err := document.Print(doc2, printSettings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortVariablesByRequired(t *testing.T) {
	var printSettings settings.Settings
	printSettings.Add(settings.WithSortVariablesByRequired)

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, printSettings)

	actual, err := document.Print(doc2, printSettings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithSortVariablesByRequired")
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
