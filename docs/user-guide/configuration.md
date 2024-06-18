---
title: "Configuration"
description: "terraform-docs configuration file, i.e. .terraform-docs.yml"
menu:
  docs:
    parent: "user-guide"
    identifier: "configuration"
    params:
      collapse: true
weight: 120
toc: true
---

The `terraform-docs` configuration file uses the [yaml format](https://yaml.org/) in order to override any default behaviors.
This is a convenient way to share the configuration amongst teammates, CI, or other toolings.

terraform-docs will locate any available configuration file without needing to explicitly pass the `--config` flag.

The default name of the configuration file is `.terraform-docs.yml`.
The path order for locating it is:

1. root of module directory
1. `.config/` folder at root of module directory <sup class="no-top">(since v0.15.0)</sup>
1. current directory
1. `.config/` folder at current directory <sup class="no-top">(since v0.15.0)</sup>
1. `$HOME/.tfdocs.d/`

if `.terraform-docs.yml` is found in any of the folders above, that will take
precedence and will override the other ones.

Here is an example for how your terraform project file structure might look, and where the `.terraform-docs.yml` file can be placed:

```bash
$ tree
.
├── main.tf
├── ...
├── ...
└── .terraform-docs.yml

$ terraform-docs .
```

To use an alternative configuration file name or path you
can use the `-c` or `--config` flag.

Or you can use a config file with any arbitrary name:

```bash
$ tree
.
├── main.tf
├── ...
├── ...
└── .tfdocs-config.yml

$ terraform-docs -c .tfdocs-config.yml .
```

{{< alert type="primary" >}}
Values passed directly as CLI flags will override all of the above.
{{< /alert >}}

## Options

Since `v0.10.0`

Below is a complete list of options that can be used with `terraform-docs`, with their
default values.

```yaml
formatter: "" # this is required

version: ""

header-from: main.tf
footer-from: ""

recursive:
  enabled: false
  path: modules
  include-main: true

sections:
  hide: []
  show: []

  hide-all: false # deprecated in v0.13.0, removed in v0.15.0
  show-all: true  # deprecated in v0.13.0, removed in v0.15.0

content: ""

output:
  file: ""
  mode: inject
  template: |-
    <!-- BEGIN_TF_DOCS -->
    {{ .Content }}
    <!-- END_TF_DOCS -->

output-values:
  enabled: false
  from: ""

sort:
  enabled: true
  by: name

settings:
  anchor: true
  color: true
  default: true
  description: false
  escape: true
  hide-empty: false
  html: true
  indent: 2
  lockfile: true
  read-comments: true
  required: true
  sensitive: true
  type: true
```

{{< alert type="info" >}}
`formatter` is the only required option.
{{< /alert >}}

## Usage

As of `v0.13.0`, `--config` flag accepts both relative and absolute paths.

```bash
$ pwd
/path/to/parent/folder

$ tree
.
├── module-a
│   └── main.tf
├── module-b
│   └── main.tf
├── ...
└── .terraform-docs.yml

# executing from parent
$ terraform-docs -c .terraform-docs.yml module-a/

# executing from child
$ cd module-a/
$ terraform-docs -c ../.terraform-docs.yml .

# or an absolute path
$ terraform-docs -c /path/to/parent/folder/.terraform-docs.yml .
```
