package doc_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/stretchr/testify/assert"
)

var inputUnquoted = doc.Input{
	Name:        "unquoted",
	Description: "",
	Default:     nil,
	Type:        "string",
}

var inputString3 = doc.Input{
	Name:        "string-3",
	Description: "",
	Default: &doc.Value{
		Type:  "string",
		Value: "",
	},
	Type: "string",
}

var inputString2 = doc.Input{
	Name:        "string-2",
	Description: "It's string number two.",
	Default:     nil,
	Type:        "string",
}

var inputString1 = doc.Input{
	Name:        "string-1",
	Description: "It's string number one.",
	Default: &doc.Value{
		Type:  "string",
		Value: "bar",
	},
	Type: "string",
}

var inputMap3 = doc.Input{
	Name:        "map-3",
	Description: "",
	Default: &doc.Value{
		Type:  "map",
		Value: map[string]interface{}{},
	},
	Type: "map",
}

var inputMap2 = doc.Input{
	Name:        "map-2",
	Description: "It's map number two.",
	Default:     nil,
	Type:        "map",
}

var inputMap1 = doc.Input{
	Name:        "map-1",
	Description: "It's map number one.",
	Default: &doc.Value{
		Type: "map",
		Value: map[string]interface{}{
			"a": float64(1),
			"b": float64(2),
			"c": float64(3),
		},
	},
	Type: "map",
}

var inputList3 = doc.Input{
	Name:        "list-3",
	Description: "",
	Default: &doc.Value{
		Type:  "list",
		Value: []interface{}{},
	},
	Type: "list",
}

var inputList2 = doc.Input{
	Name:        "list-2",
	Description: "It's list number two.",
	Default:     nil,
	Type:        "list",
}

var inputList1 = doc.Input{
	Name:        "list-1",
	Description: "It's list number one.",
	Default: &doc.Value{
		Type: "list",
		Value: []interface{}{
			"a",
			"b",
			"c",
		},
	},
	Type: "list",
}

var inputWithUnderscores = doc.Input{
	Name:        "input_with_underscores",
	Description: "A variable with underscores.",
	Type:        "string",
	Default:     nil,
}

var inputWithPipe = doc.Input{
	Name:        "input-with-pipe",
	Description: "It includes v1 | v2 | v3",
	Default: &doc.Value{
		Type:  "string",
		Value: "v1",
	},
	Type: "string",
}

var inputWithCodeBlock = doc.Input{
	Name:        "input-with-code-block",
	Description: "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n",
	Default: &doc.Value{
		Type: "list",
		Value: []interface{}{
			"name rack:location",
		},
	},
	Type: "list",
}

var output1 = doc.Output{
	Name:        "output-1",
	Description: "It's output number one.",
}

var output2 = doc.Output{
	Name:        "output-2",
	Description: "It's output number two.",
}

var outputUnquoted = doc.Output{
	Name:        "unquoted",
	Description: "It's unquoted output.",
}

func TestComment(t *testing.T) {
	actual := doc.TestDoc(t, ".").Comment

	expected := `Usage:

module "foo" {
  source = "github.com/foo/bar"

  id   = "1234567890"
  name = "baz"

  zones = ["us-east-1", "us-west-1"]

  tags = {
    Name         = "baz"
    Created-By   = "first.last@email.com"
    Date-Created = "20180101"
  }
}
`

	assert.Equal(t, expected, actual)
}

func TestCommentFromMainTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "main.tf").Comment

	expected := `Usage:

module "foo" {
  source = "github.com/foo/bar"

  id   = "1234567890"
  name = "baz"

  zones = ["us-east-1", "us-west-1"]

  tags = {
    Name         = "baz"
    Created-By   = "first.last@email.com"
    Date-Created = "20180101"
  }
}
`

	assert.Equal(t, expected, actual)
}

func TestCommentFromOutputsTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "outputs.tf").Comment

	expected := ""

	assert.Equal(t, expected, actual)
}

func TestCommentFromVariablesTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "variables.tf").Comment

	expected := ""

	assert.Equal(t, expected, actual)
}

func TestInputs(t *testing.T) {
	actual := doc.TestDoc(t, ".").Inputs

	expected := []doc.Input{
		inputUnquoted,
		inputString3,
		inputString2,
		inputString1,
		inputMap3,
		inputMap2,
		inputMap1,
		inputList3,
		inputList2,
		inputList1,
		inputWithUnderscores,
		inputWithPipe,
		inputWithCodeBlock,
	}

	assert.Equal(t, expected, actual)
}

func TestInputsFromMainTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "main.tf").Inputs
	assert.Equal(t, 0, len(actual))
}

func TestInputsFromOutputsTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "outputs.tf").Inputs
	assert.Equal(t, 0, len(actual))
}

func TestInputsFromVariablesTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "variables.tf").Inputs

	expected := []doc.Input{
		inputUnquoted,
		inputString3,
		inputString2,
		inputString1,
		inputMap3,
		inputMap2,
		inputMap1,
		inputList3,
		inputList2,
		inputList1,
		inputWithUnderscores,
		inputWithPipe,
		inputWithCodeBlock,
	}

	assert.Equal(t, expected, actual)
}

func TestInputsSortedByName(t *testing.T) {
	actual := doc.TestDoc(t, ".").Inputs

	doc.SortInputsByName(actual)

	expected := []doc.Input{
		inputWithCodeBlock,
		inputWithPipe,
		inputWithUnderscores,
		inputList1,
		inputList2,
		inputList3,
		inputMap1,
		inputMap2,
		inputMap3,
		inputString1,
		inputString2,
		inputString3,
		inputUnquoted,
	}

	assert.Equal(t, expected, actual)
}

func TestInputsSortedByRequired(t *testing.T) {
	actual := doc.TestDoc(t, ".").Inputs

	doc.SortInputsByRequired(actual)

	expected := []doc.Input{
		inputWithUnderscores,
		inputList2,
		inputMap2,
		inputString2,
		inputUnquoted,
		inputWithCodeBlock,
		inputWithPipe,
		inputList1,
		inputList3,
		inputMap1,
		inputMap3,
		inputString1,
		inputString3,
	}

	assert.Equal(t, expected, actual)
}

func TestOutputs(t *testing.T) {
	actual := doc.TestDoc(t, ".").Outputs

	expected := []doc.Output{
		outputUnquoted,
		output2,
		output1,
	}

	assert.Equal(t, expected, actual)
}

func TestOutputsFromMainTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "main.tf").Outputs
	assert.Equal(t, 0, len(actual))
}

func TestOutputsFromOutputsTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "outputs.tf").Outputs

	expected := []doc.Output{
		outputUnquoted,
		output2,
		output1,
	}

	assert.Equal(t, expected, actual)
}

func TestOutputsFromVariablesTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "variables.tf").Outputs
	assert.Equal(t, 0, len(actual))
}
