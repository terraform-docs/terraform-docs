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

	var printSettings settings.Settings

	module, diag := tfconfig.LoadModule("../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, printSettings)

	actual, err := json.Print(doc2)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}

func TestPrintWithSortInputsByRequired(t *testing.T) {
	var printSettings settings.Settings
	printSettings.Add(settings.WithSortVariablesByRequired)

	module, diag := tfconfig.LoadModule("../../../../examples")
	if diag != nil && diag.HasErrors() {
		t.Fatal(diag)
	}

	doc2, err := doc.Create(module, printSettings)

	actual, err := json.Print(doc2)
	if err != nil {
		t.Fatal(err)
	}

	expected, err := print.ReadGoldenFile("json-WithSortInputsByRequired")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, expected, actual)
}
