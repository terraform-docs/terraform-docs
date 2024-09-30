---
title: "asciidoc table"
description: "Generate AsciiDoc tables of inputs and outputs"
menu:
  docs:
    parent: "asciidoc"
weight: 953
toc: true
---

## Synopsis

Generate AsciiDoc tables of inputs and outputs.

```console
terraform-docs asciidoc table [PATH] [flags]
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
      --footer-from string          relative path of a file to read footer from (default "")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --hide-empty                  hide empty sections (default false)
      --indent int                  indention level of AsciiDoc sections [1, 2, 3, 4, 5] (default 2)
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
      --validation                  show Validation column or section (default true)
```

## Example

Given the [`examples`][examples] module:

```shell
terraform-docs asciidoc table --footer-from footer.md ./examples/
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

    == Requirements

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Version
    |[[requirement_terraform]] <<requirement_terraform,terraform>> |>= 0.12
    |[[requirement_aws]] <<requirement_aws,aws>> |>= 2.15.0
    |[[requirement_foo]] <<requirement_foo,foo>> |>= 1.0
    |[[requirement_random]] <<requirement_random,random>> |>= 2.2.0
    |===

    == Providers

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Version
    |[[provider_aws]] <<provider_aws,aws>> |>= 2.15.0
    |[[provider_aws.ident]] <<provider_aws.ident,aws.ident>> |>= 2.15.0
    |[[provider_foo]] <<provider_foo,foo>> |>= 1.0
    |[[provider_null]] <<provider_null,null>> |n/a
    |[[provider_tls]] <<provider_tls,tls>> |n/a
    |===

    == Modules

    [cols="a,a,a",options="header,autowidth"]
    |===
    |Name |Source |Version
    |[[module_bar]] <<module_bar,bar>> |baz |4.5.6
    |[[module_baz]] <<module_baz,baz>> |baz |4.5.6
    |[[module_foo]] <<module_foo,foo>> |bar |1.2.3
    |[[module_foobar]] <<module_foobar,foobar>> |git@github.com:module/path |v7.8.9
    |===

    == Resources

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Type
    |foo_resource.baz |resource
    |https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource[null_resource.foo] |resource
    |https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key[tls_private_key.baz] |resource
    |https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity[aws_caller_identity.current] |data source
    |https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity[aws_caller_identity.ident] |data source
    |===

    == Inputs

    [cols="a,a,a,a,a",options="header,autowidth"]
    |===
    |Name |Description |Type |Validation |Default |Required
    |[[input_bool-1]] <<input_bool-1,bool-1>>
    |It's bool number one.
    |`bool`
    |None
    |`true`
    |no

    |[[input_bool-2]] <<input_bool-2,bool-2>>
    |It's bool number two.
    |`bool`
    |None
    |`false`
    |no

    |[[input_bool-3]] <<input_bool-3,bool-3>>
    |n/a
    |`bool`
    |None
    |`true`
    |no

    |[[input_bool_default_false]] <<input_bool_default_false,bool_default_false>>
    |n/a
    |`bool`
    |None
    |`false`
    |no

    |[[input_input-with-code-block]] <<input_input-with-code-block,input-with-code-block>>
    |This is a complicated one. We need a newline.  
    And an example in a code block
    [source]
    ----
    default     = [
      "machine rack01:neptune"
    ]
    ----

    |`list`
    |None
    |

    [source]
    ----
    [
      "name rack:location"
    ]
    ----

    |no

    |[[input_input-with-pipe]] <<input_input-with-pipe,input-with-pipe>>
    |It includes v1 \| v2 \| v3
    |`string`
    |None
    |`"v1"`
    |no

    |[[input_input_with_underscores]] <<input_input_with_underscores,input_with_underscores>>
    |A variable with underscores.
    |`any`
    |None
    |n/a
    |yes

    |[[input_list-1]] <<input_list-1,list-1>>
    |It's list number one.
    |`list`
    |None
    |

    [source]
    ----
    [
      "a",
      "b",
      "c"
    ]
    ----

    |no

    |[[input_list-2]] <<input_list-2,list-2>>
    |It's list number two.
    |`list`
    |None
    |n/a
    |yes

    |[[input_list-3]] <<input_list-3,list-3>>
    |n/a
    |`list`
    |None
    |`[]`
    |no

    |[[input_list_default_empty]] <<input_list_default_empty,list_default_empty>>
    |n/a
    |`list(string)`
    |None
    |`[]`
    |no

    |[[input_long_type]] <<input_long_type,long_type>>
    |This description is itself markdown.

    It spans over multiple lines.

    |

    [source]
    ----
    object({
        name = string,
        foo  = object({ foo = string, bar = string }),
        bar  = object({ foo = string, bar = string }),
        fizz = list(string),
        buzz = list(string)
      })
    ----

    |None
    |

    [source]
    ----
    {
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
    }
    ----

    |no

    |[[input_map-1]] <<input_map-1,map-1>>
    |It's map number one.
    |`map`
    |None
    |

    [source]
    ----
    {
      "a": 1,
      "b": 2,
      "c": 3
    }
    ----

    |no

    |[[input_map-2]] <<input_map-2,map-2>>
    |It's map number two.
    |`map`
    |None
    |n/a
    |yes

    |[[input_map-3]] <<input_map-3,map-3>>
    |n/a
    |`map`
    |None
    |`{}`
    |no

    |[[input_no-escape-default-value]] <<input_no-escape-default-value,no-escape-default-value>>
    |The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.
    |`string`
    |None
    |`"VALUE_WITH_UNDERSCORE"`
    |no

    |[[input_number-1]] <<input_number-1,number-1>>
    |It's number number one.
    |`number`
    |None
    |`42`
    |no

    |[[input_number-2]] <<input_number-2,number-2>>
    |It's number number two.
    |`number`
    |None
    |n/a
    |yes

    |[[input_number-3]] <<input_number-3,number-3>>
    |n/a
    |`number`
    |None
    |`"19"`
    |no

    |[[input_number-4]] <<input_number-4,number-4>>
    |n/a
    |`number`
    |None
    |`15.75`
    |no

    |[[input_number_default_zero]] <<input_number_default_zero,number_default_zero>>
    |n/a
    |`number`
    |None
    |`0`
    |no

    |[[input_object_default_empty]] <<input_object_default_empty,object_default_empty>>
    |n/a
    |`object({})`
    |None
    |`{}`
    |no

    |[[input_string-1]] <<input_string-1,string-1>>
    |It's string number one.
    |`string`
    |None
    |`"bar"`
    |no

    |[[input_string-2]] <<input_string-2,string-2>>
    |It's string number two.
    |`string`
    |None
    |n/a
    |yes

    |[[input_string-3]] <<input_string-3,string-3>>
    |n/a
    |`string`
    |None
    |`""`
    |no

    |[[input_string-special-chars]] <<input_string-special-chars,string-special-chars>>
    |n/a
    |`string`
    |None
    |`"\\.<>[]{}_-"`
    |no

    |[[input_string_default_empty]] <<input_string_default_empty,string_default_empty>>
    |n/a
    |`string`
    |None
    |`""`
    |no

    |[[input_string_default_null]] <<input_string_default_null,string_default_null>>
    |n/a
    |`string`
    |None
    |`null`
    |no

    |[[input_string_no_default]] <<input_string_no_default,string_no_default>>
    |n/a
    |`string`
    |None
    |n/a
    |yes

    |[[input_unquoted]] <<input_unquoted,unquoted>>
    |n/a
    |`any`
    |None
    |n/a
    |yes

    |[[input_variable_with_no_validation]] <<input_variable_with_no_validation,variable_with_no_validation>>
    |This variable has no validation
    |`string`
    |None
    |`""`
    |no

    |[[input_variable_with_one_validation]] <<input_variable_with_one_validation,variable_with_one_validation>>
    |This variable has one validation
    |`string`
    |var.variable_with_one_validation must be empty or 10 characters long.
    |`""`
    |no

    |[[input_variable_with_two_validations]] <<input_variable_with_two_validations,variable_with_two_validations>>
    |This variable has two validations
    |`string`
    |var.variable_with_two_validations must be 10 characters long.+var.variable_with_two_validations must start with 'magic'.
    |n/a
    |yes

    |[[input_with-url]] <<input_with-url,with-url>>
    |The description contains url. https://www.domain.com/foo/bar_baz.html
    |`string`
    |None
    |`""`
    |no

    |===

    == Outputs

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Description
    |[[output_output-0.12]] <<output_output-0.12,output-0.12>> |terraform 0.12 only
    |[[output_output-1]] <<output_output-1,output-1>> |It's output number one.
    |[[output_output-2]] <<output_output-2,output-2>> |It's output number two.
    |[[output_unquoted]] <<output_unquoted,unquoted>> |It's unquoted output.
    |===

    ## This is an example of a footer

    It looks exactly like a header, but is placed at the end of the document

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
