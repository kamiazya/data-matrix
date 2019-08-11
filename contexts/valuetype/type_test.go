package valuetype

import "testing"

func TestRegister(t *testing.T) {
	type testcase struct {
		name string

		exceptErr  error
		exceptType Type
	}

	tc := []testcase{
		{
			// add hoge type
			name:       "hoge",
			exceptErr:  nil,
			exceptType: Type(1), // it may be 1st type
		},
		{
			// add fuga type
			name:       "fuga",
			exceptErr:  nil,
			exceptType: Type(2), // is may be 2nd type
		},
		{
			// add type huga
			name:      "fuga",
			exceptErr: ErrSameTypeExist, // it alleady exist type.
		},
	}
	for _, c := range tc {
		typ, err := Register(c.name)

		if err != c.exceptErr {
			t.Error("register catch unexcepted error.", err.Error())
		} else if typ.String() != c.exceptType.String() {
			t.Error("register catch unexcepted type.", typ.String())
		}
	}
}
