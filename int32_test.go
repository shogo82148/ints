package ints

import (
	"fmt"
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

func TestInt32_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Int32
		z, r Int32
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
			t.Errorf("Int32(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Int32(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Int32(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Int32(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt32_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Int32
		q, r Int32
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
			t.Errorf("Int32(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Int32(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Int32(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Int32(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt32_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int32
		i    uint
		want Int32
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int32(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt32_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int32
		i    uint
		want Int32
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{math.MaxInt32, 1, math.MaxInt32 >> 1},

		// Sign extension
		{-128, 0, -128},
		{-128, 1, -64},
		{-128, 2, -32},
		{-128, 3, -16},
		{-128, 4, -8},
		{-128, 5, -4},
		{-128, 6, -2},
		{-128, 7, -1},
		{-128, 8, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int32(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
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

func TestInt32_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int32
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxInt32, math.MaxInt32, 0},
		{math.MinInt32, math.MinInt32, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int32(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
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

func TestInt32_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int32
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
			-2147483648,
			"-2147483648",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
