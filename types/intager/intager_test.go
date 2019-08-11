package intager

import (
	"bytes"
	"testing"
)

func TestInteger(t *testing.T) {

	i := &Intager{
		bs: []byte("hoge"),
		i:  10,
	}

	t.Run("type", func(t *testing.T) {
		if i.Type() != (*Intager)(nil).Type() {
			t.Error("unecepted type")
		}
	})

	t.Run("int", func(t *testing.T) {
		if i.Int() != 10 {
			t.Error("unexcepted int. it must return 10.")
		}
	})

	t.Run("string", func(t *testing.T) {
		if i.String() != "10" {
			t.Error("unexcepted string. it must return \"10\".", i.String())
		}
	})

	t.Run("bytes", func(t *testing.T) {

		if !bytes.Equal(i.Bytes(), []byte("hoge")) {
			t.Error("unexcepted bytes. it must return \"hoge\".", string(i.Bytes()))
		}
	})
}
