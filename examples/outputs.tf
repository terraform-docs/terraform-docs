output unquoted {
  description = "It's unquoted output."
  value       = ""
}

output "output-2" {
  description = "It's output number two."
  value       = "2"
}

// It's output number one.
output "output-1" {
  value = "1"
}

output "output-0.12" {
  value       = join(",", var.list-3)
  description = "terraform 0.12 only"
}

// terraform-docs-ignore
output "ignored" {
  value = "ignored"
}
