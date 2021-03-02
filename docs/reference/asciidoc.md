---
title: "asciidoc"
description: "Generate AsciiDoc of inputs and outputs."
menu:
  docs:
    parent: "terraform-docs"
weight: 951
toc: true
---

## Synopsis

Generate AsciiDoc of inputs and outputs.

```console
terraform-docs asciidoc [PATH] [flags]
```

## Options

```console
      --anchor       create anchor links (default true)
  -h, --help         help for asciidoc
      --indent int   indention level of AsciiDoc sections [1, 2, 3, 4, 5] (default 2)
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

- [terraform-docs asciidoc document]({{< ref "asciidoc-document" >}})
- [terraform-docs asciidoc table]({{< ref "asciidoc-table" >}})
