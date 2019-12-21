package pretty_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/print"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	"github.com/segmentio/terraform-docs/internal/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	// TODO remove SortByName when --no-sort for Terraform 0.12 is implemented
	var settings = &print.Settings{
		SortByName: true,
	}

	module, err := tfconf.CreateModule("../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := pretty.Print(module, settings)
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

	module, err := tfconf.CreateModule("../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := pretty.Print(module, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("pretty-WithSortByName")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrettyWithSortVariablesByRequired(t *testing.T) {
	var settings = &print.Settings{
		SortByName:              true,
		SortVariablesByRequired: true,
	}

	module, err := tfconf.CreateModule("../../../../examples")
	if err != nil {
		t.Fatal(err)
	}

	actual, err := pretty.Print(module, settings)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("pretty-WithSortVariablesByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
