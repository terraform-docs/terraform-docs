---
title: "sort"
description: "sort configuration"
menu:
  docs:
    parent: "configuration"
weight: 130
toc: true
---

Since `v0.10.0`

To enable sorting of elements `sort.enabled` can be used. This will indicate
sorting is enabled or not, but consecutively type of sorting can also be specified
with `sort.by`. The following sort types are supported:

- `name` (default): name of items
- `required`: by name of inputs AND show required ones first
- `type`: type of inputs

## Options

Available options with their default values.

```yaml
sort:
  enabled: true
  by: name
```

{{< alert type="warning" >}}
As of `v0.13.0`, `sort.by` is converted from `list` to `string`.
{{< /alert >}}

The following error is an indicator that `.terraform-docs.yml` still uses
list for `sort.by`.

```text
Error: unable to decode config, 1 error(s) decoding:

* 'sort.by' expected type 'string', got unconvertible type '[]interface {}'
```

## Examples

Disable sorting:

```yaml
sort:
  enabled: false
```

Sort by name (terraform-docs `>= v0.13.0`):

```yaml
sort:
  enabled: true
  by: name
```

Sort by required (terraform-docs `< v0.13.0`):

```yaml
sort:
  enabled: true
  by:
    - required
```
