---
title: "Recursive Submodules"
description: "How to generate submodules documentation recursively with terraform-docs."
menu:
  docs:
    parent: "how-to"
weight: 206
toc: false
---

Since `v0.15.0`

Considering the file strucutre below of main module and its submodules, it is
possible to generate documentation for them main and all its submodules in one
execution, with `--recursive` flag.

{{< alert type="warning" >}}
Generating documentation recursively is allowed only with `--output-file`
set.
{{< /alert >}}

Path to find submodules can be configured with `--recursive-path` (defaults to
`modules`).

Each submodule can also have their own `.terraform-docs.yml` config file, to
override configuration from root module.

```bash
$ pwd
/path/to/module

$ tree .
.
├── README.md
├── main.tf
├── modules
│   └── my-sub-module
│       ├── README.md
│       ├── main.tf
│       ├── variables.tf
│       └── versions.tf
├── outputs.tf
├── variables.tf
└── versions.tf
```
