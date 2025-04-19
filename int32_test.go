package ints

import (
	"math"
	"math/big"
	"testing"
)

func FuzzInt32_Text(f *testing.F) {
	f.Add(int32(0), 10)
	f.Add(int32(0), 62)
	f.Add(int32(math.MinInt32), 2)
	f.Add(int32(math.MaxInt32), 2)
	f.Add(int32(math.MinInt32), 10)
	f.Add(int32(math.MaxInt32), 10)
	f.Add(int32(math.MinInt32), 62)
	f.Add(int32(math.MaxInt32), 62)
	f.Fuzz(func(t *testing.T, x int32, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Int32(x)
		b := big.NewInt(int64(x))
		got := a.Text(base)
		want := b.Text(base)
		if got != want {
			t.Errorf("Int32(%d).Text(%d) = %q, want %q", x, base, got, want)
		}
	})
}

func FuzzInt32_Append(f *testing.F) {
	f.Add(int32(0), 10)
	f.Add(int32(0), 62)
	f.Add(int32(math.MinInt32), 2)
	f.Add(int32(math.MaxInt32), 2)
	f.Add(int32(math.MinInt32), 10)
	f.Add(int32(math.MaxInt32), 10)
	f.Add(int32(math.MinInt32), 62)
	f.Add(int32(math.MaxInt32), 62)
	f.Fuzz(func(t *testing.T, x int32, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Int32(x)
		b := big.NewInt(int64(x))
		got := a.Append(nil, base)
		want := b.Text(base)
		if string(got) != want {
			t.Errorf("Int32(%d).Append(buf, %d) = %q, want %q", x, base, string(got), want)
		}
	})
}

func FuzzInt32_AppendText(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, x int32) {
		a := Int32(x)
		b := big.NewInt(int64(x))
		got, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		want := b.String()
		if string(got) != want {
			t.Errorf("Int32(%d).AppendText(buf) = %q, want %q", x, string(got), want)
		}
	})
}

func FuzzInt32_String(f *testing.F) {
	f.Add(int32(0))
	f.Add(int32(math.MinInt32))
	f.Add(int32(math.MaxInt32))
	f.Fuzz(func(t *testing.T, x int32) {
		a := Int32(x)
		b := big.NewInt(int64(x))
		got := a.String()
		want := b.String()
		if string(got) != want {
			t.Errorf("Int32(%d).String() = %q, want %q", x, got, want)
		}
	})
}
