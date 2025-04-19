package ints

import (
	"math"
	"math/big"
	"testing"
)

func FuzzInt32_Add(f *testing.F) {
	f.Add(int32(0), int32(0))
	f.Add(int32(1), int32(0))
	f.Add(int32(math.MaxInt32), int32(math.MaxInt32))
	f.Add(int32(math.MinInt32), int32(math.MinInt32))
	f.Fuzz(func(t *testing.T, x, y int32) {
		a := Int32(x)
		b := Int32(y)
		got := a.Add(b)
		want := Int32(int32(x + y))
		if got != want {
			t.Errorf("Int32(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt32_Sub(f *testing.F) {
	f.Add(int32(0), int32(0))
	f.Add(int32(1), int32(0))
	f.Add(int32(math.MaxInt32), int32(-1))
	f.Add(int32(math.MinInt32), int32(1))
	f.Fuzz(func(t *testing.T, x, y int32) {
		a := Int32(x)
		b := Int32(y)
		got := a.Sub(b)
		want := Int32(int32(x - y))
		if got != want {
			t.Errorf("Int32(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt32_Mul(f *testing.F) {
	f.Add(int32(0), int32(0))
	f.Add(int32(1), int32(0))
	f.Add(int32(math.MaxInt32), int32(math.MaxInt32))
	f.Add(int32(math.MinInt32), int32(math.MinInt32))
	f.Fuzz(func(t *testing.T, x, y int32) {
		a := Int32(x)
		b := Int32(y)
		got := a.Mul(b)
		want := Int32(int32(x * y))
		if got != want {
			t.Errorf("Int32(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestInt32_Sign(t *testing.T) {
	testCases := []struct {
		x    Int32
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{math.MaxInt32, 1},
		{math.MinInt32, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int32(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt32_Neg(t *testing.T) {
	testCases := []struct {
		x    Int32
		want Int32
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{math.MaxInt32, -math.MaxInt32},
		{math.MinInt32, math.MinInt32},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int32(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

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
