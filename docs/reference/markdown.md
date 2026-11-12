---
title: "markdown"
description: "Generate Markdown of inputs and outputs"
menu:
  docs:
    parent: "terraform-docs"
weight: 955
toc: true
---

## Synopsis

Generate Markdown of inputs and outputs.

```console
terraform-docs markdown [PATH] [flags]
```

## Options

```console
      --anchor[=true|false]       create anchor links (default true)
      --atx-closed[=true|false]   close ATX style headers
      --default[=true|false]      show Default column or section (default true)
      --escape[=true|false]       escape special characters (default true)
  -h, --help[=true|false]         help for markdown
      --hide-empty[=true|false]   hide empty sections (default false)
      --html[=true|false]         use HTML tags in generated output (default true)
      --indent int                indentation level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --required[=true|false]     show Required column or section (default true)
      --sensitive[=true|false]    show Sensitive column or section (default true)
      --type[=true|false]         show Type column or section (default true)
```

## Inherited Options

```console
  -c, --config string                         config file name (default ".terraform-docs.yml")
      --footer-from string                    relative path of a file to read footer from (default "")
      --header-from string                    relative path of a file to read header from (default "main.tf")
      --hide strings                          hide section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --lockfile[=true|false]                 read .terraform.lock.hcl if exist (default true)
      --output-check[=true|false]             check if content of output file is up to date (default false)
      --output-file string                    file path to insert output into (default "")
      --output-mode string                    output to file method [inject, replace] (default "inject")
      --output-template string                output template (default "<!-- BEGIN_TF_DOCS -->\n{{ .Content }}\n<!-- END_TF_DOCS -->")
      --output-values[=true|false]            inject output values into outputs (default false)
      --output-values-from string             inject output values from file into outputs (default "")
      --read-comments[=true|false]            use comments as description when description is empty (default true)
      --recursive[=true|false]                update submodules recursively (default false)
      --recursive-exclude strings             exclude directories from recursive update
      --recursive-include-main[=true|false]   include the main module (default true)
      --recursive-path string                 submodules path to recursively update (default "modules")
      --show strings                          show section [all, data-sources, footer, header, inputs, modules, outputs, providers, requirements, resources]
      --sort[=true|false]                     sort items (default true)
      --sort-by string                        sort items by criteria [name, required, type] (default "name")
```

## Subcommands

- [terraform-docs markdown document]({{< ref "markdown-document" >}})
- [terraform-docs markdown table]({{< ref "markdown-table" >}})
