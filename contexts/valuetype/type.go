package valuetype

// Type is a value kindness.
type Type uint64

// String returns type name.
func (t Type) String() string {
	return rm[t]
}

var m = map[string]Type{}
var rm = map[Type]string{}

var typeSeq uint = 0

func Register(name string) (typ Type, err error) {
	if _, exists := m[name]; exists {
		return Type(0), ErrSameTypeExist
	}

	typeSeq++
	typ = Type(typeSeq)
	m[name] = typ
	rm[typ] = name
	return
}
