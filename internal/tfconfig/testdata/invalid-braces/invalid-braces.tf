# The opening brace is required to be on the same line as the block header
# under HCL 2, and so this should produce an error, causing us to use HCL 1's
# parser instead.

variable foo
{
  default = "123"
}
