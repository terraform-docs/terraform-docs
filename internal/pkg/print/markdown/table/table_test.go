package table_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/markdown/table"
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

	actual, err := table.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:   true,
		ShowRequired: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	var settings = &print.Settings{
		SortByName: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	var settings = &print.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		EscapeMarkdown: true,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 0,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 10,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
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
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName:     true,
		MarkdownIndent: 4,
	}

	module, err := tfconf.CreateModule("../../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := table.Print(module, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("table-WithIndentationOfFour")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
