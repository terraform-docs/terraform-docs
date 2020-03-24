# terraform-config-inspect

This repository contains a helper library for extracting high-level metadata
about Terraform modules from their source code. It processes only a subset
of the information Terraform itself would process, and in return it's able
to be broadly compatible with modules written for many different versions of
Terraform.

```
$ go get github.com/hashicorp/terraform-config-inspect
```

```go
import "github.com/hashicorp/terraform-config-inspect/tfconfig"

/// ...

    module, diags := tfconfig.LoadModule(dir)
    /// ...
```

There are not yet any formal compatibility promises for the Terraform
configuration file format, so this tool can't yet promise to be compatible
with unknown future extensions to the format. However, it *should* be capable
of parsing valid configurations targeting Terraform versions between 0.10
and 0.12.

## Command Line Tool

The primary way to use this repository is as a Go library, but as a convenience
it also contains a CLI tool called `terraform-config-inspect`, installed
automatically by the `go get` command above, that allows viewing module
information in either a Markdown-like format or in JSON format.

```
$ terraform-config-inspect path/to/module

# Module `path/to/module`

Provider Requirements:
* **null:** (any version)

## Input Variables
* `a` (default `"a default"`)
* `b` (required): The b variable

## Output Values
* `a`
* `b`: I am B

## Managed Resources
* `null_resource.a` from `null`
* `null_resource.b` from `null`
```

```
$ terraform-config-inspect --json path/to/module
{
  "path": "path/to/module",
  "variables": {
    "A": {
      "name": "A",
      "default": "A default",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 1
      }
    },
    "B": {
      "name": "B",
      "description": "The B variable",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 5
      }
    }
  },
  "outputs": {
    "A": {
      "name": "A",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 9
      }
    },
    "B": {
      "name": "B",
      "description": "I am B",
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 13
      }
    }
  },
  "required_providers": {
    "null": []
  },
  "managed_resources": {
    "null_resource.A": {
      "mode": "managed",
      "type": "null_resource",
      "name": "A",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 18
      }
    },
    "null_resource.B": {
      "mode": "managed",
      "type": "null_resource",
      "name": "B",
      "provider": {
        "name": "null"
      },
      "pos": {
        "filename": "path/to/module/basics.tf",
        "line": 19
      }
    }
  },
  "data_resources": {},
  "module_calls": {}
}
```

## Contributing

This library and tool are intentionally focused on only extracting simple
top-level metadata about a single Terraform module. This is to reduce the
maintenance burden of keeping this codebase synchronized with changes to
Terraform itself: the features extracted by this package are unlikely to change
significantly in future versions.

For that reason, **we cannot accept external PRs for this codebase that add support for additional Terraform language features**.

We would be happy to review PRs to fix bugs in existing functionality or to
improve the usability of the Go package API, however. We will be hesitant about
any breaking changes to the API, since this library is used by a number of
existing tools and systems.

To work on this codebase you will need a recent version of Go installed. Please
ensure all files match the formatting rules applied by `go fmt` and that all
unit tests are passing.
