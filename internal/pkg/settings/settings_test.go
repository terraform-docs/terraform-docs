package settings_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

const (
	_ settings.Setting = iota
	A
	B
	C
)

func Test(t *testing.T) {
	var s settings.Settings

	s.Add(A)
	s.Has(A)
	assert.False(t, s.Has(B))
	assert.False(t, s.Has(C))

	s.Add(B)
	s.Has(A)
	s.Has(B)
	assert.False(t, s.Has(C))

	s.Add(C)
	s.Has(A)
	s.Has(B)
	s.Has(C)
}
