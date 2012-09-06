package gobig

var Base uint64 = 1 << 32;

type Bignum struct {
  // Digits in the defined base, in little-endian order
  value []uint32;
}

// Resize the value slice to the given length. Clients of gobig should never
// call this function.
func (this *Bignum) resize(size uint64) {
  new_slice := make([]uint32, size)
  copy(new_slice, this.value)
  this.value = new_slice
}

// Resize the value slice to exactly fit the current contents. This essentially
// just trims zeroes from the most-significant end. Clients of gobig should
// never need to call this function; all bignums returned by API calls should
// already be fitted.
func (this *Bignum) fit() {
  var first_zero uint64 = 0
  for i, v := range this.value {
    if v == 0 && first_zero == 0 {
      first_zero = uint64(i)
    } else if v != 0 {
      first_zero = 0
    }
  }

  this.resize(first_zero)
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
