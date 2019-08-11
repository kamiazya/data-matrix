package datastructure

type RowItterator interface {
	Next() bool
	Columns() ([]string, error)

	Scan(dest ...Value) error
}
