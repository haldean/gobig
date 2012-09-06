package gobig

func (this *Bignum) Cmp(other *Bignum) int8 {
  if len(this.value) < len(other.value) {
    return -1
  } else if len(this.value) > len(other.value) {
    return 1
  }

  for i := len(this.value) - 1; i >= 0; i-- {
    if this.value[i] < other.value[i] {
      return -1
    } else if this.value[i] > other.value[i] {
      return 1
    }
  }

  return 0
}

func (this *Bignum) Equals(other *Bignum) bool {
  return this.Cmp(other) == 0
}

func (this *Bignum) NotEquals(other *Bignum) bool {
  return this.Cmp(other) != 0
}

func (this *Bignum) Less(other *Bignum) bool {
  return this.Cmp(other) < 0
}

func (this *Bignum) LessOrEqual(other *Bignum) bool {
  return this.Cmp(other) <= 0
}

func (this *Bignum) Greater(other *Bignum) bool {
  return this.Cmp(other) > 0
}

func (this *Bignum) GreaterOrEqual(other *Bignum) bool {
  return this.Cmp(other) >= 0
}
