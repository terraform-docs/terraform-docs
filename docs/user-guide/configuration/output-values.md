---
title: "output-values"
description: "output-values configuration"
menu:
  docs:
    parent: "configuration"
weight: 126
toc: true
---

Since `v0.10.0`

Optional value field can be added to Outputs section which contains the current
value of an output variable as it is found in state via `terraform output`.

## Options

Available options with their default values.

```yaml
output-values:
  enabled: false
  from: ""
```

## Examples

First generate output values file in JSON format:

```bash
$ pwd
/path/to/module

$ terraform output --json > output_values.json
```

and then use the following to render them in the generated output:

```yaml
output-values:
  enabled: true
  from: "output_values.json"
```
