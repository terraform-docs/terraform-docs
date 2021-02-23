---
title: "tfvars"
description: "Generate terraform.tfvars of inputs."
menu:
  docs:
    parent: "terraform-docs"
weight: 959
toc: true
---

## Synopsis

Generate terraform.tfvars of inputs.

## Options

```console
  -h, --help   help for tfvars
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

- [terraform-docs tfvars hcl]({{< ref "tfvars-hcl" >}})
- [terraform-docs tfvars json]({{< ref "tfvars-json" >}})
