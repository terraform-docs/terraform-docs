
  `terraform-docs(1)` &sdot; a quick utility to generate docs from terraform modules.

## Features

  - View docs for inputs and outputs
  - Generate docs for inputs and outputs
  - Generate JSON docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs

## Installation

```bash
go get github.com/segmentio/terraform-docs
```

## Usage

```bash

  Usage:
    terraform-docs <dir>
    terraform-docs json <dir>
    terraform-docs markdown <dir>
    terraform-docs md <dir>
    terraform-docs -h | --help

  Examples:

    # View inputs and outputs
    $ teraform-docs ./my-module

    # Generate a JSON of inputs and outputs
    $ teraform-docs json ./my-module

    # Generate markdown tables of inputs and outputs
    $ teraform-docs md ./my-module

  Options:
    -h, --help     show help information

```

## Example

Given a simple module at `./_example`:

```tf

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
  "Comment": "",
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

Released under the MIT License

(The MIT License)

Copyright (c) 2016 Segment friends@segment.com

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
