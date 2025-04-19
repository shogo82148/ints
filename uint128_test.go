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
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint128_String(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}
