package datastructure

import "github.com/kamiazya/data-matrix/contexts/valuetype"

type Value interface {
	Type() valuetype.Type
	String() string
	Bytes() []byte
}
