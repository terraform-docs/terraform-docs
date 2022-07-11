---
title: "recursive"
description: "recursive configuration"
menu:
  docs:
    parent: "configuration"
weight: 127
toc: true
---

Since `v0.16.0`

Documentation for main module and its submodules can be generated all in one
execution using `recursive` config. It can be enabled with `recursive.enabled: true`.

Path to find submodules can be configured with `recursive.path` (defaults to
`modules`).

{{< alert type="warning" >}}
Generating documentation recursively is allowed only with `output.file`
set.
{{< /alert >}}

Each submodule can also have their own `.terraform-docs.yml` config file, to
override configuration from root module.

## Options

Available options with their default values.

```yaml
recursive:
  enabled: false
  path: modules
  include-main: true
```

## Examples

Enable recursive mode for submodules folder.

```yaml
recursive:
  enabled: true
```

Provide alternative name of submodules folder.

```yaml
recursive:
  enabled: true
  path: submodules-folder
```

Skip the main module document, and only generate documents for submodules.

```yaml
recursive:
  enabled: true
  path: submodules-folder
  include-main: false
```
