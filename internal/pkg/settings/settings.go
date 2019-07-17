package settings

// Settings represents all settings
type Settings struct {
	// ShowRequired show "Required" column when generating Markdown (default: true)
	// scope: Markdown
	ShowRequired bool

	// EscapeMarkdown escapes special Markdown characters (such as | _ * and etc) (default: true)
	// scope: Markdown
	EscapeMarkdown bool

	// SortByName sorted rendering of inputs and outputs (default: true)
	// scope: Global
	SortByName bool

	// SortInputsByRequired sort inputs by name and prints required inputs first (default: false)
	// scope: Global
	SortInputsByRequired bool

	// AggregateTypeDefaults print default values of aggregate types (default: false)
	// scope: Global
	AggregateTypeDefaults bool
}
