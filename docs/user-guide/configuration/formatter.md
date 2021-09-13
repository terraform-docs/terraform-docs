---
title: "formatter"
description: "formatter configuration"
menu:
  docs:
    parent: "configuration"
weight: 123
toc: true
---

Since `v0.10.0`

The following options are supported out of the box by terraform-docs and can be
used for `FORMATTER_NAME`:

- `asciidoc` <sup class="no-top">[reference]({{< ref "asciidoc" >}})</sup>
- `asciidoc document` <sup class="no-top">[reference]({{< ref "asciidoc-document" >}})</sup>
- `asciidoc table` <sup class="no-top">[reference]({{< ref "asciidoc-table" >}})</sup>
- `json` <sup class="no-top">[reference]({{< ref "json" >}})</sup>
- `markdown` <sup class="no-top">[reference]({{< ref "markdown" >}})</sup>
- `markdown document` <sup class="no-top">[reference]({{< ref "markdown-document" >}})</sup>
- `markdown table` <sup class="no-top">[reference]({{< ref "markdown-table" >}})</sup>
- `pretty` <sup class="no-top">[reference]({{< ref "pretty" >}})</sup>
- `tfvars hcl` <sup class="no-top">[reference]({{< ref "tfvars-hcl" >}})</sup>
- `tfvars json` <sup class="no-top">[reference]({{< ref "tfvars-json" >}})</sup>
- `toml` <sup class="no-top">[reference]({{< ref "toml" >}})</sup>
- `xml` <sup class="no-top">[reference]({{< ref "xml" >}})</sup>
- `yaml` <sup class="no-top">[reference]({{< ref "yaml" >}})</sup>

{{< alert type="info" >}}
Short version of formatters can also be used:

- `adoc` instead of `asciidoc`
- `md` instead of `markdown`
- `doc` instead of `document`
- `tbl` instead of `table`
{{< /alert >}}

{{< alert type="info" >}}
You need to pass name of a plugin as `formatter` in order to be able to
use the plugin. For example, if plugin binary file is called `tfdocs-format-foo`,
formatter name must be set to `foo`.
{{< /alert >}}

## Options

Available options with their default values.

```yaml
formatter: ""
```

{{< alert type="info" >}}
`formatter` is required and cannot be empty in `.terraform-docs.yml`.
{{< /alert >}}

## Examples

Format as Markdown table:

```yaml
formatter: "markdown table"
```

Format as Markdown document:

```yaml
formatter: "md doc"
```

Format as AsciiDoc document:

```yaml
formatter: "asciidoc document"
```

Format as `tfdocs-format-myplugin`:

```yaml
formatter: "myplugin"
```
