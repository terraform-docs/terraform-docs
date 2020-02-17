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
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

### Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs toml ./examples/
```

generates the following output:

    header = "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n\nEven inline **formatting** in _here_ is possible.\nand some [link](https://domain.com/)\n\n* list item 3\n* list item 4\n\n```hcl\nmodule \"foo_bar\" {\n  source = \"github.com/foo/bar\"\n\n  id   = \"1234567890\"\n  name = \"baz\"\n\n  zones = [\"us-east-1\", \"us-west-1\"]\n\n  tags = {\n    Name         = \"baz\"\n    Created-By   = \"first.last@email.com\"\n    Date-Created = \"20180101\"\n  }\n}\n```\n\nHere is some trailing text after code block,\nfollowed by another line of text.\n\n| Name | Description     |\n|------|-----------------|\n| Foo  | Foo description |\n| Bar  | Bar description |"

    [[inputs]]
      name = "input-with-code-block"
      type = "list"
      description = "This is a complicated one. We need a newline.  \nAnd an example in a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n"
      default = ["name rack:location"]

    [[inputs]]
      name = "input-with-pipe"
      type = "string"
      description = "It includes v1 | v2 | v3"
      default = "v1"

    [[inputs]]
      name = "input_with_underscores"
      type = "any"
      description = "A variable with underscores."

    [[inputs]]
      name = "list-1"
      type = "list"
      description = "It's list number one."
      default = ["a", "b", "c"]

    [[inputs]]
      name = "list-2"
      type = "list"
      description = "It's list number two."

    [[inputs]]
      name = "list-3"
      type = "list"
      description = ""
      default = []

    [[inputs]]
      name = "long_type"
      type = "object({\n    name = string,\n    foo  = object({ foo = string, bar = string }),\n    bar  = object({ foo = string, bar = string }),\n    fizz = list(string),\n    buzz = list(string)\n  })"
      description = "This description is itself markdown.\n\nIt spans over multiple lines.\n"
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
      [inputs.default]
        a = 1.0
        b = 2.0
        c = 3.0

    [[inputs]]
      name = "map-2"
      type = "map"
      description = "It's map number two."

    [[inputs]]
      name = "map-3"
      type = "map"
      description = ""
      [inputs.default]

    [[inputs]]
      name = "no-escape-default-value"
      type = "string"
      description = "The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'."
      default = "VALUE_WITH_UNDERSCORE"

    [[inputs]]
      name = "string-1"
      type = "string"
      description = "It's string number one."
      default = "bar"

    [[inputs]]
      name = "string-2"
      type = "string"
      description = "It's string number two."

    [[inputs]]
      name = "string-3"
      type = "string"
      description = ""
      default = ""

    [[inputs]]
      name = "unquoted"
      type = "any"
      description = ""

    [[inputs]]
      name = "with-url"
      type = "string"
      description = "The description contains url. https://www.domain.com/foo/bar_baz.html"
      default = ""

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



###### Auto generated by spf13/cobra on 17-Feb-2020
