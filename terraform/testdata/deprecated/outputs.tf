output "legacy_id" {
  value       = "abc123"
  description = "Legacy output scheduled for removal."
  deprecated  = "Use output.id; will be removed in v2.0."
}

output "id" {
  value       = "abc123"
  description = "Replacement output."
}
