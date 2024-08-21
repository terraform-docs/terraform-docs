// D description
variable "D" {
  default = "d"
}

variable "B" {
  default = "b"
}

variable "E" {
  default = ""
}

# A Description
# in multiple lines
variable A {}

variable "C" {
  description = "C description"
  default = "c"
}

variable "F" {
  description = "F description"
}

variable "G" {
  description = "G description"
  default     = null
}

# terraform-docs-ignore
variable "ignored" {
  description = "H description"
  default = null
}
