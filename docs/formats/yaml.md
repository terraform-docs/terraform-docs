## terraform-docs yaml

Generate YAML of inputs and outputs

### Synopsis

Generate YAML of inputs and outputs

```
terraform-docs yaml [PATH] [flags]
```

### Options

```
  -h, --help   help for yaml
```

### Options inherited from parent commands

```
      --header-from string             relative path of a file to read header from (default "main.tf")
      --no-header                      do not show module header
      --no-inputs                      do not show inputs
      --no-outputs                     do not show outputs
      --no-providers                   do not show providers
      --no-requirements                do not show module requirements
      --no-sort                        do no sort items
      --output-values                  inject output values into outputs
      --output-values-from string      inject output values from file into outputs
      --sort-by-required               sort items by name and print required ones first
      --sort-inputs-by-required        [deprecated] use '--sort-by-required' instead
      --with-aggregate-type-defaults   [deprecated] print default values of aggregate types
```

### Example

Given the [`examples`](/examples/) module:

```shell
terraform-docs yaml ./examples/
```

generates the following output:

    header: |-
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
    inputs:
    - name: bool-1
      type: bool
      description: It's bool number one.
      default: true
    - name: bool-2
      type: bool
      description: It's bool number two.
      default: false
    - name: bool-3
      type: bool
      description: null
      default: true
    - name: bool_default_false
      type: bool
      description: null
      default: false
    - name: input-with-code-block
      type: list
      description: "This is a complicated one. We need a newline.  \nAnd an example in
        a code block\n```\ndefault     = [\n  \"machine rack01:neptune\"\n]\n```\n"
      default:
      - name rack:location
    - name: input-with-pipe
      type: string
      description: It includes v1 | v2 | v3
      default: v1
    - name: input_with_underscores
      type: any
      description: A variable with underscores.
      default: null
    - name: list-1
      type: list
      description: It's list number one.
      default:
      - a
      - b
      - c
    - name: list-2
      type: list
      description: It's list number two.
      default: null
    - name: list-3
      type: list
      description: null
      default: []
    - name: list_default_empty
      type: list(string)
      description: null
      default: []
    - name: long_type
      type: |-
        object({
            name = string,
            foo  = object({ foo = string, bar = string }),
            bar  = object({ foo = string, bar = string }),
            fizz = list(string),
            buzz = list(string)
          })
      description: |
        This description is itself markdown.

        It spans over multiple lines.
      default:
        bar:
          bar: bar
          foo: bar
        buzz:
        - fizz
        - buzz
        fizz: []
        foo:
          bar: foo
          foo: foo
        name: hello
    - name: map-1
      type: map
      description: It's map number one.
      default:
        a: 1
        b: 2
        c: 3
    - name: map-2
      type: map
      description: It's map number two.
      default: null
    - name: map-3
      type: map
      description: null
      default: {}
    - name: no-escape-default-value
      type: string
      description: The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'.
      default: VALUE_WITH_UNDERSCORE
    - name: number-1
      type: number
      description: It's number number one.
      default: 42
    - name: number-2
      type: number
      description: It's number number two.
      default: null
    - name: number-3
      type: number
      description: null
      default: "19"
    - name: number-4
      type: number
      description: null
      default: 15.75
    - name: number_default_zero
      type: number
      description: null
      default: 0
    - name: object_default_empty
      type: object({})
      description: null
      default: {}
    - name: string-1
      type: string
      description: It's string number one.
      default: bar
    - name: string-2
      type: string
      description: It's string number two.
      default: null
    - name: string-3
      type: string
      description: null
      default: ""
    - name: string_default_empty
      type: string
      description: null
      default: ""
    - name: string_default_null
      type: string
      description: null
      default: "null"
    - name: string_no_default
      type: string
      description: null
      default: null
    - name: unquoted
      type: any
      description: null
      default: null
    - name: with-url
      type: string
      description: The description contains url. https://www.domain.com/foo/bar_baz.html
      default: ""
    outputs:
    - name: output-0.12
      description: terraform 0.12 only
    - name: output-1
      description: It's output number one.
    - name: output-2
      description: It's output number two.
    - name: unquoted
      description: It's unquoted output.
    providers:
    - name: aws
      alias: null
      version: '>= 2.15.0'
    - name: aws
      alias: ident
      version: '>= 2.15.0'
    - name: "null"
      alias: null
      version: null
    - name: tls
      alias: null
      version: null
    requirements:
    - name: terraform
      version: '>= 0.12'
    - name: aws
      version: '>= 2.15.0'


###### Auto generated by spf13/cobra on 27-Mar-2020
