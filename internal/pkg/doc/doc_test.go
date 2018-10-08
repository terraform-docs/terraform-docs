package doc_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/stretchr/testify/assert"
)

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
		doc.Input{
			Name:        "unquoted",
			Description: "",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "string-2",
			Description: "It's string number two.",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "string-1",
			Description: "It's string number one.",
			Default: &doc.Value{
				Type:  "string",
				Value: "bar",
			},
			Type: "string",
		},
		doc.Input{
			Name:        "map-3",
			Description: "",
			Default: &doc.Value{
				Type:  "map",
				Value: map[string]interface{}{},
			},
			Type: "map",
		},
		doc.Input{
			Name:        "map-2",
			Description: "It's map number two.",
			Default:     nil,
			Type:        "map",
		},
		doc.Input{
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
		},
		doc.Input{
			Name:        "list-3",
			Description: "",
			Default: &doc.Value{
				Type:  "list",
				Value: []interface{}{},
			},
			Type: "list",
		},
		doc.Input{
			Name:        "list-2",
			Description: "It's list number two.",
			Default:     nil,
			Type:        "list",
		},
		doc.Input{
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
		},
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
		doc.Input{
			Name:        "unquoted",
			Description: "",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "string-2",
			Description: "It's string number two.",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "string-1",
			Description: "It's string number one.",
			Default: &doc.Value{
				Type:  "string",
				Value: "bar",
			},
			Type: "string",
		},
		doc.Input{
			Name:        "map-3",
			Description: "",
			Default: &doc.Value{
				Type:  "map",
				Value: map[string]interface{}{},
			},
			Type: "map",
		},
		doc.Input{
			Name:        "map-2",
			Description: "It's map number two.",
			Default:     nil,
			Type:        "map",
		},
		doc.Input{
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
		},
		doc.Input{
			Name:        "list-3",
			Description: "",
			Default: &doc.Value{
				Type:  "list",
				Value: []interface{}{},
			},
			Type: "list",
		},
		doc.Input{
			Name:        "list-2",
			Description: "It's list number two.",
			Default:     nil,
			Type:        "list",
		},
		doc.Input{
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
		},
	}

	assert.Equal(t, expected, actual)
}

func TestInputsSortedByName(t *testing.T) {
	actual := doc.TestDoc(t, ".").Inputs

	doc.SortInputsByName(actual)

	expected := []doc.Input{
		doc.Input{
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
		},
		doc.Input{
			Name:        "list-2",
			Description: "It's list number two.",
			Default:     nil,
			Type:        "list",
		},
		doc.Input{
			Name:        "list-3",
			Description: "",
			Default: &doc.Value{
				Type:  "list",
				Value: []interface{}{},
			},
			Type: "list",
		},
		doc.Input{
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
		},
		doc.Input{
			Name:        "map-2",
			Description: "It's map number two.",
			Default:     nil,
			Type:        "map",
		},
		doc.Input{
			Name:        "map-3",
			Description: "",
			Default: &doc.Value{
				Type:  "map",
				Value: map[string]interface{}{},
			},
			Type: "map",
		},
		doc.Input{
			Name:        "string-1",
			Description: "It's string number one.",
			Default: &doc.Value{
				Type:  "string",
				Value: "bar",
			},
			Type: "string",
		},
		doc.Input{
			Name:        "string-2",
			Description: "It's string number two.",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "unquoted",
			Description: "",
			Default:     nil,
			Type:        "string",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestInputsSortedByRequired(t *testing.T) {
	actual := doc.TestDoc(t, ".").Inputs

	doc.SortInputsByRequired(actual)

	expected := []doc.Input{
		doc.Input{
			Name:        "list-2",
			Description: "It's list number two.",
			Default:     nil,
			Type:        "list",
		},
		doc.Input{
			Name:        "map-2",
			Description: "It's map number two.",
			Default:     nil,
			Type:        "map",
		},
		doc.Input{
			Name:        "string-2",
			Description: "It's string number two.",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
			Name:        "unquoted",
			Description: "",
			Default:     nil,
			Type:        "string",
		},
		doc.Input{
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
		},
		doc.Input{
			Name:        "list-3",
			Description: "",
			Default: &doc.Value{
				Type:  "list",
				Value: []interface{}{},
			},
			Type: "list",
		},
		doc.Input{
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
		},
		doc.Input{
			Name:        "map-3",
			Description: "",
			Default: &doc.Value{
				Type:  "map",
				Value: map[string]interface{}{},
			},
			Type: "map",
		},
		doc.Input{
			Name:        "string-1",
			Description: "It's string number one.",
			Default: &doc.Value{
				Type:  "string",
				Value: "bar",
			},
			Type: "string",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestOutputs(t *testing.T) {
	actual := doc.TestDoc(t, ".").Outputs

	expected := []doc.Output{
		doc.Output{
			Name:        "unquoted",
			Description: "It's unquoted output.",
		},
		doc.Output{
			Name:        "output-2",
			Description: "It's output number two.",
		},
		doc.Output{
			Name:        "output-1",
			Description: "It's output number one.",
		},
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
		doc.Output{
			Name:        "unquoted",
			Description: "It's unquoted output.",
		},
		doc.Output{
			Name:        "output-2",
			Description: "It's output number two.",
		},
		doc.Output{
			Name:        "output-1",
			Description: "It's output number one.",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestOutputsFromVariablesTf(t *testing.T) {
	actual := doc.TestDocFromFile(t, ".", "variables.tf").Outputs
	assert.Equal(t, 0, len(actual))
}
