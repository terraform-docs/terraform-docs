## terraform-docs json

Generate JSON of inputs and outputs

### Synopsis

Generate JSON of inputs and outputs

```
terraform-docs json [PATH] [flags]
```

### Options

```
      --escape   escape special characters (default true)
  -h, --help     help for json
```

### Options inherited from parent commands

```
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, outputs, providers, requirements]
      --hide-all                    hide all sections (default false)
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [header, inputs, outputs, providers, requirements]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

### Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs json ./examples/
```

generates the following output:

    {
      "header": "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n\nEven inline **formatting** in _here_ is possible.\nand some [link](https://domain.com/)\n\n* list item 3\n* list item 4\n\n```hcl\nmodule \"foo_bar\" {\n  source = \"github.com/foo/bar\"\n\n  id   = \"1234567890\"\n  name = \"baz\"\n\n  zones = [\"us-east-1\", \"us-west-1\"]\n\n  tags = {\n    Name         = \"baz\"\n    Created-By   = \"first.last@email.com\"\n    Date-Created = \"20180101\"\n  }\n}\n```\n\nHere is some trailing text after code block,\nfollowed by another line of text.\n\n| Name | Description     |\n|------|-----------------|\n| Foo  | Foo description |\n| Bar  | Bar description |",
      "inputs": [
        {
          "name": "bool-1",
          "type": "bool",
          "description": "It's bool number one.",
          "default": true,
          "required": false
        },
        {
          "name": "bool-2",
          "type": "bool",
          "description": "It's bool number two.",
          "default": false,
          "required": false
        },
        {
          "name": "bool-3",
          "type": "bool",
          "description": null,
          "default": true,
          "required": false
        },
        {
          "name": "bool_default_false",
          "type": "bool",
          "description": null,
          "default": false,
          "required": false
        },
        {
          "name": "input-with-code-block",
          "type": "list",
          "description": "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n",
          "default": [
            "name rack:location"
          ],
          "required": false
        },
        {
          "name": "input-with-pipe",
          "type": "string",
          "description": "It includes v1 | v2 | v3",
          "default": "v1",
          "required": false
        },
        {
          "name": "input_with_underscores",
          "type": "any",
          "description": "A variable with underscores.",
          "default": null,
          "required": true
        },
        {
          "name": "list-1",
          "type": "list",
          "description": "It's list number one.",
          "default": [
            "a",
            "b",
            "c"
          ],
          "required": false
        },
        {
          "name": "list-2",
          "type": "list",
          "description": "It's list number two.",
          "default": null,
          "required": true
        },
        {
          "name": "list-3",
          "type": "list",
          "description": null,
          "default": [],
          "required": false
        },
        {
          "name": "list_default_empty",
          "type": "list(string)",
          "description": null,
          "default": [],
          "required": false
        },
        {
          "name": "long_type",
          "type": "object({\n    name = string,\n    foo  = object({ foo = string, bar = string }),\n    bar  = object({ foo = string, bar = string }),\n    fizz = list(string),\n    buzz = list(string)\n  })",
          "description": "This description is itself markdown.\n\nIt spans over multiple lines.\n",
          "default": {
            "bar": {
              "bar": "bar",
              "foo": "bar"
            },
            "buzz": [
              "fizz",
              "buzz"
            ],
            "fizz": [],
            "foo": {
              "bar": "foo",
              "foo": "foo"
            },
            "name": "hello"
          },
          "required": false
        },
        {
          "name": "map-1",
          "type": "map",
          "description": "It's map number one.",
          "default": {
            "a": 1,
            "b": 2,
            "c": 3
          },
          "required": false
        },
        {
          "name": "map-2",
          "type": "map",
          "description": "It's map number two.",
          "default": null,
          "required": true
        },
        {
          "name": "map-3",
          "type": "map",
          "description": null,
          "default": {},
          "required": false
        },
        {
          "name": "no-escape-default-value",
          "type": "string",
          "description": "The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.",
          "default": "VALUE_WITH_UNDERSCORE",
          "required": false
        },
        {
          "name": "number-1",
          "type": "number",
          "description": "It's number number one.",
          "default": 42,
          "required": false
        },
        {
          "name": "number-2",
          "type": "number",
          "description": "It's number number two.",
          "default": null,
          "required": true
        },
        {
          "name": "number-3",
          "type": "number",
          "description": null,
          "default": "19",
          "required": false
        },
        {
          "name": "number-4",
          "type": "number",
          "description": null,
          "default": 15.75,
          "required": false
        },
        {
          "name": "number_default_zero",
          "type": "number",
          "description": null,
          "default": 0,
          "required": false
        },
        {
          "name": "object_default_empty",
          "type": "object({})",
          "description": null,
          "default": {},
          "required": false
        },
        {
          "name": "string-1",
          "type": "string",
          "description": "It's string number one.",
          "default": "bar",
          "required": false
        },
        {
          "name": "string-2",
          "type": "string",
          "description": "It's string number two.",
          "default": null,
          "required": true
        },
        {
          "name": "string-3",
          "type": "string",
          "description": null,
          "default": "",
          "required": false
        },
        {
          "name": "string-special-chars",
          "type": "string",
          "description": null,
          "default": "\\.\u003c\u003e[]{}_-",
          "required": false
        },
        {
          "name": "string_default_empty",
          "type": "string",
          "description": null,
          "default": "",
          "required": false
        },
        {
          "name": "string_default_null",
          "type": "string",
          "description": null,
          "default": null,
          "required": false
        },
        {
          "name": "string_no_default",
          "type": "string",
          "description": null,
          "default": null,
          "required": true
        },
        {
          "name": "unquoted",
          "type": "any",
          "description": null,
          "default": null,
          "required": true
        },
        {
          "name": "with-url",
          "type": "string",
          "description": "The description contains url. https://www.domain.com/foo/bar_baz.html",
          "default": "",
          "required": false
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
      ],
      "requirements": [
        {
          "name": "terraform",
          "version": "\u003e= 0.12"
        },
        {
          "name": "aws",
          "version": "\u003e= 2.15.0"
        },
        {
          "name": "random",
          "version": "\u003e= 2.2.0"
        }
      ]
    }

###### Auto generated by spf13/cobra on 25-Nov-2020
