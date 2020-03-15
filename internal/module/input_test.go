package module

import (
	"sort"
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestInputsSortedByName(t *testing.T) {
	assert := assert.New(t)
	inputs := sampleInputs()

	sort.Sort(inputsSortedByName(inputs))

	expected := []string{"a", "b", "c", "d", "e", "f"}
	actual := make([]string, len(inputs))

	for k, i := range inputs {
		actual[k] = i.Name
	}

	assert.Equal(expected, actual)
}

func TestInputsSortedByRequired(t *testing.T) {
	assert := assert.New(t)
	inputs := sampleInputs()

	sort.Sort(inputsSortedByRequired(inputs))

	expected := []string{"b", "d", "a", "c", "e", "f"}
	actual := make([]string, len(inputs))

	for k, i := range inputs {
		actual[k] = i.Name
	}

	assert.Equal(expected, actual)
}

func TestInputsSortedByPosition(t *testing.T) {
	assert := assert.New(t)
	inputs := sampleInputs()

	sort.Sort(inputsSortedByPosition(inputs))

	expected := []string{"a", "d", "e", "b", "c", "f"}
	actual := make([]string, len(inputs))

	for k, i := range inputs {
		actual[k] = i.Name
	}

	assert.Equal(expected, actual)
}

func sampleInputs() []*tfconf.Input {
	return []*tfconf.Input{
		&tfconf.Input{
			Name:        "e",
			Type:        types.String(""),
			Description: types.String("description of e"),
			Default:     types.ValueOf(true),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 35},
		},
		&tfconf.Input{
			Name:        "a",
			Type:        types.String("string"),
			Description: types.String(""),
			Default:     types.ValueOf("a"),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 10},
		},
		&tfconf.Input{
			Name:        "d",
			Type:        types.String("string"),
			Description: types.String("description for d"),
			Default:     types.ValueOf(nil),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 23},
		},
		&tfconf.Input{
			Name:        "b",
			Type:        types.String("number"),
			Description: types.String("description of b"),
			Default:     types.ValueOf(nil),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 42},
		},
		&tfconf.Input{
			Name:        "c",
			Type:        types.String("list"),
			Description: types.String("description of c"),
			Default:     types.ValueOf("c"),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 51},
		},
		&tfconf.Input{
			Name:        "f",
			Type:        types.String("string"),
			Description: types.String("description of f"),
			Default:     new(types.Null),
			Position:    tfconf.Position{Filename: "foo/variables.tf", Line: 59},
		},
	}
}
