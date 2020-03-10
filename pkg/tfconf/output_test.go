package tfconf

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/segmentio/terraform-docs/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestOutputValue(t *testing.T) {
	outputs := sampleOutputs()
	tests := []struct {
		name          string
		output        Output
		expectValue   string
		expectDefault bool
	}{
		{
			name:          "output Value and HasDefault",
			output:        outputs[0],
			expectValue:   "",
			expectDefault: false,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[1],
			expectValue:   "",
			expectDefault: false,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[2],
			expectValue:   "false",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[3],
			expectValue:   "\"\"",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[4],
			expectValue:   "\"foo\"",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[5],
			expectValue:   "",
			expectDefault: false,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[6],
			expectValue:   "\"\\u003csensitive\\u003e\"",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[7],
			expectValue:   "[\n  \"a\",\n  \"b\",\n  \"c\"\n]",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[8],
			expectValue:   "[]",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[9],
			expectValue:   "{\n  \"a\": 1,\n  \"b\": 2,\n  \"c\": 3\n}",
			expectDefault: true,
		},
		{
			name:          "output Value and HasDefault",
			output:        outputs[10],
			expectValue:   "{}",
			expectDefault: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(tt.expectValue, tt.output.GetValue())
			assert.Equal(tt.expectDefault, tt.output.HasDefault())
		})
	}
}

func TestOutputMarshalJSON(t *testing.T) {
	outputs := sampleOutputs()
	tests := []struct {
		name     string
		output   Output
		expected string
	}{
		{
			name:     "output marshal JSON",
			output:   outputs[0],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":null,\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[1],
			expected: "{\"name\":\"output\",\"description\":\"description\"}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[2],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":false,\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[3],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":\"\",\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[4],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":\"foo\",\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[5],
			expected: "{\"name\":\"output\",\"description\":\"description\"}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[6],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":\"<sensitive>\",\"sensitive\":true}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[7],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":[\"a\",\"b\",\"c\"],\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[8],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":[],\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[9],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":{\"a\":1,\"b\":2,\"c\":3},\"sensitive\":false}\n",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[10],
			expected: "{\"name\":\"output\",\"description\":\"description\",\"value\":{},\"sensitive\":false}\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := tt.output.MarshalJSON()

			assert.Nil(err)
			assert.Equal(tt.expected, string(actual))
		})
	}
}

func TestOutputMarshalXML(t *testing.T) {
	outputs := sampleOutputs()
	tests := []struct {
		name     string
		output   Output
		expected string
	}{
		{
			name:     "output marshal XML",
			output:   outputs[0],
			expected: "<output><name>output</name><description>description</description><value xsi:nil=\"true\"></value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[1],
			expected: "<output><name>output</name><description>description</description></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[2],
			expected: "<output><name>output</name><description>description</description><value>false</value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[3],
			expected: "<output><name>output</name><description>description</description><value></value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[4],
			expected: "<output><name>output</name><description>description</description><value>foo</value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[5],
			expected: "<output><name>output</name><description>description</description></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[6],
			expected: "<output><name>output</name><description>description</description><value>&lt;sensitive&gt;</value><sensitive>true</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[7],
			expected: "<output><name>output</name><description>description</description><value><item>a</item><item>b</item><item>c</item></value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[8],
			expected: "<output><name>output</name><description>description</description><value></value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[9],
			expected: "<output><name>output</name><description>description</description><value><a>1</a><b>2</b><c>3</c></value><sensitive>false</sensitive></output>",
		},
		{
			name:     "output marshal XML",
			output:   outputs[10],
			expected: "<output><name>output</name><description>description</description><value></value><sensitive>false</sensitive></output>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			var b bytes.Buffer
			encoder := xml.NewEncoder(&b)
			start := xml.StartElement{Name: xml.Name{Local: "output"}}

			err := tt.output.MarshalXML(encoder, start)
			assert.Nil(err)

			err = encoder.Flush()
			assert.Nil(err)

			assert.Nil(err)
			assert.Equal(tt.expected, b.String())
		})
	}
}

func TestOutputMarshalYAML(t *testing.T) {
	outputs := sampleOutputs()
	tests := []struct {
		name     string
		output   Output
		expected string
	}{
		{
			name:     "output marshal JSON",
			output:   outputs[0],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[1],
			expected: "*tfconf.Output",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[2],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[3],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[4],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[5],
			expected: "*tfconf.Output",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[6],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[7],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[8],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[9],
			expected: "tfconf.withvalue",
		},
		{
			name:     "output marshal JSON",
			output:   outputs[10],
			expected: "tfconf.withvalue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			actual, err := tt.output.MarshalYAML()

			assert.Nil(err)
			assert.Equal(tt.expected, reflect.TypeOf(actual).String())
		})
	}
}

func sampleOutputs() []Output {
	name := "output"
	description := types.String("description")
	position := Position{Filename: "foo.tf", Line: 13}
	return []Output{
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(nil),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Position:    position,
			ShowValue:   false,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(false),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(""),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf("foo"),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf("this should be hidden"),
			Sensitive:   false,
			Position:    position,
			ShowValue:   false,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf("<sensitive>"),
			Sensitive:   true,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(types.List{"a", "b", "c"}.Underlying()),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(types.List{}.Underlying()),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(types.Map{"a": 1, "b": 2, "c": 3}.Underlying()),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
		{
			Name:        name,
			Description: description,
			Value:       types.ValueOf(types.Map{}.Underlying()),
			Sensitive:   false,
			Position:    position,
			ShowValue:   true,
		},
	}
}
