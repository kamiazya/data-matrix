package datastructure

type data interface {
	Get(i uint64) (cell Cell, err error)
	Size() uint64
}
