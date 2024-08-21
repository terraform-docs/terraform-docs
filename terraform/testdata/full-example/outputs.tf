output C {
  description = "It's unquoted output."
  value       = "c"
}

output "A" {
  description = "A description"
  value       = "a"
}

// B description
output "B" {
  value = "b"
}

// D null result
output "D" {
  value = null
}

# terraform-docs-ignore
output "ignored" {
  value = "e"
}
