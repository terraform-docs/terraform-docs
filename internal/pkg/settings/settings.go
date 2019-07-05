package settings

// Settings represents all settings
type Settings struct {
	// ShowRequired show "Required" column when generating Markdown
	ShowRequired bool

	// SortByName sorted rendering of inputs and outputs (default: true)
	SortByName bool

	// SortInputsByRequired sort inputs by name and prints required inputs first (default: false)
	SortInputsByRequired bool

	// AggregateTypeDefaults print default values of aggregate types (default: false)
	AggregateTypeDefaults bool
}
