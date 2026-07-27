variable "old_name" {
  type        = string
  default     = "x"
  description = "Legacy input scheduled for removal."
  deprecated  = "Will be removed in v2.0; use var.new_name."
}

variable "new_name" {
  type        = string
  default     = "x"
  description = "Replacement input."
}
