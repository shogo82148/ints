package ints

import (
	"math"
	"math/big"
	"testing"
)

func FuzzInt64_Text(f *testing.F) {
	f.Add(int64(0), 10)
	f.Add(int64(0), 62)
	f.Add(int64(math.MinInt64), 2)
	f.Add(int64(math.MaxInt64), 2)
	f.Add(int64(math.MinInt64), 10)
	f.Add(int64(math.MaxInt64), 10)
	f.Add(int64(math.MinInt64), 62)
	f.Add(int64(math.MaxInt64), 62)
	f.Fuzz(func(t *testing.T, x int64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Int64(x)
		b := big.NewInt(int64(x))
		got := a.Text(base)
		want := b.Text(base)
		if got != want {
			t.Errorf("Int64(%d).Text(%d) = %q, want %q", x, base, got, want)
		}
	})
}

func FuzzInt64_Append(f *testing.F) {
	f.Add(int64(0), 10)
	f.Add(int64(0), 62)
	f.Add(int64(math.MinInt64), 2)
	f.Add(int64(math.MaxInt64), 2)
	f.Add(int64(math.MinInt64), 10)
	f.Add(int64(math.MaxInt64), 10)
	f.Add(int64(math.MinInt64), 62)
	f.Add(int64(math.MaxInt64), 62)
	f.Fuzz(func(t *testing.T, x int64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Int64(x)
		b := big.NewInt(int64(x))
		got := a.Append(nil, base)
		want := b.Text(base)
		if string(got) != want {
			t.Errorf("Int64(%d).Append(buf, %d) = %q, want %q", x, base, string(got), want)
		}
	})
}

func FuzzInt64_AppendText(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, x int64) {
		a := Int64(x)
		b := big.NewInt(int64(x))
		got, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		want := b.String()
		if string(got) != want {
			t.Errorf("Int64(%d).AppendText(buf) = %q, want %q", x, string(got), want)
		}
	})
}

func FuzzInt64_String(f *testing.F) {
	f.Add(int64(0))
	f.Add(int64(math.MinInt64))
	f.Add(int64(math.MaxInt64))
	f.Fuzz(func(t *testing.T, x int64) {
		a := Int64(x)
		b := big.NewInt(int64(x))
		got := a.String()
		want := b.String()
		if string(got) != want {
			t.Errorf("Int64(%d).String() = %q, want %q", x, got, want)
		}
	})
}
