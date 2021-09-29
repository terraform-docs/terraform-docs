---
title: "sections"
description: "sections configuration"
menu:
  docs:
    parent: "configuration"
weight: 128
toc: true
---

Since `v0.10.0`

The following options are supported and can be used for `sections.show` and
`sections.hide`:

- `all` <sup class="no-top">(since v0.15.0)</sup>
- `data-sources` <sup class="no-top">(since v0.13.0)</sup>
- `header`
- `footer` <sup class="no-top">(since v0.12.0)</sup>
- `inputs`
- `modules` <sup class="no-top">(since v0.11.0)</sup>
- `outputs`
- `providers`
- `requirements`
- `resources` <sup class="no-top">(since v0.11.0)</sup>

{{< alert type="warning" >}}
The following options cannot be used together:

- `sections.hide` and `sections.show`
- `sections.hide-all` and `sections.show-all`
- `sections.hide-all` and `sections.hide`
- `sections.show-all` and `sections.show`
{{< /alert >}}

{{< alert type="info" >}}
As of `v0.13.0`, `sections.hide-all` and `sections.show-all` are deprecated
in favor of explicit use of `sections.hide` and `sections.show`, and they are removed
as of `v0.15.0`.
{{< /alert >}}

## Options

Available options with their default values.

```yaml
sections:
  hide: []
  show: []

  hide-all: false # deprecated in v0.13.0, removed in v0.15.0
  show-all: true  # deprecated in v0.13.0, removed in v0.15.0
```

## Examples

Show only `providers`, `inputs`, and `outputs`.

```yaml
sections:
  show:
    - providers
    - inputs
    - outputs
```

Show everything except `providers`.

```yaml
sections:
  hide:
    - providers
```
