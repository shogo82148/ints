package ints

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestUint8_Text(t *testing.T) {
	var b big.Int
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int8(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint8_String(t *testing.T) {
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int8(%d).String() = %q, want %q", i, got, want)
		}
	}
}
