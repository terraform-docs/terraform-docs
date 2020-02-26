package tfconf

// Position represents position of Terraform input or output in a file.
type Position struct {
	Filename string `json:"-" xml:"-" yaml:"-"`
	Line     int    `json:"-" xml:"-" yaml:"-"`
}
