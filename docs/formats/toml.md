## terraform-docs toml

Generate TOML of inputs and outputs

### Synopsis

Generate TOML of inputs and outputs

```
terraform-docs toml [PATH] [flags]
```

### Options

```
  -h, --help   help for toml
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
terraform-docs toml ./examples/
```

generates the following output:

    header = "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n\nEven inline **formatting** in _here_ is possible.\nand some [link](https://domain.com/)\n\n* list item 3\n* list item 4\n\n```hcl\nmodule \"foo_bar\" {\n  source = \"github.com/foo/bar\"\n\n  id   = \"1234567890\"\n  name = \"baz\"\n\n  zones = [\"us-east-1\", \"us-west-1\"]\n\n  tags = {\n    Name         = \"baz\"\n    Created-By   = \"first.last@email.com\"\n    Date-Created = \"20180101\"\n  }\n}\n```\n\nHere is some trailing text after code block,\nfollowed by another line of text.\n\n| Name | Description     |\n|------|-----------------|\n| Foo  | Foo description |\n| Bar  | Bar description |"

    [[inputs]]
      name = "bool-1"
      type = "bool"
      description = "It's bool number one."
      default = true
      required = false

    [[inputs]]
      name = "bool-2"
      type = "bool"
      description = "It's bool number two."
      default = false
      required = false

    [[inputs]]
      name = "bool-3"
      type = "bool"
      description = ""
      default = true
      required = false

    [[inputs]]
      name = "bool_default_false"
      type = "bool"
      description = ""
      default = false
      required = false

    [[inputs]]
      name = "input-with-code-block"
      type = "list"
      description = "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n"
      default = ["name rack:location"]
      required = false

    [[inputs]]
      name = "input-with-pipe"
      type = "string"
      description = "It includes v1 | v2 | v3"
      default = "v1"
      required = false

    [[inputs]]
      name = "input_with_underscores"
      type = "any"
      description = "A variable with underscores."
      required = true
      [inputs.default]

    [[inputs]]
      name = "list-1"
      type = "list"
      description = "It's list number one."
      default = ["a", "b", "c"]
      required = false

    [[inputs]]
      name = "list-2"
      type = "list"
      description = "It's list number two."
      required = true
      [inputs.default]

    [[inputs]]
      name = "list-3"
      type = "list"
      description = ""
      default = []
      required = false

    [[inputs]]
      name = "list_default_empty"
      type = "list(string)"
      description = ""
      default = []
      required = false

    [[inputs]]
      name = "long_type"
      type = "object({\n    name = string,\n    foo  = object({ foo = string, bar = string }),\n    bar  = object({ foo = string, bar = string }),\n    fizz = list(string),\n    buzz = list(string)\n  })"
      description = "This description is itself markdown.\n\nIt spans over multiple lines.\n"
      required = false
      [inputs.default]
        buzz = ["fizz", "buzz"]
        fizz = []
        name = "hello"
        [inputs.default.bar]
          bar = "bar"
          foo = "bar"
        [inputs.default.foo]
          bar = "foo"
          foo = "foo"

    [[inputs]]
      name = "map-1"
      type = "map"
      description = "It's map number one."
      required = false
      [inputs.default]
        a = 1.0
        b = 2.0
        c = 3.0

    [[inputs]]
      name = "map-2"
      type = "map"
      description = "It's map number two."
      required = true
      [inputs.default]

    [[inputs]]
      name = "map-3"
      type = "map"
      description = ""
      required = false
      [inputs.default]

    [[inputs]]
      name = "no-escape-default-value"
      type = "string"
      description = "The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'."
      default = "VALUE_WITH_UNDERSCORE"
      required = false

    [[inputs]]
      name = "number-1"
      type = "number"
      description = "It's number number one."
      default = 42.0
      required = false

    [[inputs]]
      name = "number-2"
      type = "number"
      description = "It's number number two."
      required = true
      [inputs.default]

    [[inputs]]
      name = "number-3"
      type = "number"
      description = ""
      default = "19"
      required = false

    [[inputs]]
      name = "number-4"
      type = "number"
      description = ""
      default = 15.75
      required = false

    [[inputs]]
      name = "number_default_zero"
      type = "number"
      description = ""
      default = 0.0
      required = false

    [[inputs]]
      name = "object_default_empty"
      type = "object({})"
      description = ""
      required = false
      [inputs.default]

    [[inputs]]
      name = "string-1"
      type = "string"
      description = "It's string number one."
      default = "bar"
      required = false

    [[inputs]]
      name = "string-2"
      type = "string"
      description = "It's string number two."
      required = true
      [inputs.default]

    [[inputs]]
      name = "string-3"
      type = "string"
      description = ""
      default = ""
      required = false

    [[inputs]]
      name = "string-special-chars"
      type = "string"
      description = ""
      default = "\\.<>[]{}_-"
      required = false

    [[inputs]]
      name = "string_default_empty"
      type = "string"
      description = ""
      default = ""
      required = false

    [[inputs]]
      name = "string_default_null"
      type = "string"
      description = ""
      required = false
      [inputs.default]

    [[inputs]]
      name = "string_no_default"
      type = "string"
      description = ""
      required = true
      [inputs.default]

    [[inputs]]
      name = "unquoted"
      type = "any"
      description = ""
      required = true
      [inputs.default]

    [[inputs]]
      name = "with-url"
      type = "string"
      description = "The description contains url. https://www.domain.com/foo/bar_baz.html"
      default = ""
      required = false

    [[outputs]]
      name = "output-0.12"
      description = "terraform 0.12 only"

    [[outputs]]
      name = "output-1"
      description = "It's output number one."

    [[outputs]]
      name = "output-2"
      description = "It's output number two."

    [[outputs]]
      name = "unquoted"
      description = "It's unquoted output."

    [[providers]]
      name = "aws"
      alias = ""
      version = ">= 2.15.0"

    [[providers]]
      name = "aws"
      alias = "ident"
      version = ">= 2.15.0"

    [[providers]]
      name = "null"
      alias = ""
      version = ""

    [[providers]]
      name = "tls"
      alias = ""
      version = ""

    [[requirements]]
      Name = "terraform"
      Version = ">= 0.12"

    [[requirements]]
      Name = "aws"
      Version = ">= 2.15.0"

    [[requirements]]
      Name = "random"
      Version = ">= 2.2.0"

###### Auto generated by spf13/cobra on 25-Nov-2020
