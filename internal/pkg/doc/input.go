package doc

// Input represents a Terraform input variable.
type Input struct {
	Name        string
	Description string
	Default     *Value
	Type        string
}

// HasDescription indicates if a Terraform input has a description.
func (i *Input) HasDescription() bool {
	return i.Description != ""
}

// IsOptional indicates if a Terraform input is optional.
func (i *Input) IsOptional() bool {
	return !i.IsRequired()
}

// IsRequired indicates if a Terraform input is required.
func (i *Input) IsRequired() bool {
	return i.Default == nil
}
