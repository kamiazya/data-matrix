package intager

import "github.com/kamiazya/data-matrix/contexts/valuetype"

var Type valuetype.Type

func init() {
	Type, _ = valuetype.Register("intager")
}
