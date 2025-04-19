package ints

import (
	"math"
	"math/big"
	"testing"
)

func uint1024ToBigInt(a Uint1024) *big.Int {
	b := new(big.Int)
	b.SetUint64(a[0])
	for i := 1; i < len(a); i++ {
		b.Lsh(b, 64)
		b.Add(b, new(big.Int).SetUint64(a[i]))
	}
	return b
}

func FuzzUint1024_Text(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		got := a.Text(base)

		b := uint1024ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint1024(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint1024_Append(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint1024ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint1024(%s).Append(buf, %d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint1024_AppendText(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint1024ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint1024(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzUint1024_String(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		got := a.String()

		b := uint1024ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint1024(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}
