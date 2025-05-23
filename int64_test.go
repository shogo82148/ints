package ints

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestInt64_Add(t *testing.T) {
	testCases := []struct {
		x, y int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		a := Int64(tc.x)
		b := Int64(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Int64(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt64_Sub(t *testing.T) {
	testCases := []struct {
		x, y int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, tc := range testCases {
		a := Int64(tc.x)
		b := Int64(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Int64(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt64_Mul(t *testing.T) {
	testCases := []struct {
		x, y int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{-1, -1, 1},
	}

	for _, tc := range testCases {
		a := Int64(tc.x)
		b := Int64(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Int64(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt64_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Int64
		z, r Int64
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
		{127, -10, -12, 7},
		{-127, 10, -13, 3},
		{-127, -10, 13, 3},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Int64(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Int64(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Int64(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Int64(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt64_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Int64
		q, r Int64
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
		{127, -10, -12, 7},
		{-127, 10, -12, -7},
		{-127, -10, 12, -7},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Int64(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Int64(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Int64(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Int64(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt64_And(t *testing.T) {
	testCases := []struct {
		x    Int64
		y    Int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Int64(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt64_AndNot(t *testing.T) {
	testCases := []struct {
		x    Int64
		y    Int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{math.MaxInt64, math.MaxInt64, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Int64(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt64_Or(t *testing.T) {
	testCases := []struct {
		x    Int64
		y    Int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
		{math.MaxInt64, math.MaxInt64, math.MaxInt64},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Int64(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt64_Xor(t *testing.T) {
	testCases := []struct {
		x    Int64
		y    Int64
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{math.MaxInt64, math.MaxInt64, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Int64(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt64_Not(t *testing.T) {
	testCases := []struct {
		x    Int64
		want Int64
	}{
		{0, -1},
		{1, -2},
		{-1, 0},
		{math.MaxInt64, -math.MaxInt64 - 1},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Int64(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt64_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int64
		i    uint
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
		{1, 2, 4},
		{1, 3, 8},
		{1, 4, 16},
		{1, 5, 32},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int64(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt64_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int64
		i    uint
		want Int64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{math.MaxInt64, 1, math.MaxInt64 >> 1},

		// Sign extension
		{-128, 0, -128},
		{-128, 1, -64},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int64(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
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
