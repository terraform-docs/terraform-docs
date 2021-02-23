---
title: "markdown"
description: "Generate Markdown of inputs and outputs."
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
      --escape       escape special characters (default true)
  -h, --help         help for markdown
      --indent int   indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --required     show Required column or section (default true)
      --sensitive    show Sensitive column or section (default true)
```

## Inherited Options

```console
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, modules, outputs, providers, requirements, resources]
      --hide-all                    hide all sections (default false)
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --show strings                show section [header, inputs, modules, outputs, providers, requirements, resources]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

## Subcommands

- [terraform-docs markdown document]({{< ref "markdown-document" >}})
- [terraform-docs markdown table]({{< ref "markdown-table" >}})
