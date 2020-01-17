# JSON

Generate a JSON of inputs and outputs.

## Usage

```text
Usage:
  terraform-docs json [PATH] [flags]

Flags:
  -h, --help        help for json
      --no-escape   do not escape special characters

Global Flags:
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

## Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs json ./examples/
```

generates the following output:

```json
{
  "header": "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n```hcl\nmodule \"foo_bar\" {\n  source = \"github.com/foo/bar\"\n\n  id   = \"1234567890\"\n  name = \"baz\"\n\n  zones = [\"us-east-1\", \"us-west-1\"]\n\n  tags = {\n    Name         = \"baz\"\n    Created-By   = \"first.last@email.com\"\n    Date-Created = \"20180101\"\n  }\n}\n```",
  "inputs": [
    {
      "name": "input-with-code-block",
      "type": "list",
      "description": "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n",
      "default": "[\n  \"name rack:location\"\n]"
    },
    {
      "name": "input-with-pipe",
      "type": "string",
      "description": "It includes v1 | v2 | v3",
      "default": "\"v1\""
    },
    {
      "name": "input_with_underscores",
      "type": "any",
      "description": "A variable with underscores.",
      "default": null
    },
    {
      "name": "list-1",
      "type": "list",
      "description": "It's list number one.",
      "default": "[\n  \"a\",\n  \"b\",\n  \"c\"\n]"
    },
    {
      "name": "list-2",
      "type": "list",
      "description": "It's list number two.",
      "default": null
    },
    {
      "name": "list-3",
      "type": "list",
      "description": null,
      "default": "[]"
    },
    {
      "name": "long_type",
      "type": "object({\n    name = string,\n    foo  = object({ foo = string, bar = string }),\n    bar  = object({ foo = string, bar = string }),\n    fizz = list(string),\n    buzz = list(string)\n  })",
      "description": "This description is itself markdown.\n\nIt spans over multiple lines.\n",
      "default": "{\n  \"bar\": {\n    \"bar\": \"bar\",\n    \"foo\": \"bar\"\n  },\n  \"buzz\": [\n    \"fizz\",\n    \"buzz\"\n  ],\n  \"fizz\": [],\n  \"foo\": {\n    \"bar\": \"foo\",\n    \"foo\": \"foo\"\n  },\n  \"name\": \"hello\"\n}"
    },
    {
      "name": "map-1",
      "type": "map",
      "description": "It's map number one.",
      "default": "{\n  \"a\": 1,\n  \"b\": 2,\n  \"c\": 3\n}"
    },
    {
      "name": "map-2",
      "type": "map",
      "description": "It's map number two.",
      "default": null
    },
    {
      "name": "map-3",
      "type": "map",
      "description": null,
      "default": "{}"
    },
    {
      "name": "no-escape-default-value",
      "type": "string",
      "description": "The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.",
      "default": "\"VALUE_WITH_UNDERSCORE\""
    },
    {
      "name": "string-1",
      "type": "string",
      "description": "It's string number one.",
      "default": "\"bar\""
    },
    {
      "name": "string-2",
      "type": "string",
      "description": "It's string number two.",
      "default": null
    },
    {
      "name": "string-3",
      "type": "string",
      "description": null,
      "default": "\"\""
    },
    {
      "name": "unquoted",
      "type": "any",
      "description": null,
      "default": null
    }
  ],
  "outputs": [
    {
      "name": "output-0.12",
      "description": "terraform 0.12 only"
    },
    {
      "name": "output-1",
      "description": "It's output number one."
    },
    {
      "name": "output-2",
      "description": "It's output number two."
    },
    {
      "name": "unquoted",
      "description": "It's unquoted output."
    }
  ],
  "providers": [
    {
      "name": "aws",
      "alias": null,
      "version": "\u003e= 2.15.0"
    },
    {
      "name": "aws",
      "alias": "ident",
      "version": "\u003e= 2.15.0"
    },
    {
      "name": "null",
      "alias": null,
      "version": null
    },
    {
      "name": "tls",
      "alias": null,
      "version": null
    }
  ]
}
```
