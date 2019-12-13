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

	sgrColor1 := "\x1b[36m"
	sgrColor2 := "\x1b[90m"
	sgrReset := "\x1b[0m"

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
			"  " + sgrColor1 + "var.unquoted" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-3" + sgrReset + " (\"\")\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's string number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-1" + sgrReset + " (\"bar\")\n" +
			"  " + sgrColor2 + "It's string number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-3" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's map number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-1" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "It's map number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-3" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's list number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-1" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "It's list number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input_with_underscores" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "A variable with underscores." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-pipe" + sgrReset + " (\"v1\")\n" +
			"  " + sgrColor2 + "It includes v1 | v2 | v3" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-code-block" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgrReset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgrColor1 + "output.unquoted" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's unquoted output." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-2" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-1" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number one." + sgrReset + "\n" +
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

	sgrColor1 := "\x1b[36m"
	sgrColor2 := "\x1b[90m"
	sgrReset := "\x1b[0m"

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
			"  " + sgrColor1 + "var.unquoted" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-3" + sgrReset + " (\"\")\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's string number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-1" + sgrReset + " (\"bar\")\n" +
			"  " + sgrColor2 + "It's string number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-3" + sgrReset + " ({})\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's map number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-1" + sgrReset + " ({ \"a\": 1, \"b\": 2, \"c\": 3 })\n" +
			"  " + sgrColor2 + "It's map number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-3" + sgrReset + " ([])\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's list number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-1" + sgrReset + " ([ \"a\", \"b\", \"c\" ])\n" +
			"  " + sgrColor2 + "It's list number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input_with_underscores" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "A variable with underscores." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-pipe" + sgrReset + " (\"v1\")\n" +
			"  " + sgrColor2 + "It includes v1 | v2 | v3" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-code-block" + sgrReset + " ([ \"name rack:location\" ])\n" +
			"  " + sgrColor2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgrReset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgrColor1 + "output.unquoted" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's unquoted output." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-2" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-1" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number one." + sgrReset + "\n" +
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

	sgrColor1 := "\x1b[36m"
	sgrColor2 := "\x1b[90m"
	sgrReset := "\x1b[0m"

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
			"  " + sgrColor1 + "var.input-with-code-block" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-pipe" + sgrReset + " (\"v1\")\n" +
			"  " + sgrColor2 + "It includes v1 | v2 | v3" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input_with_underscores" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "A variable with underscores." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-1" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "It's list number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's list number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-3" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-1" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "It's map number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's map number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-3" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-1" + sgrReset + " (\"bar\")\n" +
			"  " + sgrColor2 + "It's string number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's string number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-3" + sgrReset + " (\"\")\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.unquoted" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-1" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-2" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.unquoted" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's unquoted output." + sgrReset + "\n" +
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

	sgrColor1 := "\x1b[36m"
	sgrColor2 := "\x1b[90m"
	sgrReset := "\x1b[0m"

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
			"  " + sgrColor1 + "var.input_with_underscores" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "A variable with underscores." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's list number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's map number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-2" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "It's string number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.unquoted" + sgrReset + " (required)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-code-block" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.input-with-pipe" + sgrReset + " (\"v1\")\n" +
			"  " + sgrColor2 + "It includes v1 | v2 | v3" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-1" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "It's list number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.list-3" + sgrReset + " (<list>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-1" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "It's map number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.map-3" + sgrReset + " (<map>)\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-1" + sgrReset + " (\"bar\")\n" +
			"  " + sgrColor2 + "It's string number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "var.string-3" + sgrReset + " (\"\")\n" +
			"  " + sgrColor2 + "" + sgrReset + "\n" +
			"\n" +
			"\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-1" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number one." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.output-2" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's output number two." + sgrReset + "\n" +
			"\n" +
			"  " + sgrColor1 + "output.unquoted" + sgrReset + "\n" +
			"  " + sgrColor2 + "It's unquoted output." + sgrReset + "\n" +
			"\n" +
			"\n"

	assert.Equal(t, expected, actual)
}
