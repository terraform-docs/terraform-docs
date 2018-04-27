/**
 * Module usage:
 *
 *      module "foo" {
 *        source = "github.com/foo/baz"
 *        subnet_ids = "${join(",", subnet.*.id)}"
 *      }
 *
 */

terraform {
  required_version = ">= 0.10.0"
}

provider "aws" {
  region     = "us-west-2"
  access_key = "anaccesskey"
  secret_key = "asecretkey"
}

resource "aws_instance" "web" {
  ami           = "${var.amis[us-east-1]}"
  instance_type = "t2.micro"
  vpc_id        = "${var.vpc_id}"
  subnet_id     = ["${var.subnet_ids}"]
  vpc_security_group_ids = ["${var.security_group_ids}"]
}

variable "subnet_ids" {
  description = "a comma-separated list of subnet IDs"
}

variable "security_group_ids" {
  default = "sg-a, sg-b"
}

variable "amis" {
  default = {
    "us-east-1" = "ami-8f7687e2"
    "us-west-1" = "ami-bb473cdb"
    "us-west-2" = "ami-84b44de4"
    "eu-west-1" = "ami-4e6ffe3d"
    "eu-central-1" = "ami-b0cc23df"
    "ap-northeast-1" = "ami-095dbf68"
    "ap-southeast-1" = "ami-cf03d2ac"
    "ap-southeast-2" = "ami-697a540a"
  }
}

// The VPC ID.
output "vpc_id" {
  value = "vpc-5c1f55fd"
}
