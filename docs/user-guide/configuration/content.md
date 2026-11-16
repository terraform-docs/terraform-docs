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

These variables are the generated output of individual sections in the selected
formatter. For example `{{ .Inputs }}` is Markdown Table representation of _inputs_
when formatter is set to `markdown table`.

{{< alert type="info" >}}
Sections visibility (i.e. `sections.show` and `sections.hide`) takes precedence
over the `content`.
{{< /alert >}}

`content` also has the following function:

- `{{ include "relative/path/to/file" }}`

Additionally there's also one extra special variable available to the `content`:

- `{{ .Module }}`

As opposed to the other variables mentioned above, which are generated sections
based on a selected formatter, the `{{ .Module }}` variable is just a `struct`
representing a [Terraform module].

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

  and even in between sections. also spaces will be preserved:

  - item 1
    - item 1-1
    - item 1-2
  - item 2
  - item 3

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
rendered because `providers` is not set to be shown in `sections.show`:

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

Building highly complex and highly customized content using `{{ .Module }}` struct:

```yaml
content: |-
  ## Resources

  {{ range .Module.Resources }}
  - {{ .GetMode }}.{{ .Spec }} ({{ .Position.Filename }}#{{ .Position.Line }})
  {{- end }}
```

[Terraform module]: https://pkg.go.dev/github.com/rquadling/terraform-docs/terraform#Module
