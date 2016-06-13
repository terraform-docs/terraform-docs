
  `tf-docs(1)` &sdot; a quick utility to generate markdown docs.

## Features

  - Generate docs for inputs and outputs
  - Generate JSON docs (for customizing presentation)
  - Generate markdown tables of inputs and outputs

## Installation

```bash
go get github.com/segmentio/tf-docs
```

## Usage

```bash

  Usage:
    tf-docs <dir>
    tf-docs md <dir>
    tf-docs -h | --help

  Examples:

    # Generate a JSON of inputs and outputs
    $ tf-docs ./my-module

    # Generate markdown tables of inputs and outputs
    $ tf-docs md ./my-module

  Options:
    -h, --help     show help information


```

## Example

Given a simple module at `./_example`:

```terraform

variable "subnet_ids" {
  description = "a comma-separated list of subnet IDs"
}

// The VPC ID.
output "vpc_id" {
  value = ""
}

```

To output JSON docs:

```bash
$ tf-docs _example
[
  {
    "Type": "input",
    "Name": "subnet_ids",
    "Value": "",
    "Default": "",
    "Description": "a comma-separated list of subnet IDs"
  },
  {
    "Type": "output",
    "Name": "vpc_id",
    "Value": "\"\"",
    "Default": "",
    "Description": "The VPC ID.\n"
  }
]
```

To output markdown docs:

```bash
$ tf-docs md _example
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
