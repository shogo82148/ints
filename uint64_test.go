package ints

import (
	"math"
	"math/big"
	"testing"
)

func FuzzUint64_Text(f *testing.F) {
	f.Add(uint64(0), 10)
	f.Add(uint64(0), 62)
	f.Add(uint64(math.MaxUint64), 2)
	f.Add(uint64(math.MaxUint64), 10)
	f.Add(uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, x uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.Text(base)
		want := b.Text(base)
		if got != want {
			t.Errorf("Uint64(%d).Text(%d) = %q, want %q", x, base, got, want)
		}
	})
}

func FuzzUint64_Append(f *testing.F) {
	f.Add(uint64(0), 10)
	f.Add(uint64(0), 62)
	f.Add(uint64(math.MaxUint64), 2)
	f.Add(uint64(math.MaxUint64), 10)
	f.Add(uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, x uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.Append(nil, base)
		want := b.Text(base)
		if string(got) != want {
			t.Errorf("Uint64(%d).Append(buf, %d) = %q, want %q", x, base, string(got), want)
		}
	})
}

func FuzzUint64_AppendText(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, x uint64) {
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint64(%d).AppendText(buf) = %q, want %q", x, string(got), want)
		}
	})
}

func FuzzUint64_String(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, x uint64) {
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.String()
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint64(%d).String() = %q, want %q", x, got, want)
		}
	})
}
