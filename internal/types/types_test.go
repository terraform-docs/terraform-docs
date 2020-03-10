package types

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type expected struct {
	typeName   string
	valueKind  string
	hasDefault bool
}

type testprimitive struct {
	name     string
	values   List
	types    string
	expected expected
}

type testlist struct {
	name     string
	values   []List
	types    string
	expected expected
}

type testmap struct {
	name     string
	values   []Map
	types    string
	expected expected
}

func testPrimitive(t *testing.T, tests []testprimitive) {
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv)
				actualType := TypeOf(tt.types, tv)

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}

func testList(t *testing.T, tests []testlist) {
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv.Underlying())
				actualType := TypeOf(tt.types, tv.Underlying())

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}

func testMap(t *testing.T, tests []testmap) {
	for _, tt := range tests {
		for _, tv := range tt.values {
			t.Run(tt.name, func(t *testing.T) {
				assert := assert.New(t)

				actualValue := ValueOf(tv.Underlying())
				actualType := TypeOf(tt.types, tv.Underlying())

				assert.Equal(tt.expected.typeName, string(actualType))
				assert.Equal(tt.expected.valueKind, reflect.TypeOf(actualValue).String())
				assert.Equal(tt.expected.hasDefault, actualValue.HasDefault())
			})
		}
	}
}
