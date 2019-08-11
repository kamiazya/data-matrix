package parser

import "sort"

type OrderedComponent struct {
	Order int
	Component
}

var _ sort.Interface = (*order)(nil)

type order []*OrderedComponent

func (a order) Len() int           { return len(a) }
func (a order) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a order) Less(i, j int) bool { return a[i].Order < a[j].Order }

func By(order int, c Component) *OrderedComponent {
	return &OrderedComponent{order, c}
}
