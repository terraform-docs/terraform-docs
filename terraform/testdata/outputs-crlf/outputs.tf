output "multi-line-crlf" {
  value = "foo"
  description = <<-EOT
  The quick brown fox jumps
  over the lazy dog
  EOT
}
