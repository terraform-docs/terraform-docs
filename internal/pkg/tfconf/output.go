package tfconf

// Output represents a Terraform output.
type Output struct {
	Name        string   `json:"name" yaml:"name"`
	Description String   `json:"description" yaml:"description"`
	Value       string   `json:"value,omitempty" yaml:"value,omitempty"`
	Position    Position `json:"-" yaml:"-"`
}

type outputsSortedByName []*Output

func (a outputsSortedByName) Len() int {
	return len(a)
}

func (a outputsSortedByName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a outputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type outputsSortedByPosition []*Output

func (a outputsSortedByPosition) Len() int {
	return len(a)
}

func (a outputsSortedByPosition) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a outputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}

// TerraformOutput is used for unmarshalling `terraform outputs --json` into
type TerraformOutput struct {
	Sensitive bool   `json:"sensitive"`
	Type      string `json:"type"`
	Value     string `json:"value"`
}
