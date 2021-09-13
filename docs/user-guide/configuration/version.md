---
title: "version"
description: "version configuration"
menu:
  docs:
    parent: "configuration"
weight: 130
toc: true
---

Since `v0.13.0`

terraform-docs version constraints is almost identical to the syntax used by
Terraform. A version constraint is a string literal containing one or more condition,
which are separated by commas.

Each condition consists of an operator and a version number. A version number is
a series of numbers separated by dots (e.g. `0.13.0`). Note that version number
should not have leading `v` in it.

Valid operators are as follow:

- `=` (or no operator): allows for exact version number.
- `!=`: exclude an exact version number.
- `>`, `>=`, `<`, and `<=`: comparisons against a specific version.
- `~>`: only the rightmost version component to increment.

## Options

Available options with their default values.

```yaml
version: ""
```

## Examples

Only allow terraform-docs version between `0.13.0` and `1.0.0`:

```yaml
version: ">= 0.13.0, < 1.0.0"
```
