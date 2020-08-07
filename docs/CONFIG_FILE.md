# Config File Reference

All available options for `.terraform-docs.yml`. Note that not all of them can be used at the same time (e.g. `sections.hide` and `sections.show`)

```yaml
formatter: <FORMATTER_NAME>
header-from: main.tf

sections:
  hide-all: false
  hide:
    - header
    - inputs
    - outputs
    - providers
    - requirements
  show-all: true
  show:
    - header
    - inputs
    - outputs
    - providers
    - requirements

output-values:
  enabled: false
  from: ""

sort:
  enabled: true
  by:
    - required
    - type

settings:
  color: true
  escape: true
  indent: 2
  required: true
  sensitive: true
```

Available options for `FORMATTER_NAME` are:

- `asciidoc`
- `asciidoc document`
- `asciidoc table`
- `json`
- `markdown`
- `markdown document`
- `markdown table`
- `pretty`
- `tfvars hcl`
- `tfvars json`
- `toml`
- `xml`
- `yaml`
