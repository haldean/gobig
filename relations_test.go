package gobig

import (
	"testing"
)

func TestRelations(t *testing.T) {
  a := Bignum{[]uint32{1, 3, 4, 2}}
  b := Bignum{[]uint32{1, 3, 4, 2}}
  c := Bignum{[]uint32{3, 4, 2}}
  d := Bignum{[]uint32{10, 3, 4, 2}}
  e := Bignum{[]uint32{11, 3, 4, 2}}
  f := Bignum{[]uint32{1, 11, 3, 4, 2}}

  if (!a.Equals(&b)) {
    t.Error("a should equal b")
  }

  if (!b.Equals(&b)) {
    t.Error("Equals should be reflexive")
  }

  if (!a.Less(&f)) {
    t.Error("a should be less than b")
  }

  if (!f.Greater(&a)) {
    t.Error("Less and Greater should be inverses")
  }

  if (!a.LessOrEqual(&f)) {
    t.Error("a should be less than or equal to b")
  }

  if (!f.GreaterOrEqual(&a)) {
    t.Error("LessOrEqual and GreaterOrEqual should be inverses")
  }

  if (!c.Less(&a)) {
    t.Error("c should be less than a")
  }

  if (!e.Greater(&d)) {
    t.Error("e should be greater than d")
  }

  if (!a.NotEquals(&f)) {
    t.Error("a should not equal f")
  }
}
