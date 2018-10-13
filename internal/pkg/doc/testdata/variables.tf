variable unquoted {}

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
