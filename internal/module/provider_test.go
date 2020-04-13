package module

import (
	"sort"
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/segmentio/terraform-docs/pkg/tfconf"
	"github.com/stretchr/testify/assert"
)

func TestProvidersSortedByName(t *testing.T) {
	assert := assert.New(t)
	providers := sampleProviders()

	sort.Sort(providersSortedByName(providers))

	expected := []string{"a", "b", "c", "d", "d.a", "e", "e.a"}
	actual := make([]string, len(providers))

	for k, p := range providers {
		actual[k] = p.FullName()
	}

	assert.Equal(expected, actual)
}

func TestProvidersSortedByPosition(t *testing.T) {
	assert := assert.New(t)
	providers := sampleProviders()

	sort.Sort(providersSortedByPosition(providers))

	expected := []string{"e.a", "b", "d", "d.a", "a", "e", "c"}
	actual := make([]string, len(providers))

	for k, p := range providers {
		actual[k] = p.FullName()
	}

	assert.Equal(expected, actual)
}

func sampleProviders() []*tfconf.Provider {
	return []*tfconf.Provider{
		{
			Name:     "d",
			Alias:    types.String(""),
			Version:  types.String("1.3.2"),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 21},
		},
		{
			Name:     "d",
			Alias:    types.String("a"),
			Version:  types.String("> 1.x"),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 25},
		},
		{
			Name:     "b",
			Alias:    types.String(""),
			Version:  types.String("= 2.1.0"),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 13},
		},
		{
			Name:     "a",
			Alias:    types.String(""),
			Version:  types.String(""),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 39},
		},
		{
			Name:     "c",
			Alias:    types.String(""),
			Version:  types.String("~> 0.5.0"),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 53},
		},
		{
			Name:     "e",
			Alias:    types.String(""),
			Version:  types.String(""),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 47},
		},
		{
			Name:     "e",
			Alias:    types.String("a"),
			Version:  types.String("> 1.0"),
			Position: tfconf.Position{Filename: "foo/main.tf", Line: 5},
		},
	}
}
