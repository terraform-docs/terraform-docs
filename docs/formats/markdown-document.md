## terraform-docs markdown document

Generate Markdown document of inputs and outputs

### Synopsis

Generate Markdown document of inputs and outputs

```
terraform-docs markdown document [PATH] [flags]
```

### Options

```
  -h, --help   help for document
```

### Options inherited from parent commands

```
  -c, --config string               config file name (default ".terraform-docs.yml")
      --escape                      escape special characters (default true)
      --header-from string          relative path of a file to read header from (default "main.tf")
      --hide strings                hide section [header, inputs, outputs, providers, requirements]
      --hide-all                    hide all sections (default false)
      --indent int                  indention level of Markdown sections [1, 2, 3, 4, 5] (default 2)
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
terraform-docs markdown document ./examples/
```

generates the following output:

    Usage:

    Example of 'foo\_bar' module in `foo_bar.tf`.

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

    ## Requirements

    The following requirements are needed by this module:

    - terraform (>= 0.12)

    - aws (>= 2.15.0)

    - random (>= 2.2.0)

    ## Providers

    The following providers are used by this module:

    - aws (>= 2.15.0)

    - aws.ident (>= 2.15.0)

    - null

    - tls

    ## Required Inputs

    The following input variables are required:

    ### input\_with\_underscores

    Description: A variable with underscores.

    Type: `any`

    ### list-2

    Description: It's list number two.

    Type: `list`

    ### map-2

    Description: It's map number two.

    Type: `map`

    ### number-2

    Description: It's number number two.

    Type: `number`

    ### string-2

    Description: It's string number two.

    Type: `string`

    ### string\_no\_default

    Description: n/a

    Type: `string`

    ### unquoted

    Description: n/a

    Type: `any`

    ## Optional Inputs

    The following input variables are optional (have default values):

    ### bool-1

    Description: It's bool number one.

    Type: `bool`

    Default: `true`

    ### bool-2

    Description: It's bool number two.

    Type: `bool`

    Default: `false`

    ### bool-3

    Description: n/a

    Type: `bool`

    Default: `true`

    ### bool\_default\_false

    Description: n/a

    Type: `bool`

    Default: `false`

    ### input-with-code-block

    Description: This is a complicated one. We need a newline.  
    And an example in a code block
    ```
    default     = [
      "machine rack01:neptune"
    ]
    ```

    Type: `list`

    Default:

    ```json
    [
      "name rack:location"
    ]
    ```

    ### input-with-pipe

    Description: It includes v1 \| v2 \| v3

    Type: `string`

    Default: `"v1"`

    ### list-1

    Description: It's list number one.

    Type: `list`

    Default:

    ```json
    [
      "a",
      "b",
      "c"
    ]
    ```

    ### list-3

    Description: n/a

    Type: `list`

    Default: `[]`

    ### list\_default\_empty

    Description: n/a

    Type: `list(string)`

    Default: `[]`

    ### long\_type

    Description: This description is itself markdown.

    It spans over multiple lines.

    Type:

    ```hcl
    object({
        name = string,
        foo  = object({ foo = string, bar = string }),
        bar  = object({ foo = string, bar = string }),
        fizz = list(string),
        buzz = list(string)
      })
    ```

    Default:

    ```json
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
    ```

    ### map-1

    Description: It's map number one.

    Type: `map`

    Default:

    ```json
    {
      "a": 1,
      "b": 2,
      "c": 3
    }
    ```

    ### map-3

    Description: n/a

    Type: `map`

    Default: `{}`

    ### no-escape-default-value

    Description: The description contains `something_with_underscore`. Defaults to 'VALUE\_WITH\_UNDERSCORE'.

    Type: `string`

    Default: `"VALUE_WITH_UNDERSCORE"`

    ### number-1

    Description: It's number number one.

    Type: `number`

    Default: `42`

    ### number-3

    Description: n/a

    Type: `number`

    Default: `"19"`

    ### number-4

    Description: n/a

    Type: `number`

    Default: `15.75`

    ### number\_default\_zero

    Description: n/a

    Type: `number`

    Default: `0`

    ### object\_default\_empty

    Description: n/a

    Type: `object({})`

    Default: `{}`

    ### string-1

    Description: It's string number one.

    Type: `string`

    Default: `"bar"`

    ### string-3

    Description: n/a

    Type: `string`

    Default: `""`

    ### string-special-chars

    Description: n/a

    Type: `string`

    Default: `"\\.<>[]{}_-"`

    ### string\_default\_empty

    Description: n/a

    Type: `string`

    Default: `""`

    ### string\_default\_null

    Description: n/a

    Type: `string`

    Default: `null`

    ### with-url

    Description: The description contains url. https://www.domain.com/foo/bar_baz.html

    Type: `string`

    Default: `""`

    ## Outputs

    The following outputs are exported:

    ### output-0.12

    Description: terraform 0.12 only

    ### output-1

    Description: It's output number one.

    ### output-2

    Description: It's output number two.

    ### unquoted

    Description: It's unquoted output.

###### Auto generated by spf13/cobra on 25-Nov-2020
