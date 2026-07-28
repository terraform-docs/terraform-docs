variable unquoted {}

variable "bool-3" {
  default = true
}

variable "bool-2" {
  description = "It's bool number two."
  default     = false
}

// It's bool number one.
variable "bool-1" {
  default = true
}

# terraform-docs-ignore
variable "ignored" {
  default = ""
}

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

variable "string-special-chars" {
  default = "\\.<>[]{}_-"
}

variable "number-3" {
  type    = number
  default = "19"
}

variable "number-4" {
  type    = number
  default = 15.75
}

variable "number-2" {
  description = "It's number number two."
  type        = "number"
}

// It's number number one.
variable "number-1" {
  default = 42
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
  default = [
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

variable "string_default_empty" {
  type    = string
  default = ""
}

variable "string_default_null" {
  type    = string
  default = null
}

variable "string_no_default" {
  type = string
}

variable "number_default_zero" {
  type    = number
  default = 0
}

variable "bool_default_false" {
  type    = bool
  default = false
}

variable "list_default_empty" {
  type    = list(string)
  default = []
}

variable "object_default_empty" {
  type    = object({})
  default = {}
}

variable "variable_with_no_validation" {
    description = "This variable has no validation"
    type        = string
    default     = ""
}

variable "variable_with_one_validation" {
    description = "This variable has one validation"
    type        = string
    default     = ""

    validation {
        condition     = (length(var.variable_with_one_validation) == 0 || length(var.variable_with_one_validation) == 10)
        error_message = "var.variable_with_one_validation must be empty or 10 characters long."
    }
}

variable "variable_with_two_validations" {
    description = "This variable has two validations"
    type        = string

    validation {
        condition     = (length(var.variable_with_two_validations) == 10)
        error_message = "var.variable_with_two_validations must be 10 characters long."
    }

    validation {
        condition     = (startswith(var.variable_with_two_validations, "magic"))
        error_message = "var.variable_with_two_validations must start with 'magic'."
    }
}
