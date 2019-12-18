package pretty_test

import (
	"testing"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
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

	actual, err := pretty.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("pretty")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrettyWithSortByName(t *testing.T) {
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

	actual, err := pretty.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("pretty-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrettyWithSortInputsByRequired(t *testing.T) {
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

	actual, err := pretty.Print(document, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("pretty-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
