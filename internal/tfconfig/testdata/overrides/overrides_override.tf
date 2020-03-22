variable "A" {
  description = "The A variable OVERRIDDEN"
}

variable "C" {
  description = "An entirely new variable C"
}

output "A" {
  description = "I am an overridden output!"
  value       = "${var.A}"
}

module "foo" {
  version = "1.0.2_override"

  unused = 2
}
