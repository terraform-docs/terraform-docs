---
title: "terraform-docs"
description: "A utility to generate documentation from Terraform modules in various output formats"
menu:
  docs:
    parent: "reference"
weight: 950
toc: true
---

## Synopsis

A utility to generate documentation from Terraform modules in various output formats.

```console
terraform-docs [PATH] [flags]
```

## Options

```console
  -c, --config string               config file name (default ".terraform-docs.yml")
      --footer-from string          relative path of a file to read footer from (default "")
      --header-from string          relative path of a file to read header from (default "main.tf")
  -h, --help                        help for terraform-docs
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

- [terraform-docs asciidoc]({{< ref "asciidoc" >}})
  - [terraform-docs asciidoc document]({{< ref "asciidoc-document" >}})
  - [terraform-docs asciidoc table]({{< ref "asciidoc-table" >}})
- [terraform-docs json]({{< ref "json" >}})
- [terraform-docs markdown]({{< ref "markdown" >}})
  - [terraform-docs markdown document]({{< ref "markdown-document" >}})
  - [terraform-docs markdown table]({{< ref "markdown-table" >}})
- [terraform-docs pretty]({{< ref "pretty" >}})
- [terraform-docs tfvars]({{< ref "tfvars" >}})
  - [terraform-docs tfvars hcl]({{< ref "tfvars-hcl" >}})
  - [terraform-docs tfvars json]({{< ref "tfvars-json" >}})
- [terraform-docs toml]({{< ref "toml" >}})
- [terraform-docs xml]({{< ref "xml" >}})
- [terraform-docs yaml]({{< ref "yaml" >}})
