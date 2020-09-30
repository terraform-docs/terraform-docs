# terraform-docs

[![Build Status](https://github.com/terraform-docs/terraform-docs/workflows/build/badge.svg)](https://github.com/terraform-docs/terraform-docs/actions) [![GoDoc](https://godoc.org/github.com/terraform-docs/terraform-docs?status.svg)](https://godoc.org/github.com/terraform-docs/terraform-docs) [![Go Report Card](https://goreportcard.com/badge/github.com/terraform-docs/terraform-docs)](https://goreportcard.com/report/github.com/terraform-docs/terraform-docs) [![Codecov Report](https://codecov.io/gh/terraform-docs/terraform-docs/branch/master/graph/badge.svg)](https://codecov.io/gh/terraform-docs/terraform-docs) [![License](https://img.shields.io/github/license/terraform-docs/terraform-docs)](https://github.com/terraform-docs/terraform-docs/blob/master/LICENSE) [![Latest release](https://img.shields.io/github/v/release/terraform-docs/terraform-docs)](https://github.com/terraform-docs/terraform-docs/releases)

![terraform-docs-teaser](./images/terraform-docs-teaser.png)

## What is terraform-docs

A utility to generate documentation from Terraform modules in various output formats.

``` bash
terraform-docs asciidoc ./my-terraform-module          # generate asciidoc table
terraform-docs asciidoc table ./my-terraform-module    # generate asciidoc table
terraform-docs asciidoc document ./my-terraform-module # generate asciidoc document
terraform-docs json ./my-terraform-module              # generate json
terraform-docs markdown ./my-terraform-module          # generate markdown table
terraform-docs markdown table ./my-terraform-module    # generate markdown table
terraform-docs markdown document ./my-terraform-module # generate markdown document
terraform-docs pretty ./my-terraform-module            # generate colorized pretty
terraform-docs tfvars hcl ./my-terraform-module        # generate hcl format of terraform.tfvars
terraform-docs tfvars json ./my-terraform-module       # generate json format of terraform.tfvars
terraform-docs toml ./my-terraform-module              # generate toml
terraform-docs xml ./my-terraform-module               # generate xml
terraform-docs yaml ./my-terraform-module              # generate yaml
```

## Documentation

- **Users**
  - Read the [User Guide](./docs/USER_GUIDE.md) to learn how to use terraform-docs
  - Read the [Formats Guide](./docs/FORMATS_GUIDE.md) to learn about different output formats of terraform-docs
  - Refer to [Config File Reference](./docs/CONFIG_FILE.md) for all the available configuration options
- **Developers**
  - Read [Contributing Guide](CONTRIBUTING.md) before submitting a pull request

Visit [./docs](./docs/) for all documentation.

## Installation

The latest version can be installed using `go get`:

```bash
GO111MODULE="on" go get github.com/terraform-docs/terraform-docs@v0.10.1
```

**NOTE:** to download any version **before** `v0.9.1` (inclusive) you need to use to old module namespace (`segmentio`):

```bash
# only for v0.9.1 and before
GO111MODULE="on" go get github.com/segmentio/terraform-docs@v0.9.1
```

**NOTE:** please use the latest go to do this, we use 1.15.1 but ideally go 1.14 or greater.

This will put `terraform-docs` in `$(go env GOPATH)/bin`. If you encounter the error `terraform-docs: command not found` after installation then you may need to either add that directory to your `$PATH` as shown [here](https://golang.org/doc/code.html#GOPATH) or do a manual installation by cloning the repo and run `make build` from the repository which will put `terraform-docs` in:

```bash
$(go env GOPATH)/src/github.com/terraform-docs/terraform-docs/bin/$(uname | tr '[:upper:]' '[:lower:]')-amd64/terraform-docs
```
Stable binaries are also available on the [releases](https://github.com/terraform-docs/terraform-docs/releases) page. To install, download the binary for your platform from "Assets" and place this into your `$PATH`:

```bash
curl -Lo ./terraform-docs https://github.com/terraform-docs/terraform-docs/releases/download/0.10.1/terraform-docs-0.10.1-$(uname | tr '[:upper:]' '[:lower:]')-amd64
chmod +x ./terraform-docs
mv ./terraform-docs /some-dir-in-your-PATH/terraform-docs
```

**NOTE:** Windows releases are in `EXE` format.

If you are a Mac OS X user, you can use [Homebrew](https://brew.sh):

``` bash
brew install terraform-docs
```

Windows users can install using [Chocolatey](https://www.chocolatey.org):

``` bash
choco install terraform-docs
```

Alternatively you also can run `terraform-docs` as a container:

```bash
docker run quay.io/terraform-docs/terraform-docs:0.10.0
```

**NOTE:** Docker tag `latest` refers to _latest_ stable released version and `edge` refers to HEAD of `master` at any given point in time.

## Maintenance

This project was originally developed by [Segment](https://github.com/segmentio/) but now is no longer maintained by them. Instead, [Martin Etmajer](https://github.com/metmajer) from [GetCloudnative](https://github.com/getcloudnative) and [Khosrow Moossavi](https://github.com/khos2ow) from [CloudOps](https://github.com/cloudops) are maintaining the project with help from these awesome [contributors](AUTHORS). Note that maintainers are unaffiliated with Segment.

## License

MIT License - Copyright (c) 2020 The terraform-docs Authors.
