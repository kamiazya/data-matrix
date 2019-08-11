package schema

import "github.com/eknkc/amber/parser"

type ColumnSchema struct {
	Name string
	parser.Parser
}
