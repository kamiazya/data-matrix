package parser

import (
	"sort"

	"github.com/kamiazya/data-matrix/contexts/datastructure"
)

func New(cs ...*OrderedComponent) Paeser {
	sort.Sort(order(cs))
	p := &parser{
		cs: make([]Component, len(cs)),
	}
	for i, c := range cs {
		p.cs[i] = c
	}
	return p
}

type Paeser interface {
	Parse(bs []byte) datastructure.Value
}
