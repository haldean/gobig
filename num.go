package gobig

var Base uint64 = 1 << 32;

type Bignum struct {
  value []uint32;
}

func Decompose(X uint64) *Bignum {
  c := make([]uint32, 0);

  // Read into slice. This puts elements in little-endian order
  for X > 0 {
    c = append(c, uint32(X % Base));
    X /= Base
  }

  return &Bignum{value: c}
}

func Compose(c *Bignum) uint64 {
  if len(c.value) == 0 {
    return 0
  }
  var X uint64 = uint64(c.value[0]);
  if len(c.value) > 1 {
    X += Base * uint64(c.value[1]);
  }
  return X;
}
