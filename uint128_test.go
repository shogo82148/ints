package ints

import (
	"math"
	"math/big"
	"runtime"
	"testing"
)

func uint128ToBigInt(a Uint128) *big.Int {
	var b, c big.Int
	b.SetUint64(a[0])
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[1]))
	return &b
}

func FuzzUint128_Add(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(math.MaxUint64),
		uint64(0), uint64(1),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Add(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint128_Sub(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0),
		uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Sub(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint128_Mul(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0),
		uint64(1), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Mul(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestUint128_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want int
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{math.MaxUint64, math.MaxUint64}, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint128(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{0, 1}, Uint128{math.MaxUint64, math.MaxUint64}},
		{Uint128{math.MaxUint64, math.MaxUint64}, Uint128{0, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint128(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func FuzzUint128_Text(f *testing.F) {
	f.Add(uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, u0, u1 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint128{u0, u1}
		got := a.Text(base)

		b := uint128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_Append(f *testing.F) {
	f.Add(uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, u0, u1 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint128{u0, u1}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Uint128{u0, u1}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_String(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Uint128{u0, u1}
		got := a.String()

		b := uint128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func BenchmarkUint128_Text2(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint128_Text10(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkUint128_Text62(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkUint128_String(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}
