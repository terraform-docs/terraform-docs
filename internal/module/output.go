package module

import (
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

// TerraformOutput is used for unmarshalling `terraform outputs --json` into
type TerraformOutput struct {
	Sensitive bool        `json:"sensitive"`
	Type      interface{} `json:"type"`
	Value     interface{} `json:"value"`
}

type outputsSortedByName []*tfconf.Output

func (on outputsSortedByName) Len() int      { return len(on) }
func (on outputsSortedByName) Swap(i, j int) { on[i], on[j] = on[j], on[i] }
func (on outputsSortedByName) Less(i, j int) bool {
	return on[i].Name < on[j].Name
}

type outputsSortedByPosition []*tfconf.Output

func (op outputsSortedByPosition) Len() int      { return len(op) }
func (op outputsSortedByPosition) Swap(i, j int) { op[i], op[j] = op[j], op[i] }
func (op outputsSortedByPosition) Less(i, j int) bool {
	return op[i].Position.Filename < op[j].Position.Filename || op[i].Position.Line < op[j].Position.Line
}
