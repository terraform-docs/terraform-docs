package tfconf

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/terraform-docs/terraform-docs/internal/types"
)

func TestProviderNameWithoutAlias(t *testing.T) {
	assert := assert.New(t)
	provider := Provider{
		Name:     "provider",
		Alias:    types.String(""),
		Version:  types.String(">= 1.2.3"),
		Position: Position{Filename: "foo.tf", Line: 13},
	}
	assert.Equal("provider", provider.FullName())
}

func TestProviderNameWithAlias(t *testing.T) {
	assert := assert.New(t)
	provider := Provider{
		Name:     "provider",
		Alias:    types.String("alias"),
		Version:  types.String(">= 1.2.3"),
		Position: Position{Filename: "foo.tf", Line: 13},
	}
	assert.Equal("provider.alias", provider.FullName())
}
