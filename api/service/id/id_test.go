package id

import (
	"testing"
)

func TestSummonID(t *testing.T) {
	s := NewSummonID()

	var x, y ID
	for i := 0; i < 1e4; i++ {
		y = s.Generate()
		str := y.Base2()

		t.Logf("\033[31m%v\033[0m-\033[34m%v\033[33m%v\033[0m", y.String(), str[:42], str[42:])

		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}
