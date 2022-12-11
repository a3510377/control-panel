package service

import (
	"testing"
)

func TestSummonID(t *testing.T) {
	s := NewSummonID()

	var x, y ID
	for i := 0; i < 1e4; i++ {
		y = s.Generate()
		t.Log(y.Base2())
		if x == y {
			t.Errorf("x(%d) & y(%d) are the same", x, y)
		}
		x = y
	}
}
