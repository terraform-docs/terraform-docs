variable "log_categories" {
  default = ["one", "two", "three"]
}
variable "enabled" {
  default = true
}
variable "retention_days" {
  default = 7
}

locals {
  logs = {
    for category in var.log_categories :
    category => {
      enabled        = var.enabled
      retention_days = var.retention_days
    }
  }
}
