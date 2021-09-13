---
title: "content"
description: "content configuration"
menu:
  docs:
    parent: "configuration"
weight: 121
toc: true
---

Since `v0.14.0`

Generated content can be customized further away with `content` in configuration.
If the `content` is empty the default order of sections is used.

{{< alert type="info" >}}
Compatible formatters for customized content are `asciidoc` and `markdown`. `content`
will be ignored for other formatters.
{{< /alert >}}

`content` is a Go template with following additional variables:

- `{{ .Header }}`
- `{{ .Footer }}`
- `{{ .Inputs }}`
- `{{ .Modules }}`
- `{{ .Outputs }}`
- `{{ .Providers }}`
- `{{ .Requirements }}`
- `{{ .Resources }}`

and following functions:

- `{{ include "relative/path/to/file" }}`

These variables are the generated output of individual sections in the selected
formatter. For example `{{ .Inputs }}` is Markdown Table representation of _inputs_
when formatter is set to `markdown table` and so on.

{{< alert type="info" >}}
Sections visibility (i.e. `sections.show` and `sections.hide`) takes
precedence over the `content`.
{{< /alert >}}

## Options

Available options with their default values.

```yaml
content: ""
```

## Examples

Content can be customized, rearranged. It can have arbitrary text in between
sections:

```yaml
content: |-
  Any arbitrary text can be placed anywhere in the content

  {{ .Header }}

  and even in between sections

  {{ .Providers }}

  and they don't even need to be in the default order

  {{ .Outputs }}

  {{ .Inputs }}
```

Relative files can be included in the `content`:

```yaml
content: |-
  include any relative files

  {{ include "relative/path/to/file" }}
```

`include` can be used to add example snippet code in the `content`:

````yaml
content: |-
  # Examples

  ```hcl
  {{ include "examples/foo/main.tf" }}
  ```
````

In the following example, although `{{ .Providers }}` is defined it won't be
rendered because `providers` is not set to be shown in `sections.show`.

```yaml
sections:
  show:
    - header
    - inputs
    - outputs

content: |-
  {{ .Header }}

  Some more information can go here.

  {{ .Providers }}

  {{ .Inputs }}

  {{ .Outputs }}
```
