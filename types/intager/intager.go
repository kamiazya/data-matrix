package intager

import (
	"strconv"

	"github.com/kamiazya/data-matrix/contexts/datastructure"
	"github.com/kamiazya/data-matrix/contexts/valuetype"
)

var _ datastructure.Value = (*Intager)(nil)

type Intager struct {
	bs []byte
	i  int
}

func (*Intager) Type() valuetype.Type {
	return Type
}

func (i *Intager) Bytes() []byte {
	return i.bs
}

func (i *Intager) String() string {
	return strconv.Itoa(i.i)
}

func (i *Intager) Int() int {
	return i.i
}
