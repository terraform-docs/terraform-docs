output "multi-line-lf" {
  value = "foo"
  description = <<-EOT
  The quick brown fox jumps
  over the lazy dog
  EOT
}
