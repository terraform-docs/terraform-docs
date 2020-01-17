variable unquoted {}

variable "string-3" {
  default = ""
}

variable "string-2" {
  description = "It's string number two."
  type        = "string"
}

// It's string number one.
variable "string-1" {
  default = "bar"
}

variable "map-3" {
  default = {}
}

variable "map-2" {
  description = "It's map number two."
  type        = "map"
}

// It's map number one.
variable "map-1" {
  default = {
    a = 1
    b = 2
    c = 3
  }

  type = "map"
}

variable "list-3" {
  default = []
}

variable "list-2" {
  description = "It's list number two."
  type        = "list"
}

// It's list number one.
variable "list-1" {
  default = ["a", "b", "c"]
  type    = "list"
}

// A variable with underscores.
variable "input_with_underscores" {}

// A variable with pipe in the description
variable "input-with-pipe" {
  description = "It includes v1 | v2 | v3"
  default     = "v1"
}

variable "input-with-code-block" {
  description = <<EOD
This is a complicated one. We need a newline.  
And an example in a code block
```
default     = [
  "machine rack01:neptune"
]
```
EOD
  default     = [
    "name rack:location"
  ]
}

variable "long_type" {
  type = object({
    name = string,
    foo  = object({ foo = string, bar = string }),
    bar  = object({ foo = string, bar = string }),
    fizz = list(string),
    buzz = list(string)
  })
  default = {
    name = "hello"
    foo = {
      foo = "foo"
      bar = "foo"
    }
    bar = {
      foo = "bar"
      bar = "bar"
    },
    fizz = []
    buzz = ["fizz", "buzz"]
  }
  description = <<EOF
This description is itself markdown.

It spans over multiple lines.
EOF
}

variable "no-escape-default-value" {
  description = "The description contains `something_with_underscore`. Defaults to 'VALUE_WITH_UNDERSCORE'."
  default     = "VALUE_WITH_UNDERSCORE"
}

variable "with-url" {
  description = "The description contains url. https://www.domain.com/foo/bar_baz.html"
  default     = ""
}
