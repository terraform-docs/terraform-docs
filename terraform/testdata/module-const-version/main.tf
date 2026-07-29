variable "version" {
  type    = string
  default = "6.6.1"
  const   = true
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = var.version
}
