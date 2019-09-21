package pretty_test

import (
	"testing"

	"github.com/segmentio/terraform-docs/internal/pkg/doc"
	"github.com/segmentio/terraform-docs/internal/pkg/print/pretty"
	_settings "github.com/segmentio/terraform-docs/internal/pkg/settings"
	"github.com/stretchr/testify/assert"
)

func TestPretty(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings = &_settings.Settings{}

	actual, err := pretty.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	sgr_color_1 := "\x1b[36m"
	sgr_color_2 := "\x1b[90m"
	sgr_reset := "\x1b[0m"

	expected :=
		"\nUsage:\n" +
			"\n" +
			"module \"foo\" {\n" +
			"  source = \"github.com/foo/bar\"\n" +
			"\n" +
			"  id   = \"1234567890\"\n" +
			"  name = \"baz\"\n" +
			"\n" +
			"  zones = [\"us-east-1\", \"us-west-1\"]\n" +
			"\n" +
			"  tags = {\n" +
			"    Name         = \"baz\"\n" +
			"    Created-By   = \"first.last@email.com\"\n" +
			"    Date-Created = \"20180101\"\n" +
			"  }\n" +
			"}\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "var.unquoted" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-3" + sgr_reset + " (\"\")\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's string number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-1" + sgr_reset + " (\"bar\")\n" +
			"  " + sgr_color_2 + "It's string number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-3" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's map number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-1" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "It's map number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-3" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's list number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-1" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "It's list number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input_with_underscores" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "A variable with underscores." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-pipe" + sgr_reset + " (\"v1\")\n" +
			"  " + sgr_color_2 + "It includes v1 | v2 | v3" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-code-block" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgr_reset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "output.unquoted" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's unquoted output." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-2" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-1" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number one." + sgr_reset + "\n" +
			"\n" +
			"\n"

	assert.Equal(t, expected, actual)
}

func TestPrettyWithWithAggregateTypeDefaults(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings = &_settings.Settings{
		AggregateTypeDefaults: true,
	}

	actual, err := pretty.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	sgr_color_1 := "\x1b[36m"
	sgr_color_2 := "\x1b[90m"
	sgr_reset := "\x1b[0m"

	expected :=
		"\nUsage:\n" +
			"\n" +
			"module \"foo\" {\n" +
			"  source = \"github.com/foo/bar\"\n" +
			"\n" +
			"  id   = \"1234567890\"\n" +
			"  name = \"baz\"\n" +
			"\n" +
			"  zones = [\"us-east-1\", \"us-west-1\"]\n" +
			"\n" +
			"  tags = {\n" +
			"    Name         = \"baz\"\n" +
			"    Created-By   = \"first.last@email.com\"\n" +
			"    Date-Created = \"20180101\"\n" +
			"  }\n" +
			"}\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "var.unquoted" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-3" + sgr_reset + " (\"\")\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's string number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-1" + sgr_reset + " (\"bar\")\n" +
			"  " + sgr_color_2 + "It's string number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-3" + sgr_reset + " ({})\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's map number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-1" + sgr_reset + " ({ \"a\": 1, \"b\": 2, \"c\": 3 })\n" +
			"  " + sgr_color_2 + "It's map number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-3" + sgr_reset + " ([])\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's list number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-1" + sgr_reset + " ([ \"a\", \"b\", \"c\" ])\n" +
			"  " + sgr_color_2 + "It's list number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input_with_underscores" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "A variable with underscores." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-pipe" + sgr_reset + " (\"v1\")\n" +
			"  " + sgr_color_2 + "It includes v1 | v2 | v3" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-code-block" + sgr_reset + " ([ \"name rack:location\" ])\n" +
			"  " + sgr_color_2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgr_reset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "output.unquoted" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's unquoted output." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-2" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-1" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number one." + sgr_reset + "\n" +
			"\n" +
			"\n"

	assert.Equal(t, expected, actual)
}

func TestPrettyWithSortByName(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings = &_settings.Settings{
		SortByName: true,
	}

	actual, err := pretty.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	sgr_color_1 := "\x1b[36m"
	sgr_color_2 := "\x1b[90m"
	sgr_reset := "\x1b[0m"

	expected :=
		"\nUsage:\n" +
			"\n" +
			"module \"foo\" {\n" +
			"  source = \"github.com/foo/bar\"\n" +
			"\n" +
			"  id   = \"1234567890\"\n" +
			"  name = \"baz\"\n" +
			"\n" +
			"  zones = [\"us-east-1\", \"us-west-1\"]\n" +
			"\n" +
			"  tags = {\n" +
			"    Name         = \"baz\"\n" +
			"    Created-By   = \"first.last@email.com\"\n" +
			"    Date-Created = \"20180101\"\n" +
			"  }\n" +
			"}\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-code-block" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-pipe" + sgr_reset + " (\"v1\")\n" +
			"  " + sgr_color_2 + "It includes v1 | v2 | v3" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input_with_underscores" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "A variable with underscores." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-1" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "It's list number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's list number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-3" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-1" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "It's map number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's map number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-3" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-1" + sgr_reset + " (\"bar\")\n" +
			"  " + sgr_color_2 + "It's string number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's string number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-3" + sgr_reset + " (\"\")\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.unquoted" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-1" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-2" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.unquoted" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's unquoted output." + sgr_reset + "\n" +
			"\n" +
			"\n"

	assert.Equal(t, expected, actual)
}

func TestPrettyWithSortInputsByRequired(t *testing.T) {
	doc := doc.TestDoc(t, "..")

	var settings = &_settings.Settings{
		SortByName:           true,
		SortInputsByRequired: true,
	}

	actual, err := pretty.Print(doc, settings)
	if err != nil {
		t.Fatal(err)
	}

	sgr_color_1 := "\x1b[36m"
	sgr_color_2 := "\x1b[90m"
	sgr_reset := "\x1b[0m"

	expected :=
		"\nUsage:\n" +
			"\n" +
			"module \"foo\" {\n" +
			"  source = \"github.com/foo/bar\"\n" +
			"\n" +
			"  id   = \"1234567890\"\n" +
			"  name = \"baz\"\n" +
			"\n" +
			"  zones = [\"us-east-1\", \"us-west-1\"]\n" +
			"\n" +
			"  tags = {\n" +
			"    Name         = \"baz\"\n" +
			"    Created-By   = \"first.last@email.com\"\n" +
			"    Date-Created = \"20180101\"\n" +
			"  }\n" +
			"}\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input_with_underscores" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "A variable with underscores." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's list number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's map number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-2" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "It's string number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.unquoted" + sgr_reset + " (required)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-code-block" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.input-with-pipe" + sgr_reset + " (\"v1\")\n" +
			"  " + sgr_color_2 + "It includes v1 | v2 | v3" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-1" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "It's list number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.list-3" + sgr_reset + " (<list>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-1" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "It's map number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.map-3" + sgr_reset + " (<map>)\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-1" + sgr_reset + " (\"bar\")\n" +
			"  " + sgr_color_2 + "It's string number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "var.string-3" + sgr_reset + " (\"\")\n" +
			"  " + sgr_color_2 + "" + sgr_reset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-1" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number one." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.output-2" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's output number two." + sgr_reset + "\n" +
			"\n" +
			"  " + sgr_color_1 + "output.unquoted" + sgr_reset + "\n" +
			"  " + sgr_color_2 + "It's unquoted output." + sgr_reset + "\n" +
			"\n" +
			"\n"

	assert.Equal(t, expected, actual)
}
