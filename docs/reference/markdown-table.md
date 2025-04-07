---
title: "markdown table"
description: "Generate Markdown tables of inputs and outputs"
menu:
  docs:
    parent: "markdown"
weight: 957
toc: true
---

## Synopsis

Generate Markdown tables of inputs and outputs.

```console
terraform-docs markdown table [PATH] [flags]
```

## Options

```console
  -h, --help   help for table
```

## Inherited Options

```console
      --anchor                      create anchor links (default true)
  -c, --config string               config file name (default ".terraform-docs.yml")
      --default                     show Default column or section (default true)
      --escape                      escape special characters (default true)
      --footer-from string          relative path of a file to read footer from (default "")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --hide-empty                  hide empty sections (default false)
      --html                        use HTML tags in genereted output (default true)
      --indent int                  indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
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
      --required                    show Required column or section (default true)
      --sensitive                   show Sensitive column or section (default true)
      --show strings                show section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --sort                        sort items (default true)
      --sort-by string              sort items by criteria [name, required, type] (default "name")
      --type                        show Type column or section (default true)
```

## Example

Given the [`examples`][examples] module:

```shell
terraform-docs markdown table --footer-from footer.md ./examples/
```

generates the following output:

    Usage:

    Example of 'foo\_bar' module in `foo_bar.tf`.

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

    ## Requirements

    | Name | Version |
    |------|---------|
    | <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 0.12 |
    | <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 2.15.0 |
    | <a name="requirement_foo"></a> [foo](#requirement\_foo) | >= 1.0 |
    | <a name="requirement_random"></a> [random](#requirement\_random) | >= 2.2.0 |

    ## Providers

    | Name | Version |
    |------|---------|
    | <a name="provider_aws"></a> [aws](#provider\_aws) | >= 2.15.0 |
    | <a name="provider_aws.ident"></a> [aws.ident](#provider\_aws.ident) | >= 2.15.0 |
    | <a name="provider_foo"></a> [foo](#provider\_foo) | >= 1.0 |
    | <a name="provider_null"></a> [null](#provider\_null) | n/a |
    | <a name="provider_tls"></a> [tls](#provider\_tls) | n/a |

    ## Modules

    | Name | Source | Version |
    |------|--------|---------|
    | <a name="module_bar"></a> [bar](#module\_bar) | baz | 4.5.6 |
    | <a name="module_baz"></a> [baz](#module\_baz) | baz | 4.5.6 |
    | <a name="module_foo"></a> [foo](#module\_foo) | bar | 1.2.3 |
    | <a name="module_foobar"></a> [foobar](#module\_foobar) | git@github.com:module/path | v7.8.9 |

    ## Resources

    | Name | Type |
    |------|------|
    | foo_resource.baz | resource |
    | [null_resource.foo](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) | resource |
    | [tls_private_key.baz](https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key) | resource |
    | [aws_caller_identity.current](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |
    | [aws_caller_identity.ident](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity) | data source |

    ## Inputs

    | Name | Description | Type | Default | Required |
    |------|-------------|------|---------|:--------:|
    | <a name="input_bool-1"></a> [bool-1](#input\_bool-1) | It's bool number one. | `bool` | `true` | no |
    | <a name="input_bool-2"></a> [bool-2](#input\_bool-2) | It's bool number two. | `bool` | `false` | no |
    | <a name="input_bool-3"></a> [bool-3](#input\_bool-3) | n/a | `bool` | `true` | no |
    | <a name="input_bool_default_false"></a> [bool\_default\_false](#input\_bool\_default\_false) | n/a | `bool` | `false` | no |
    | <a name="input_input-with-code-block"></a> [input-with-code-block](#input\_input-with-code-block) | This is a complicated one. We need a newline.<br/>And an example in a code block<pre>default     = [<br/>  "machine rack01:neptune"<br/>]</pre> | `list` | <pre>[<br/>  "name rack:location"<br/>]</pre> | no |
    | <a name="input_input-with-pipe"></a> [input-with-pipe](#input\_input-with-pipe) | It includes v1 \| v2 \| v3 | `string` | `"v1"` | no |
    | <a name="input_input_with_underscores"></a> [input\_with\_underscores](#input\_input\_with\_underscores) | A variable with underscores. | `any` | n/a | yes |
    | <a name="input_list-1"></a> [list-1](#input\_list-1) | It's list number one. | `list` | <pre>[<br/>  "a",<br/>  "b",<br/>  "c"<br/>]</pre> | no |
    | <a name="input_list-2"></a> [list-2](#input\_list-2) | It's list number two. | `list` | n/a | yes |
    | <a name="input_list-3"></a> [list-3](#input\_list-3) | n/a | `list` | `[]` | no |
    | <a name="input_list_default_empty"></a> [list\_default\_empty](#input\_list\_default\_empty) | n/a | `list(string)` | `[]` | no |
    | <a name="input_long_type"></a> [long\_type](#input\_long\_type) | This description is itself markdown.<br/><br/>It spans over multiple lines. | <pre>object({<br/>    name = string,<br/>    foo  = object({ foo = string, bar = string }),<br/>    bar  = object({ foo = string, bar = string }),<br/>    fizz = list(string),<br/>    buzz = list(string)<br/>  })</pre> | <pre>{<br/>  "bar": {<br/>    "bar": "bar",<br/>    "foo": "bar"<br/>  },<br/>  "buzz": [<br/>    "fizz",<br/>    "buzz"<br/>  ],<br/>  "fizz": [],<br/>  "foo": {<br/>    "bar": "foo",<br/>    "foo": "foo"<br/>  },<br/>  "name": "hello"<br/>}</pre> | no |
    | <a name="input_map-1"></a> [map-1](#input\_map-1) | It's map number one. | `map` | <pre>{<br/>  "a": 1,<br/>  "b": 2,<br/>  "c": 3<br/>}</pre> | no |
    | <a name="input_map-2"></a> [map-2](#input\_map-2) | It's map number two. | `map` | n/a | yes |
    | <a name="input_map-3"></a> [map-3](#input\_map-3) | n/a | `map` | `{}` | no |
    | <a name="input_no-escape-default-value"></a> [no-escape-default-value](#input\_no-escape-default-value) | The description contains `something_with_underscore`. Defaults to 'VALUE\_WITH\_UNDERSCORE'. | `string` | `"VALUE_WITH_UNDERSCORE"` | no |
    | <a name="input_number-1"></a> [number-1](#input\_number-1) | It's number number one. | `number` | `42` | no |
    | <a name="input_number-2"></a> [number-2](#input\_number-2) | It's number number two. | `number` | n/a | yes |
    | <a name="input_number-3"></a> [number-3](#input\_number-3) | n/a | `number` | `"19"` | no |
    | <a name="input_number-4"></a> [number-4](#input\_number-4) | n/a | `number` | `15.75` | no |
    | <a name="input_number_default_zero"></a> [number\_default\_zero](#input\_number\_default\_zero) | n/a | `number` | `0` | no |
    | <a name="input_object_default_empty"></a> [object\_default\_empty](#input\_object\_default\_empty) | n/a | `object({})` | `{}` | no |
    | <a name="input_string-1"></a> [string-1](#input\_string-1) | It's string number one. | `string` | `"bar"` | no |
    | <a name="input_string-2"></a> [string-2](#input\_string-2) | It's string number two. | `string` | n/a | yes |
    | <a name="input_string-3"></a> [string-3](#input\_string-3) | n/a | `string` | `""` | no |
    | <a name="input_string-special-chars"></a> [string-special-chars](#input\_string-special-chars) | n/a | `string` | `"\\.<>[]{}_-"` | no |
    | <a name="input_string_default_empty"></a> [string\_default\_empty](#input\_string\_default\_empty) | n/a | `string` | `""` | no |
    | <a name="input_string_default_null"></a> [string\_default\_null](#input\_string\_default\_null) | n/a | `string` | `null` | no |
    | <a name="input_string_no_default"></a> [string\_no\_default](#input\_string\_no\_default) | n/a | `string` | n/a | yes |
    | <a name="input_unquoted"></a> [unquoted](#input\_unquoted) | n/a | `any` | n/a | yes |
    | <a name="input_with-url"></a> [with-url](#input\_with-url) | The description contains url. <https://www.domain.com/foo/bar_baz.html> | `string` | `""` | no |

    ## Outputs

    | Name | Description |
    |------|-------------|
    | <a name="output_output-0.12"></a> [output-0.12](#output\_output-0.12) | terraform 0.12 only |
    | <a name="output_output-1"></a> [output-1](#output\_output-1) | It's output number one. |
    | <a name="output_output-2"></a> [output-2](#output\_output-2) | It's output number two. |
    | <a name="output_unquoted"></a> [unquoted](#output\_unquoted) | It's unquoted output. |

    ## This is an example of a footer

    It looks exactly like a header, but is placed at the end of the document

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
