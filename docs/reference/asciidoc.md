---
title: "asciidoc"
description: "Generate AsciiDoc of inputs and outputs"
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
      --default      show Default column or section (default true)
  -h, --help         help for asciidoc
      --hide-empty   hide empty sections (default false)
      --indent int   indention level of AsciiDoc sections [1, 2, 3, 4, 5] (default 2)
      --required     show Required column or section (default true)
      --sensitive    show Sensitive column or section (default true)
      --type         show Type column or section (default true)
      --validation   show Validation column or section (default true)
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

## Subcommands

- [terraform-docs asciidoc document]({{< ref "asciidoc-document" >}})
- [terraform-docs asciidoc table]({{< ref "asciidoc-table" >}})
