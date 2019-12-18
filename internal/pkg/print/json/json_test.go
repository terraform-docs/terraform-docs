package json_test

import (
	"testing"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/json"
	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName: true,
	}

	module, diag := tfconfig.LoadModule("../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	document, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := json.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortByName(t *testing.T) {
	var settings = &print.Settings{
		SortByName: true,
	}

	module, diag := tfconfig.LoadModule("../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	document, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := json.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json-WithSortByName")
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

	module, diag := tfconfig.LoadModule("../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	document, err := doc.Create(module)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := json.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
