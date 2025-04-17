---
title: "toml"
description: "Generate TOML of inputs and outputs"
menu:
  docs:
    parent: "terraform-docs"
weight: 962
toc: true
---

## Synopsis

Generate TOML of inputs and outputs.

```console
terraform-docs toml [PATH] [flags]
```

## Options

```console
  -h, --help   help for toml
```

## Inherited Options

```console
  -c, --config string               config file name (default ".terraform-docs.yml")
      --footer-from string          relative path of a file to read footer from (default "")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --lockfile                    read .terraform.lock.hcl if exist (default true)
      --output-check                check if content of output file is up to date (default false)
      --output-file string          file path to insert output into (default "")
      --output-mode string          output to file method [inject, replace] (default "inject")
      --output-template string      output template (default "<!-- BEGIN_TF_DOCS -->\n{{ .Content }}\n<!-- END_TF_DOCS -->")
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --read-comments               use comments as description when description is empty (default true)
      --recursive                   update submodules recursively (default false)
      --recursive-include-main      include the main module (default true)
      --recursive-path string       submodules path to recursively update (default "modules")
      --show strings                show section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --sort                        sort items (default true)
      --sort-by string              sort items by criteria [name, required, type] (default "name")
```

## Example

Given the [`examples`][examples] module:

```shell
terraform-docs toml --footer-from footer.md ./examples/
```

generates the following output:

    header = "Usage:\n\nExample of 'foo_bar' module in `foo_bar.tf`.\n\n- list item 1\n- list item 2\n\nEven inline **formatting** in _here_ is possible.\nand some [link](https://domain.com/)\n\n* list item 3\n* list item 4\n\n```hcl\nmodule \"foo_bar\" {\n  source = \"github.com/foo/bar\"\n\n  id   = \"1234567890\"\n  name = \"baz\"\n\n  zones = [\"us-east-1\", \"us-west-1\"]\n\n  tags = {\n    Name         = \"baz\"\n    Created-By   = \"first.last@email.com\"\n    Date-Created = \"20180101\"\n  }\n}\n```\n\nHere is some trailing text after code block,\nfollowed by another line of text.\n\n| Name | Description     |\n|------|-----------------|\n| Foo  | Foo description |\n| Bar  | Bar description |"
    footer = "## This is an example of a footer\n\nIt looks exactly like a header, but is placed at the end of the document"

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
      description = "The description contains url. <https://www.domain.com/foo/bar_baz.html>"
      default = ""
      required = false

    [[modules]]
      name = "bar"
      source = "baz"
      version = "4.5.6"
      description = ""

    [[modules]]
      name = "baz"
      source = "baz"
      version = "4.5.6"
      description = ""

    [[modules]]
      name = "foo"
      source = "bar"
      version = "1.2.3"
      description = "another type of description for module foo"

    [[modules]]
      name = "foobar"
      source = "git@github.com:module/path"
      version = "v7.8.9"
      description = ""

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
      name = "foo"
      alias = ""
      version = ">= 1.0"

    [[providers]]
      name = "null"
      alias = ""
      version = ""

    [[providers]]
      name = "tls"
      alias = ""
      version = ""

    [[requirements]]
      name = "terraform"
      version = ">= 0.12"

    [[requirements]]
      name = "aws"
      version = ">= 2.15.0"

    [[requirements]]
      name = "foo"
      version = ">= 1.0"

    [[requirements]]
      name = "random"
      version = ">= 2.2.0"

    [[resources]]
      type = "resource"
      name = "baz"
      provider = "foo"
      source = "https://registry.acme.com/foo"
      mode = "managed"
      version = "latest"
      description = ""

    [[resources]]
      type = "resource"
      name = "foo"
      provider = "null"
      source = "hashicorp/null"
      mode = "managed"
      version = "latest"
      description = ""

    [[resources]]
      type = "private_key"
      name = "baz"
      provider = "tls"
      source = "hashicorp/tls"
      mode = "managed"
      version = "latest"
      description = "this description for tls_private_key.baz which can be multiline."

    [[resources]]
      type = "caller_identity"
      name = "current"
      provider = "aws"
      source = "hashicorp/aws"
      mode = "data"
      version = "latest"
      description = ""

    [[resources]]
      type = "caller_identity"
      name = "ident"
      provider = "aws"
      source = "hashicorp/aws"
      mode = "data"
      version = "latest"
      description = ""

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
