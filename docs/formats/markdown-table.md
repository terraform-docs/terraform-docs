## terraform-docs markdown table

Generate Markdown tables of inputs and outputs

### Synopsis

Generate Markdown tables of inputs and outputs

```
terraform-docs markdown table [PATH] [flags]
```

### Options

```
  -h, --help   help for table
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
terraform-docs markdown table ./examples/
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

    | Name | Version |
    |------|---------|
    | terraform | >= 0.12 |
    | aws | >= 2.15.0 |
    | random | >= 2.2.0 |

    ## Providers

    | Name | Version |
    |------|---------|
    | aws | >= 2.15.0 |
    | aws.ident | >= 2.15.0 |
    | null | n/a |
    | tls | n/a |

    ## Inputs

    | Name | Description | Type | Default | Required |
    |------|-------------|------|---------|:--------:|
    | bool-1 | It's bool number one. | `bool` | `true` | no |
    | bool-2 | It's bool number two. | `bool` | `false` | no |
    | bool-3 | n/a | `bool` | `true` | no |
    | bool\_default\_false | n/a | `bool` | `false` | no |
    | input-with-code-block | This is a complicated one. We need a newline.<br>And an example in a code block<pre>default     = [<br>  &#34;machine rack01:neptune&#34;<br>]</pre> | `list` | <pre>[<br>  &#34;name rack:location&#34;<br>]</pre> | no |
    | input-with-pipe | It includes v1 \| v2 \| v3 | `string` | `"v1"` | no |
    | input\_with\_underscores | A variable with underscores. | `any` | n/a | yes |
    | list-1 | It's list number one. | `list` | <pre>[<br>  &#34;a&#34;,<br>  &#34;b&#34;,<br>  &#34;c&#34;<br>]</pre> | no |
    | list-2 | It's list number two. | `list` | n/a | yes |
    | list-3 | n/a | `list` | `[]` | no |
    | list\_default\_empty | n/a | `list(string)` | `[]` | no |
    | long\_type | This description is itself markdown.<br><br>It spans over multiple lines. | <pre>object({<br>    name = string,<br>    foo  = object({ foo = string, bar = string }),<br>    bar  = object({ foo = string, bar = string }),<br>    fizz = list(string),<br>    buzz = list(string)<br>  })</pre> | <pre>{<br>  &#34;bar&#34;: {<br>    &#34;bar&#34;: &#34;bar&#34;,<br>    &#34;foo&#34;: &#34;bar&#34;<br>  },<br>  &#34;buzz&#34;: [<br>    &#34;fizz&#34;,<br>    &#34;buzz&#34;<br>  ],<br>  &#34;fizz&#34;: [],<br>  &#34;foo&#34;: {<br>    &#34;bar&#34;: &#34;foo&#34;,<br>    &#34;foo&#34;: &#34;foo&#34;<br>  },<br>  &#34;name&#34;: &#34;hello&#34;<br>}</pre> | no |
    | map-1 | It's map number one. | `map` | <pre>{<br>  &#34;a&#34;: 1,<br>  &#34;b&#34;: 2,<br>  &#34;c&#34;: 3<br>}</pre> | no |
    | map-2 | It's map number two. | `map` | n/a | yes |
    | map-3 | n/a | `map` | `{}` | no |
    | no-escape-default-value | The description contains `something_with_underscore`. Defaults to 'VALUE\_WITH\_UNDERSCORE'. | `string` | `"VALUE_WITH_UNDERSCORE"` | no |
    | number-1 | It's number number one. | `number` | `42` | no |
    | number-2 | It's number number two. | `number` | n/a | yes |
    | number-3 | n/a | `number` | `"19"` | no |
    | number-4 | n/a | `number` | `15.75` | no |
    | number\_default\_zero | n/a | `number` | `0` | no |
    | object\_default\_empty | n/a | `object({})` | `{}` | no |
    | string-1 | It's string number one. | `string` | `"bar"` | no |
    | string-2 | It's string number two. | `string` | n/a | yes |
    | string-3 | n/a | `string` | `""` | no |
    | string-special-chars | n/a | `string` | `"\\.<>[]{}_-"` | no |
    | string\_default\_empty | n/a | `string` | `""` | no |
    | string\_default\_null | n/a | `string` | `null` | no |
    | string\_no\_default | n/a | `string` | n/a | yes |
    | unquoted | n/a | `any` | n/a | yes |
    | with-url | The description contains url. https://www.domain.com/foo/bar_baz.html | `string` | `""` | no |

    ## Outputs

    | Name | Description |
    |------|-------------|
    | output-0.12 | terraform 0.12 only |
    | output-1 | It's output number one. |
    | output-2 | It's output number two. |
    | unquoted | It's unquoted output. |

###### Auto generated by spf13/cobra on 25-Nov-2020
