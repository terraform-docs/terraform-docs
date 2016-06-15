/**
 * Module usage:
 *
 *      module "foo" {
 *        source = "github.com/foo/baz"
 *        subnet_ids = "${join(",", subnet.*.id)}"
 *      }
 *
 */

variable "subnet_ids" {
  description = "a comma-separated list of subnet IDs"
}

// The VPC ID.
output "vpc_id" {
  value = "vpc-5c1f55fd"
}
