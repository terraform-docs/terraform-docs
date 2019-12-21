package document_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/document"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortVariablesByRequired(t *testing.T) {
	var settings = &print.Settings{
		SortByName:              true,
		SortVariablesByRequired: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		EscapeMarkdown: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := document.Print(module, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("document-WithIndentationOfFour")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
