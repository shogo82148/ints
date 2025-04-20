package ints

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func FuzzInt64_Add(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(1), int64(0))
	f.Add(int64(math.MaxInt64), int64(math.MaxInt64))
	f.Add(int64(math.MinInt64), int64(math.MinInt64))
	f.Fuzz(func(t *testing.T, x, y int64) {
		a := Int64(x)
		b := Int64(y)
		got := a.Add(b)
		want := Int64(int64(x + y))
		if got != want {
			t.Errorf("Int64(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt64_Sub(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(1), int64(0))
	f.Add(int64(math.MaxInt64), int64(-1))
	f.Add(int64(math.MinInt64), int64(1))
	f.Fuzz(func(t *testing.T, x, y int64) {
		a := Int64(x)
		b := Int64(y)
		got := a.Sub(b)
		want := Int64(int64(x - y))
		if got != want {
			t.Errorf("Int64(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt64_Mul(f *testing.F) {
	f.Add(int64(0), int64(0))
	f.Add(int64(1), int64(0))
	f.Add(int64(math.MaxInt64), int64(math.MaxInt64))
	f.Add(int64(math.MinInt64), int64(math.MinInt64))
	f.Fuzz(func(t *testing.T, x, y int64) {
		a := Int64(x)
		b := Int64(y)
		got := a.Mul(b)
		want := Int64(int64(x * y))
		if got != want {
			t.Errorf("Int64(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestInt64_Sign(t *testing.T) {
	testCases := []struct {
		x    Int64
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{math.MaxInt64, 1},
		{math.MinInt64, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int64(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt64_Neg(t *testing.T) {
	testCases := []struct {
		x    Int64
		want Int64
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{math.MaxInt64, -math.MaxInt64},
		{math.MinInt64, math.MinInt64},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int64(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt64_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int64
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxInt64, math.MaxInt64, 0},
		{math.MinInt64, math.MinInt64, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int64(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

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

func TestInt64_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int64
		want   string
	}{
		// decimal
		{
			"%d",
			0,
			"0",
		},
		{
			"%d",
			-9223372036854775808,
			"-9223372036854775808",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
