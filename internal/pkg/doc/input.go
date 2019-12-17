package doc

// Input represents a Terraform input.
type Input struct {
	Name        string `json:"name"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Default     string `json:"default,omitempty"`
}

// HasDefault indicates if a Terraform variable has a default value set.
func (i *Input) HasDefault() bool {
	return len(i.Default) > 0
}

type variablesSortedByName []Input

func (a variablesSortedByName) Len() int {
	return len(a)
}

func (a variablesSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a variablesSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type variablesSortedByRequired []Input

func (a variablesSortedByRequired) Len() int {
	return len(a)
}

func (a variablesSortedByRequired) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a variablesSortedByRequired) Less(i, j int) bool {
	switch {
	// i required, j not: i gets priority
	case !a[i].HasDefault() && a[j].HasDefault():
		return true
	// j required, i not: i does not get priority
	case a[i].HasDefault() && !a[j].HasDefault():
		return false
	// Otherwise, sort by name
	default:
		return a[i].Name < a[j].Name
	}
}
