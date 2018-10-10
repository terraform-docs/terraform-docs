## terraform-docs

[![CircleCI](https://circleci.com/gh/segmentio/terraform-docs.svg?style=svg)](https://circleci.com/gh/segmentio/terraform-docs)

A utility to generate documentation from Terraform modules.

<img width="1284" alt="screen shot 2016-06-14 at 5 38 37 pm" src="https://cloud.githubusercontent.com/assets/1661587/16049202/1ad63c16-3257-11e6-9e2c-6bb83e684ba4.png">

## Maintenance

This project is no longer maintained by Segment. Instead, [Martin Etmajer](https://github.com/metmajer), unaffiliated with Segment, from [GetCloudnative](https://github.com/getcloudnative), is maintaining the project with help from these awesome [contributors](AUTHORS).

## Features

  - View docs for inputs and outputs
  - Generate docs for inputs and outputs
  - Generate JSON docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs

## Installation

  - `go get github.com/segmentio/terraform-docs`
  - [Binaries](https://github.com/segmentio/terraform-docs/releases)
  - `brew install terraform-docs` (on macOS)

## Usage

```bash

  Usage:
    terraform-docs [--no-required] [--no-sort | --sort-inputs-by-required] [--with-aggregate-type-defaults] [json | markdown | md] <path>...
    terraform-docs -h | --help

  Examples:

    # View inputs and outputs
    $ terraform-docs ./my-module

    # View inputs and outputs for variables.tf and outputs.tf only
    $ terraform-docs variables.tf outputs.tf

    # Generate a JSON of inputs and outputs
    $ terraform-docs json ./my-module

    # Generate markdown tables of inputs and outputs
    $ terraform-docs md ./my-module

    # Generate markdown tables of inputs and outputs for the given module and ../config.tf
    $ terraform-docs md ./my-module ../config.tf

  Options:
    -h, --help                       show help information
    --no-required                    omit "Required" column when generating markdown
    --no-sort                        omit sorted rendering of inputs and ouputs
    --sort-inputs-by-required        sort inputs by name and prints required inputs first
    --with-aggregate-type-defaults   print default values of aggregate types
    --version                        print version

```

## Example

Given a simple module at `./_example`:

```tf
/**
 * This module has a variable and an output.  This text here will be output before any inputs or outputs!
 */

variable "subnet_ids" {
  description = "a comma-separated list of subnet IDs"
}

// The VPC ID.
output "vpc_id" {
  value = "vpc-5c1f55fd"
}

```

To view docs:

```bash
$ terraform-docs _example
```

To output JSON docs:

```bash
$ terraform-docs json _example
{
  "Comment": "This module has a variable and an output.  This text here will be output before any inputs or outputs!\n",
  "Inputs": [
    {
      "Name": "subnet_ids",
      "Description": "a comma-separated list of subnet IDs",
      "Default": ""
    }
  ],
  "Outputs": [
    {
      "Name": "vpc_id",
      "Description": "The VPC ID.\n"
    }
  ]
}
```

To output markdown docs:

```bash
$ terraform-docs md _example
This module has a variable and an output.  This text here will be output before any inputs or outputs!


## Inputs

| Name | Description | Default | Required |
|------|-------------|:-----:|:-----:|
| subnet_ids | a comma-separated list of subnet IDs | - | yes |

## Outputs

| Name | Description |
|------|-------------|
| vpc_id | The VPC ID. |

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
