package module

import (
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

type inputsSortedByName []*tfconf.Input

func (in inputsSortedByName) Len() int      { return len(in) }
func (in inputsSortedByName) Swap(i, j int) { in[i], in[j] = in[j], in[i] }
func (in inputsSortedByName) Less(i, j int) bool {
	return in[i].Name < in[j].Name
}

type inputsSortedByRequired []*tfconf.Input

func (ir inputsSortedByRequired) Len() int      { return len(ir) }
func (ir inputsSortedByRequired) Swap(i, j int) { ir[i], ir[j] = ir[j], ir[i] }
func (ir inputsSortedByRequired) Less(i, j int) bool {
	if ir[i].HasDefault() == ir[j].HasDefault() {
		return ir[i].Name < ir[j].Name
	}
	return !ir[i].HasDefault() && ir[j].HasDefault()
}

type inputsSortedByPosition []*tfconf.Input

func (ip inputsSortedByPosition) Len() int      { return len(ip) }
func (ip inputsSortedByPosition) Swap(i, j int) { ip[i], ip[j] = ip[j], ip[i] }
func (ip inputsSortedByPosition) Less(i, j int) bool {
	return ip[i].Position.Filename < ip[j].Position.Filename || ip[i].Position.Line < ip[j].Position.Line
}

type inputsSortedByType []*tfconf.Input

func (it inputsSortedByType) Len() int      { return len(it) }
func (it inputsSortedByType) Swap(i, j int) { it[i], it[j] = it[j], it[i] }
func (it inputsSortedByType) Less(i, j int) bool {
	if it[i].Type == it[j].Type {
		return it[i].Name < it[j].Name
	}
	return it[i].Type < it[j].Type
}
