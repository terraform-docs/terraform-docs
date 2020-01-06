package print

// Settings represents all settings
type Settings struct {
	// AggregateTypeDefaults print default values of aggregate types (default: false)
	// scope: Global
	AggregateTypeDefaults bool

	// EscapeCharacters escapes special characters (such as | _ * in Markdown and > < in JSON) (default: true)
	// scope: Markdown
	EscapeCharacters bool

	// MarkdownIndent control the indentation of Markdown headers [available: 1, 2, 3, 4, 5] (default: 2)
	// scope: Markdown
	MarkdownIndent int

	// ShowColor print "colorized" version of result in the terminal (default: true)
	// scope: Pretty
	ShowColor bool

	// ShowInputs show "Inputs" information (default: true)
	// scope: Global
	ShowInputs bool

	// ShowOutputs show "Outputs" information (default: true)
	// scope: Global
	ShowOutputs bool

	// ShowProviders show "Providers" information (default: true)
	// scope: Global
	ShowProviders bool

	// ShowRequired show "Required" column when generating Markdown (default: true)
	// scope: Markdown
	ShowRequired bool

	// SortByName sorted rendering of inputs and outputs (default: true)
	// scope: Global
	SortByName bool

	// SortInputsByRequired sort inputs by name and prints required inputs first (default: false)
	// scope: Global
	SortInputsByRequired bool
}

//NewSettings returns new instance of Settings
func NewSettings() *Settings {
	return &Settings{
		AggregateTypeDefaults: false,
		EscapeCharacters:      true,
		MarkdownIndent:        2,
		ShowColor:             true,
		ShowInputs:            true,
		ShowOutputs:           true,
		ShowProviders:         true,
		ShowRequired:          true,
		SortByName:            true,
		SortInputsByRequired:  false,
	}
}
