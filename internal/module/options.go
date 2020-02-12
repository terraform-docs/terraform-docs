package module

// SortBy contains different sort criteria corresponding
// to available flags (e.g. name, required, etc)
type SortBy struct {
	Name     bool
	Required bool
}

// Options contains required options to load a Module from path
type Options struct {
	Path   string
	SortBy *SortBy
}
