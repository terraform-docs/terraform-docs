---
title: "tfvars json"
description: "Generate JSON format of terraform.tfvars of inputs"
menu:
  docs:
    parent: "tfvars"
weight: 961
toc: true
---

## Synopsis

Generate JSON format of terraform.tfvars of inputs.

```console
terraform-docs tfvars json [PATH] [flags]
```

## Options

```console
  -h, --help   help for json
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
terraform-docs tfvars json --footer-from footer.md ./examples/
```

generates the following output:

    {
      "bool-1": true,
      "bool-2": false,
      "bool-3": true,
      "bool_default_false": false,
      "input-with-code-block": [
        "name rack:location"
      ],
      "input-with-pipe": "v1",
      "input_with_underscores": null,
      "list-1": [
        "a",
        "b",
        "c"
      ],
      "list-2": null,
      "list-3": [],
      "list_default_empty": [],
      "long_type": {
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
      "map-1": {
        "a": 1,
        "b": 2,
        "c": 3
      },
      "map-2": null,
      "map-3": {},
      "no-escape-default-value": "VALUE_WITH_UNDERSCORE",
      "number-1": 42,
      "number-2": null,
      "number-3": "19",
      "number-4": 15.75,
      "number_default_zero": 0,
      "object_default_empty": {},
      "string-1": "bar",
      "string-2": null,
      "string-3": "",
      "string-special-chars": "\\.<>[]{}_-",
      "string_default_empty": "",
      "string_default_null": null,
      "string_no_default": null,
      "unquoted": null,
      "variable_with_no_validation": "",
      "variable_with_one_validation": "",
      "variable_with_two_validations": null,
      "with-url": ""
    }

[examples]: https://github.com/terraform-docs/terraform-docs/tree/master/examples
