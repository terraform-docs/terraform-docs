---
title: "asciidoc document"
description: "Generate AsciiDoc document of inputs and outputs"
menu:
  docs:
    parent: "asciidoc"
weight: 952
toc: true
---

## Synopsis

Generate AsciiDoc document of inputs and outputs.

```console
terraform-docs asciidoc document [PATH] [flags]
```

## Options

```console
  -h, --help   help for document
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
terraform-docs asciidoc document --footer-from footer.md ./examples/
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

    The following requirements are needed by this module:

    - [[requirement_terraform]] <<requirement_terraform,terraform>> (>= 0.12)

    - [[requirement_aws]] <<requirement_aws,aws>> (>= 2.15.0)

    - [[requirement_foo]] <<requirement_foo,foo>> (>= 1.0)

    - [[requirement_random]] <<requirement_random,random>> (>= 2.2.0)

    == Providers

    The following providers are used by this module:

    - [[provider_aws]] <<provider_aws,aws>> (>= 2.15.0)

    - [[provider_aws.ident]] <<provider_aws.ident,aws.ident>> (>= 2.15.0)

    - [[provider_foo]] <<provider_foo,foo>> (>= 1.0)

    - [[provider_null]] <<provider_null,null>>

    - [[provider_tls]] <<provider_tls,tls>>

    == Modules

    The following Modules are called:

    === [[module_bar]] <<module_bar,bar>>

    Source: baz

    Version: 4.5.6

    === [[module_baz]] <<module_baz,baz>>

    Source: baz

    Version: 4.5.6

    === [[module_foo]] <<module_foo,foo>>

    Source: bar

    Version: 1.2.3

    === [[module_foobar]] <<module_foobar,foobar>>

    Source: git@github.com:module/path

    Version: v7.8.9

    == Resources

    The following resources are used by this module:

    - foo_resource.baz (resource)
    - https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource[null_resource.foo] (resource)
    - https://registry.terraform.io/providers/hashicorp/tls/latest/docs/resources/private_key[tls_private_key.baz] (resource)
    - https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity[aws_caller_identity.current] (data source)
    - https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/caller_identity[aws_caller_identity.ident] (data source)

    == Required Inputs

    The following input variables are required:

    === [[input_input_with_underscores]] <<input_input_with_underscores,input_with_underscores>>

    Description: A variable with underscores.

    Type: `any`

    === [[input_list-2]] <<input_list-2,list-2>>

    Description: It's list number two.

    Type: `list`

    === [[input_map-2]] <<input_map-2,map-2>>

    Description: It's map number two.

    Type: `map`

    === [[input_number-2]] <<input_number-2,number-2>>

    Description: It's number number two.

    Type: `number`

    === [[input_string-2]] <<input_string-2,string-2>>

    Description: It's string number two.

    Type: `string`

    === [[input_string_no_default]] <<input_string_no_default,string_no_default>>

    Description: n/a

    Type: `string`

    === [[input_unquoted]] <<input_unquoted,unquoted>>

    Description: n/a

    Type: `any`

    === [[input_variable_with_two_validations]] <<input_variable_with_two_validations,variable_with_two_validations>>

    Description: This variable has two validations

    Type: `string`

    Validations:

    - var.variable_with_two_validations must be 10 characters long.
    - var.variable_with_two_validations must start with 'magic'.

    == Optional Inputs

    The following input variables are optional (have default values):

    === [[input_bool-1]] <<input_bool-1,bool-1>>

    Description: It's bool number one.

    Type: `bool`

    Default: `true`

    === [[input_bool-2]] <<input_bool-2,bool-2>>

    Description: It's bool number two.

    Type: `bool`

    Default: `false`

    === [[input_bool-3]] <<input_bool-3,bool-3>>

    Description: n/a

    Type: `bool`

    Default: `true`

    === [[input_bool_default_false]] <<input_bool_default_false,bool_default_false>>

    Description: n/a

    Type: `bool`

    Default: `false`

    === [[input_input-with-code-block]] <<input_input-with-code-block,input-with-code-block>>

    Description: This is a complicated one. We need a newline.
    And an example in a code block
    ```
    default     = [
      "machine rack01:neptune"
    ]
    ```

    Type: `list`

    Default:
    [source,json]
    ----
    [
      "name rack:location"
    ]
    ----

    === [[input_input-with-pipe]] <<input_input-with-pipe,input-with-pipe>>

    Description: It includes v1 | v2 | v3

    Type: `string`

    Default: `"v1"`

    === [[input_list-1]] <<input_list-1,list-1>>

    Description: It's list number one.

    Type: `list`

    Default:
    [source,json]
    ----
    [
      "a",
      "b",
      "c"
    ]
    ----

    === [[input_list-3]] <<input_list-3,list-3>>

    Description: n/a

    Type: `list`

    Default: `[]`

    === [[input_list_default_empty]] <<input_list_default_empty,list_default_empty>>

    Description: n/a

    Type: `list(string)`

    Default: `[]`

    === [[input_long_type]] <<input_long_type,long_type>>

    Description: This description is itself markdown.

    It spans over multiple lines.

    Type:
    [source,hcl]
    ----
    object({
        name = string,
        foo  = object({ foo = string, bar = string }),
        bar  = object({ foo = string, bar = string }),
        fizz = list(string),
        buzz = list(string)
      })
    ----

    Default:
    [source,json]
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

    === [[input_map-1]] <<input_map-1,map-1>>

    Description: It's map number one.

    Type: `map`

    Default:
    [source,json]
    ----
    {
      "a": 1,
      "b": 2,
      "c": 3
    }
    ----

    === [[input_map-3]] <<input_map-3,map-3>>

    Description: n/a

    Type: `map`

    Default: `{}`

    === [[input_no-escape-default-value]] <<input_no-escape-default-value,no-escape-default-value>>

    Description: The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.

    Type: `string`

    Default: `"VALUE_WITH_UNDERSCORE"`

    === [[input_number-1]] <<input_number-1,number-1>>

    Description: It's number number one.

    Type: `number`

    Default: `42`

    === [[input_number-3]] <<input_number-3,number-3>>

    Description: n/a

    Type: `number`

    Default: `"19"`

    === [[input_number-4]] <<input_number-4,number-4>>

    Description: n/a

    Type: `number`

    Default: `15.75`

    === [[input_number_default_zero]] <<input_number_default_zero,number_default_zero>>

    Description: n/a

    Type: `number`

    Default: `0`

    === [[input_object_default_empty]] <<input_object_default_empty,object_default_empty>>

    Description: n/a

    Type: `object({})`

    Default: `{}`

    === [[input_string-1]] <<input_string-1,string-1>>

    Description: It's string number one.

    Type: `string`

    Default: `"bar"`

    === [[input_string-3]] <<input_string-3,string-3>>

    Description: n/a

    Type: `string`

    Default: `""`

    === [[input_string-special-chars]] <<input_string-special-chars,string-special-chars>>

    Description: n/a

    Type: `string`

    Default: `"\\.<>[]{}_-"`

    === [[input_string_default_empty]] <<input_string_default_empty,string_default_empty>>

    Description: n/a

    Type: `string`

    Default: `""`

    === [[input_string_default_null]] <<input_string_default_null,string_default_null>>

    Description: n/a

    Type: `string`

    Default: `null`

    === [[input_variable_with_no_validation]] <<input_variable_with_no_validation,variable_with_no_validation>>

    Description: This variable has no validation

    Type: `string`

    Default: `""`

    === [[input_variable_with_one_validation]] <<input_variable_with_one_validation,variable_with_one_validation>>

    Description: This variable has one validation

    Type: `string`

    Validations:

    - var.variable_with_one_validation must be empty or 10 characters long.

    Default: `""`

    === [[input_with-url]] <<input_with-url,with-url>>

    Description: The description contains url. https://www.domain.com/foo/bar_baz.html

    Type: `string`

    Default: `""`

    == Outputs

    The following outputs are exported:

    === [[output_output-0.12]] <<output_output-0.12,output-0.12>>

    Description: terraform 0.12 only

    === [[output_output-1]] <<output_output-1,output-1>>

    Description: It's output number one.

    === [[output_output-2]] <<output_output-2,output-2>>

    Description: It's output number two.

    === [[output_unquoted]] <<output_unquoted,unquoted>>

    Description: It's unquoted output.

    ## This is an example of a footer

    It looks exactly like a header, but is placed at the end of the document

[examples]: https://github.com/rquadling/terraform-docs/tree/master/examples
