package doc

// Output represents a Terraform output.
type Output struct {
	Name        string
	Description string
}

// HasDescription indicates if a Terraform output has a description.
func (o *Output) HasDescription() bool {
	return o.Description != ""
}
