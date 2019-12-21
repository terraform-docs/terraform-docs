package tfconf

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

type inputsSortedByName []*Input

func (a inputsSortedByName) Len() int {
	return len(a)
}

func (a inputsSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a inputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type inputsSortedByRequired []*Input

func (a inputsSortedByRequired) Len() int {
	return len(a)
}

func (a inputsSortedByRequired) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a inputsSortedByRequired) Less(i, j int) bool {
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
