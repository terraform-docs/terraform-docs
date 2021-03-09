---
title: "tfvars hcl"
description: "Generate HCL format of terraform.tfvars of inputs."
menu:
  docs:
    parent: "tfvars"
weight: 960
toc: true
---

## Synopsis

Generate HCL format of terraform.tfvars of inputs.

```console
terraform-docs tfvars hcl [PATH] [flags]
```

## Options

```console
  -h, --help   help for hcl
```

## Inherited Options

```console
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, modules, outputs, providers, requirements, resources]
      --hide-all                    hide all sections (default false)
      --output-file string          File in module directory to insert output into (default "")
      --output-mode string          Output to file method [inject, replace] (default "inject")
      --output-template string      Output template (default "<!-- BEGIN_TF_DOCS -->\n{{ .Content }}\n<!-- END_TF_DOCS -->")
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [header, inputs, modules, outputs, providers, requirements, resources]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

## Example

Given the [`examples`][examples] module:

```shell
terraform-docs tfvars hcl ./examples/
```

generates the following output:

    bool-1             = true
    bool-2             = false
    bool-3             = true
    bool_default_false = false
    input-with-code-block = [
      "name rack:location"
    ]
    input-with-pipe        = "v1"
    input_with_underscores = ""
    list-1 = [
      "a",
      "b",
      "c"
    ]
    list-2             = ""
    list-3             = []
    list_default_empty = []
    long_type = {
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
    map-1 = {
      "a": 1,
      "b": 2,
      "c": 3
    }
    map-2                   = ""
    map-3                   = {}
    no-escape-default-value = "VALUE_WITH_UNDERSCORE"
    number-1                = 42
    number-2                = ""
    number-3                = "19"
    number-4                = 15.75
    number_default_zero     = 0
    object_default_empty    = {}
    string-1                = "bar"
    string-2                = ""
    string-3                = ""
    string-special-chars    = "\\.<>[]{}_-"
    string_default_empty    = ""
    string_default_null     = ""
    string_no_default       = ""
    unquoted                = ""
    with-url                = ""

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
