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

Read the [User Guide](./docs/USER_GUIDE.md) and [Formats Guide](./docs/FORMATS_GUIDE.md) for detailed documentation.

## Installation

The latest version can be installed using `go get`:

``` bash
GO111MODULE="on" go get github.com/terraform-docs/terraform-docs@v0.9.1
```

If you are a Mac OS X user, you can use [Homebrew](https://brew.sh):

``` bash
brew install terraform-docs
```

Windows users can install using [Chocolatey](https://www.chocolatey.org):

``` bash
choco install terraform-docs
```

**NOTE:** please use the latest go to do this, we use 1.14 but ideally go 1.13.5 or greater.

This will put `terraform-docs` in `$(go env GOPATH)/bin`. If you encounter the error `terraform-docs: command not found` after installation then you may need to either add that directory to your `$PATH` as shown [here](https://golang.org/doc/code.html#GOPATH) or do a manual installation by cloning the repo and run `make build` from the repository which will put `terraform-docs` in:

```bash
$(go env GOPATH)/src/github.com/terraform-docs/terraform-docs/bin/$(uname | tr '[:upper:]' '[:lower:]')-amd64/terraform-docs
```

Stable binaries are also available on the [releases](https://github.com/terraform-docs/terraform-docs/releases) page. To install, download the binary for your platform from "Assets" and place this into your `$PATH`:

```bash
curl -Lo ./terraform-docs https://github.com/terraform-docs/terraform-docs/releases/download/v0.9.1/terraform-docs-v0.9.1-$(uname | tr '[:upper:]' '[:lower:]')-amd64
chmod +x ./terraform-docs
mv ./terraform-docs /some-dir-in-your-PATH/terraform-docs
```

**NOTE:** Windows releases are in `EXE` format.

## Code Completion

The code completion for `bash` or `zsh` can be installed using:

**Note:** Shell auto-completion is not available for Windows users.

### bash

``` bash
terraform-docs completion bash > ~/.terraform-docs-completion
source ~/.terraform-docs-completion

# or simply the one-liner below
source <(terraform-docs completion bash)
```

### zsh

``` bash
terraform-docs completion zsh > /usr/local/share/zsh/site-functions/_terraform-docs
autoload -U compinit && compinit
```

To make this change permenant, the above commands can be added to your `~/.profile` file.

## Documentation

- **Users**
  - Read the [User Guide](./docs/USER_GUIDE.md) to learn how to use terraform-docs
  - Read the [Formats Guide](./docs/FORMATS_GUIDE.md) to learn about different output formats of terraform-docs
- **Developers**
  - Read [Contributing Guide](CONTRIBUTING.md) before submitting a pull request.
  - Building: not written yet
  - Releasing: not written yet

Visit [./docs](./docs/) for all documentation.

## Development Requirements

- [Go](https://golang.org/) 1.14 (ideally 1.13+)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [git-chlog](https://github.com/git-chglog/git-chglog)
- [golangci-lint](https://github.com/golangci/golangci-lint)

## Maintenance

This project is no longer maintained by Segment. Instead, [Martin Etmajer](https://github.com/metmajer) from [GetCloudnative](https://github.com/getcloudnative) and [Khosrow Moossavi](https://github.com/khos2ow) from [CloudOps](https://github.com/cloudops) are maintaining the project with help from these awesome [contributors](AUTHORS). Note that maintainers are unaffiliated with Segment.

## License

MIT License

Copyright (c) 2018 The terraform-docs Authors.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
