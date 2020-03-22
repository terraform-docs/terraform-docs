variable "A" {
    default = "A default"
}

variable "B" {
    description = "The B variable"
}

output "A" {
    value = "${var.A}"
}

output "B" {
    description = "I am B"
    value = "${var.A}"
}

resource "null_resource" "A" {}
resource "null_resource" "B" {}
