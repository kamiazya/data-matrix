package parser

import "github.com/kamiazya/data-matrix/contexts/datastructure"

type parser struct {
	cs []Component
}

func (p *parser) Parse(bs []byte) (v datastructure.Value) {
	for _, c := range []Component(p.cs) {
		if v = c.Parse(bs); v != nil {
			return v
		}
	}
	return
}
