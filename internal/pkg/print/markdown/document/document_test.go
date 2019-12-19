package document_test

import (
	"testing"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName: true,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
	}

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

func TestPrintWithRequired(t *testing.T) {
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:   true,
		ShowRequired: true,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)

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
	var settings = &print.Settings{
		SortByName: true,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
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
	var settings = &print.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)

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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		EscapeMarkdown: true,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 0,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 10,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 4,
	}

	module, diag := tfconfig.LoadModule("../../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
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
