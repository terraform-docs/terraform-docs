## terraform-docs asciidoc table

Generate AsciiDoc tables of inputs and outputs

### Synopsis

Generate AsciiDoc tables of inputs and outputs

```
terraform-docs asciidoc table [PATH] [flags]
```

### Options

```
  -h, --help   help for table
```

### Options inherited from parent commands

```
  -c, --config string               config file name (default ".terraform-docs.yml")
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, outputs, providers, requirements]
      --hide-all                    hide all sections (default false)
      --indent int                  indention level of AsciiDoc sections [1, 2, 3, 4, 5] (default 2)
      --output-values               inject output values into outputs (default false)
      --output-values-from string   inject output values from file into outputs (default "")
      --required                    show Required column or section (default true)
      --sensitive                   show Sensitive column or section (default true)
      --show strings                show section [header, inputs, outputs, providers, requirements]
      --show-all                    show all sections (default true)
      --sort                        sort items (default true)
      --sort-by-required            sort items by name and print required ones first (default false)
      --sort-by-type                sort items by type of them (default false)
```

### Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs asciidoc table ./examples/
```

generates the following output:

    Usage:

    Example of 'foo_bar' module in `foo_bar.tf`.

    - list item 1
    - list item 2

    Even inline **formatting** in _here_ is possible.  
    and some [link](https://domain.com/)

    * list item 3
    * list item 4

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

    | Name | Description     |
    |------|-----------------|
    | Foo  | Foo description |
    | Bar  | Bar description |

    == Requirements

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Version
    |terraform |>= 0.12
    |aws |>= 2.15.0
    |random |>= 2.2.0
    |===

    == Providers

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Version
    |aws |>= 2.15.0
    |aws.ident |>= 2.15.0
    |null |n/a
    |tls |n/a
    |===

    == Inputs

    [cols="a,a,a,a,a",options="header,autowidth"]
    |===
    |Name |Description |Type |Default |Required
    |bool-1
    |It's bool number one.
    |`bool`
    |`true`
    |no

    |bool-2
    |It's bool number two.
    |`bool`
    |`false`
    |no

    |bool-3
    |n/a
    |`bool`
    |`true`
    |no

    |bool_default_false
    |n/a
    |`bool`
    |`false`
    |no

    |input-with-code-block
    |This is a complicated one. We need a newline.  
    And an example in a code block
    [source]
    ----
    default     = [
      "machine rack01:neptune"
    ]
    ----

    |`list`
    |

    [source]
    ----
    [
      "name rack:location"
    ]
    ----

    |no

    |input-with-pipe
    |It includes v1 \| v2 \| v3
    |`string`
    |`"v1"`
    |no

    |input_with_underscores
    |A variable with underscores.
    |`any`
    |n/a
    |yes

    |list-1
    |It's list number one.
    |`list`
    |

    [source]
    ----
    [
      "a",
      "b",
      "c"
    ]
    ----

    |no

    |list-2
    |It's list number two.
    |`list`
    |n/a
    |yes

    |list-3
    |n/a
    |`list`
    |`[]`
    |no

    |list_default_empty
    |n/a
    |`list(string)`
    |`[]`
    |no

    |long_type
    |This description is itself markdown.

    It spans over multiple lines.

    |

    [source]
    ----
    object({
        name = string,
        foo  = object({ foo = string, bar = string }),
        bar  = object({ foo = string, bar = string }),
        fizz = list(string),
        buzz = list(string)
      })
    ----

    |

    [source]
    ----
    {
      "bar": {
        "bar": "bar",
        "foo": "bar"
      },
      "buzz": [
        "fizz",
        "buzz"
      ],
      "fizz": [],
      "foo": {
        "bar": "foo",
        "foo": "foo"
      },
      "name": "hello"
    }
    ----

    |no

    |map-1
    |It's map number one.
    |`map`
    |

    [source]
    ----
    {
      "a": 1,
      "b": 2,
      "c": 3
    }
    ----

    |no

    |map-2
    |It's map number two.
    |`map`
    |n/a
    |yes

    |map-3
    |n/a
    |`map`
    |`{}`
    |no

    |no-escape-default-value
    |The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.
    |`string`
    |`"VALUE_WITH_UNDERSCORE"`
    |no

    |number-1
    |It's number number one.
    |`number`
    |`42`
    |no

    |number-2
    |It's number number two.
    |`number`
    |n/a
    |yes

    |number-3
    |n/a
    |`number`
    |`"19"`
    |no

    |number-4
    |n/a
    |`number`
    |`15.75`
    |no

    |number_default_zero
    |n/a
    |`number`
    |`0`
    |no

    |object_default_empty
    |n/a
    |`object({})`
    |`{}`
    |no

    |string-1
    |It's string number one.
    |`string`
    |`"bar"`
    |no

    |string-2
    |It's string number two.
    |`string`
    |n/a
    |yes

    |string-3
    |n/a
    |`string`
    |`""`
    |no

    |string-special-chars
    |n/a
    |`string`
    |`"\\.<>[]{}_-"`
    |no

    |string_default_empty
    |n/a
    |`string`
    |`""`
    |no

    |string_default_null
    |n/a
    |`string`
    |`null`
    |no

    |string_no_default
    |n/a
    |`string`
    |n/a
    |yes

    |unquoted
    |n/a
    |`any`
    |n/a
    |yes

    |with-url
    |The description contains url. https://www.domain.com/foo/bar_baz.html
    |`string`
    |`""`
    |no

    |===

    == Outputs

    [cols="a,a",options="header,autowidth"]
    |===
    |Name |Description
    |output-0.12 |terraform 0.12 only
    |output-1 |It's output number one.
    |output-2 |It's output number two.
    |unquoted |It's unquoted output.
    |===

###### Auto generated by spf13/cobra on 25-Nov-2020
