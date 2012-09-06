package gobig

import (
	"testing"
)

func TestDecompose(t *testing.T) {
  c := Decompose(1 << 32 + 1);
  if (c.value[0] != 1 || c.value[1] != 1) {
    t.Error("Decomposing 1 << 32 + 1 should produce [1 1]")
  }
}

func TestCompose(t *testing.T) {
  c := new(Bignum)
  c.value = []uint32{1, 1}
  if (Compose(c) != 1 << 32 + 1) {
    t.Error("Composing [1 1] should produce 1 << 32 + 1")
  }
}

func TestDecomposeCompose(t *testing.T) {
  var x uint64 = 1 << 63 | 1 << 60 | 1 << 15;
  if (Compose(Decompose(x)) != x) {
    t.Error("Compose should be decompose's inverse.")
  }
}

func TestResize(t *testing.T) {
  c := Decompose(1 << 63 | 1)
  if len(c.value) != 2 {
    t.Error("Length of c should be 2")
  }
}
