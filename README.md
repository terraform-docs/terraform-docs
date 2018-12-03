# terraform-docs

[![CircleCI](https://circleci.com/gh/segmentio/terraform-docs.svg?style=svg)](https://circleci.com/gh/segmentio/terraform-docs) [![Go Report Card](https://goreportcard.com/badge/github.com/segmentio/terraform-docs)](https://goreportcard.com/report/github.com/segmentio/terraform-docs)

> **A utility to generate documentation from Terraform modules in various output formats.**

![terraform-docs-teaser](https://raw.githubusercontent.com/segmentio/terraform-docs/media/terraform-docs-teaser.png)

## Table of Contents

- [Maintenance](#maintenance)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [License](#license)

## Maintenance

This project is no longer maintained by Segment. Instead, [Martin Etmajer](https://github.com/metmajer), unaffiliated with Segment, from [GetCloudnative](https://github.com/getcloudnative), is maintaining the project with help from these awesome [contributors](AUTHORS).

## Installation

The latest version can be installed using `go get`:

``` bash
go get github.com/segmentio/terraform-docs
```

If you are a Mac OS X user, you can use [Homebrew](https://brew.sh):

``` bash
brew install terraform-docs
```

For other platforms, please have a look at our [binary releases](https://github.com/segmentio/terraform-docs/releases).

## Getting Started

Show help information:

``` bash
terraform-docs --help
```

Generate JSON from the Terraform configuration in folder `./examples`:

```bash
terraform-docs json ./examples
```

Generate Markdown tables from the Terraform configuration in folder `./examples`:

```bash
terraform-docs markdown table ./examples
```

Generate a Markdown document from the Terraform configuration in folder `./examples`:

```bash
terraform-docs markdown document ./examples
```

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
