package ints

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"
)

func FuzzInt8_Add(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(1), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Fuzz(func(t *testing.T, x, y int8) {
		a := Int8(x)
		b := Int8(y)
		got := a.Add(b)
		want := Int8(int8(x + y))
		if got != want {
			t.Errorf("Int8(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt8_Sub(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(1), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Fuzz(func(t *testing.T, x, y int8) {
		a := Int8(x)
		b := Int8(y)
		got := a.Sub(b)
		want := Int8(int8(x - y))
		if got != want {
			t.Errorf("Int8(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt8_Mul(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(1), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Fuzz(func(t *testing.T, x, y int8) {
		a := Int8(x)
		b := Int8(y)
		got := a.Mul(b)
		want := Int8(int8(x * y))
		if got != want {
			t.Errorf("Int8(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestInt8_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Int8
		z, r Int8
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
		{127, -10, -12, 7},
		{-127, 10, -13, 3},
		{-127, -10, 13, 3},

		// integer overflow
		{-128, -1, -128, 0},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Int8(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Int8(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Int8(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Int8(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt8_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Int8
		q, r Int8
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
		{127, -10, -12, 7},
		{-127, 10, -12, -7},
		{-127, -10, 12, -7},

		// integer overflow
		{-128, -1, -128, 0},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Int8(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Int8(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Int8(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Int8(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt8_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int8
		i    uint
		want Int8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
		{1, 2, 4},
		{1, 3, 8},
		{1, 4, 16},
		{1, 5, 32},
		{1, 6, 64},
		{1, 7, -128},
		{1, 8, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int8(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt8_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int8
		i    uint
		want Int8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{64, 0, 64},
		{64, 1, 32},
		{64, 2, 16},
		{64, 3, 8},
		{64, 4, 4},
		{64, 5, 2},
		{64, 6, 1},
		{64, 7, 0},

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
			t.Errorf("Int8(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt8_Sign(t *testing.T) {
	testCases := []struct {
		x    Int8
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 1},
		{-128, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int8(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt8_Neg(t *testing.T) {
	testCases := []struct {
		x    Int8
		want Int8
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{127, -127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int8(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt8_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int8
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{127, -128, 1},
		{-128, 127, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int8(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestInt8_Text(t *testing.T) {
	var b big.Int
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int8(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt8_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int8(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt8_AppendText(t *testing.T) {
	var buf []byte
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int8(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestInt8_String(t *testing.T) {
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int8(%d).String() = %q, want %q", i, got, want)
		}
	}
}

func TestInt8_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int8
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
			-128,
			"-128",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
