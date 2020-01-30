# Markdown Table

Generate Markdown tables of inputs and outputs.

## Usage

```text
Usage:
  terraform-docs markdown table [PATH] [flags]

Aliases:
  table, tbl

Flags:
  -h, --help   help for table

Global Flags:
      --indent int                     indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
      --no-escape                      do not escape special characters
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-required                    do not show "Required" column or section
      --no-sort                        do no sort items
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

## Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs markdown table ./examples/
```

generates the following output:

    Usage:

    Example of 'foo\_bar' module in `foo_bar.tf`.

    ```hcl
    module "foo_bar" {
        source = "github.com/foo/bar"

        id   = "1234567890"
        name = "baz"

        zones = ["us-east-1", "us-west-1"]

        tags = {
            Name         = "baz"
            Created-By   = "first.last@email.com"
            Date-Created = "20180101"
        }
    }
    ```

    Here is some trailing text after code block,
    followed by another line of text.

    ## Providers

    | Name | Version |
    |------|---------|
    | aws | >= 2.15.0 |
    | aws.ident | >= 2.15.0 |
    | null | n/a |
    | tls | n/a |

    ## Inputs

    | Name | Description | Type | Default | Required |
    |------|-------------|------|---------|:-----:|
    | input-with-code-block | This is a complicated one. We need a newline.<br>And an example in a code block<pre>default     = [<br>  "machine rack01:neptune"<br>]</pre> | `list` | <pre>[<br>  "name rack:location"<br>]</pre> | no |
    | input-with-pipe | It includes v1 \| v2 \| v3 | `string` | `"v1"` | no |
    | input\_with\_underscores | A variable with underscores. | `any` | n/a | yes |
    | list-1 | It's list number one. | `list` | <pre>[<br>  "a",<br>  "b",<br>  "c"<br>]</pre> | no |
    | list-2 | It's list number two. | `list` | n/a | yes |
    | list-3 | n/a | `list` | `[]` | no |
    | long\_type | This description is itself markdown.<br><br>It spans over multiple lines. | <pre>object({<br>    name = string,<br>    foo  = object({ foo = string, bar = string }),<br>    bar  = object({ foo = string, bar = string }),<br>    fizz = list(string),<br>    buzz = list(string)<br>  })</pre> | <pre>{<br>  "bar": {<br>    "bar": "bar",<br>    "foo": "bar"<br>  },<br>  "buzz": [<br>    "fizz",<br>    "buzz"<br>  ],<br>  "fizz": [],<br>  "foo": {<br>    "bar": "foo",<br>    "foo": "foo"<br>  },<br>  "name": "hello"<br>}</pre> | no |
    | map-1 | It's map number one. | `map` | <pre>{<br>  "a": 1,<br>  "b": 2,<br>  "c": 3<br>}</pre> | no |
    | map-2 | It's map number two. | `map` | n/a | yes |
    | map-3 | n/a | `map` | `{}` | no |
    | no-escape-default-value | The description contains `something_with_underscore`. Defaults to 'VALUE\_WITH\_UNDERSCORE'. | `string` | `"VALUE_WITH_UNDERSCORE"` | no |
    | string-1 | It's string number one. | `string` | `"bar"` | no |
    | string-2 | It's string number two. | `string` | n/a | yes |
    | string-3 | n/a | `string` | `""` | no |
    | unquoted | n/a | `any` | n/a | yes |
    | with-url | The description contains url. https://www.domain.com/foo/bar_baz.html | `string` | `""` | no |

    ## Outputs

    | Name | Description |
    |------|-------------|
    | output-0.12 | terraform 0.12 only |
    | output-1 | It's output number one. |
    | output-2 | It's output number two. |
    | unquoted | It's unquoted output. |
