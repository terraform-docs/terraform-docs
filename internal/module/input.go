package module

import (
	"github.com/segmentio/terraform-docs/pkg/tfconf"
)

type inputsSortedByName []*tfconf.Input

func (a inputsSortedByName) Len() int      { return len(a) }
func (a inputsSortedByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByName) Less(i, j int) bool {
	return a[i].Name < a[j].Name
}

type inputsSortedByRequired []*tfconf.Input

func (a inputsSortedByRequired) Len() int      { return len(a) }
func (a inputsSortedByRequired) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByRequired) Less(i, j int) bool {
	if a[i].HasDefault() == a[j].HasDefault() {
		return a[i].Name < a[j].Name
	}
	return !a[i].HasDefault() && a[j].HasDefault()
}

type inputsSortedByPosition []*tfconf.Input

func (a inputsSortedByPosition) Len() int      { return len(a) }
func (a inputsSortedByPosition) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a inputsSortedByPosition) Less(i, j int) bool {
	return a[i].Position.Filename < a[j].Position.Filename || a[i].Position.Line < a[j].Position.Line
}
