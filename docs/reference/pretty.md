---
title: "pretty"
description: "Generate colorized pretty of inputs and outputs"
menu:
  docs:
    parent: "terraform-docs"
weight: 958
toc: true
---

## Synopsis

Generate colorized pretty of inputs and outputs.

```console
terraform-docs pretty [PATH] [flags]
```

## Options

```console
      --color   colorize printed result (default true)
  -h, --help    help for pretty
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
terraform-docs pretty --footer-from footer.md --no-color ./examples/
```

generates the following output:

    Usage:

    Example of 'foo_bar' module in `foo_bar.tf`.

    - list item 1
    - list item 2

    Even inline **formatting** in _here_ is possible.
    and some [link](https://domain.com/)

    * list item 3
    * list item 4

    ```hcl
    module "foo_bar" {
      source = "github.com/foo/bar"

      id   = "1234567890"
      name = "baz"

      zones = ["us-east-1", "us-west-1"]

      tags = {
        Name         = "baz"
        Created-By   = "first.last@email.com"
        Date-Created = "20180101"
      }
    }
    ```

    Here is some trailing text after code block,
    followed by another line of text.

    | Name | Description     |
    |------|-----------------|
    | Foo  | Foo description |
    | Bar  | Bar description |


    requirement.terraform (>= 0.12)
    requirement.aws (>= 2.15.0)
    requirement.foo (>= 1.0)
    requirement.random (>= 2.2.0)


    provider.aws (>= 2.15.0)
    provider.aws.ident (>= 2.15.0)
    provider.foo (>= 1.0)
    provider.null
    provider.tls


    module.bar (baz,4.5.6)
    module.baz (baz,4.5.6)
    module.foo (bar,1.2.3)
    module.foobar (git@github.com:module/path,v7.8.9)


    resource.foo_resource.baz (resource)
    resource.null_resource.foo (resource) (https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource)
    resource.tls_private_key.baz (resource) (https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key)
    data.aws_caller_identity.current (data source) (https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)
    data.aws_caller_identity.ident (data source) (https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity)


    input.bool-1 (true)
    It's bool number one.

    input.bool-2 (false)
    It's bool number two.

    input.bool-3 (true)
    n/a

    input.bool_default_false (false)
    n/a

    input.input-with-code-block ([
      "name rack:location"
    ])
    This is a complicated one. We need a newline.
    And an example in a code block
    ```
    default     = [
      "machine rack01:neptune"
    ]
    ```

    input.input-with-pipe ("v1")
    It includes v1 | v2 | v3

    input.input_with_underscores (required)
    A variable with underscores.

    input.list-1 ([
      "a",
      "b",
      "c"
    ])
    It's list number one.

    input.list-2 (required)
    It's list number two.

    input.list-3 ([])
    n/a

    input.list_default_empty ([])
    n/a

    input.long_type ({
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
    })
    This description is itself markdown.

    It spans over multiple lines.

    input.map-1 ({
      "a": 1,
      "b": 2,
      "c": 3
    })
    It's map number one.

    input.map-2 (required)
    It's map number two.

    input.map-3 ({})
    n/a

    input.no-escape-default-value ("VALUE_WITH_UNDERSCORE")
    The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.

    input.number-1 (42)
    It's number number one.

    input.number-2 (required)
    It's number number two.

    input.number-3 ("19")
    n/a

    input.number-4 (15.75)
    n/a

    input.number_default_zero (0)
    n/a

    input.object_default_empty ({})
    n/a

    input.string-1 ("bar")
    It's string number one.

    input.string-2 (required)
    It's string number two.

    input.string-3 ("")
    n/a

    input.string-special-chars ("\\.<>[]{}_-")
    n/a

    input.string_default_empty ("")
    n/a

    input.string_default_null (null)
    n/a

    input.string_no_default (required)
    n/a

    input.unquoted (required)
    n/a

    input.with-url ("")
    The description contains url. <https://www.domain.com/foo/bar_baz.html>


    output.output-0.12
    terraform 0.12 only

    output.output-1
    It's output number one.

    output.output-2
    It's output number two.

    output.unquoted
    It's unquoted output.

    ## This is an example of a footer

    It looks exactly like a header, but is placed at the end of the document

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
