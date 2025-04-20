package ints

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestInt16_Add(t *testing.T) {
	testCases := []struct {
		x, y int16
		want Int16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		a := Int16(tc.x)
		b := Int16(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Int16(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt16_Sub(t *testing.T) {
	testCases := []struct {
		x, y int16
		want Int16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
	}

	for _, tc := range testCases {
		a := Int16(tc.x)
		b := Int16(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Int16(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt16_Mul(t *testing.T) {
	testCases := []struct {
		x, y int16
		want Int16
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
		{-1, -1, 1},
	}

	for _, tc := range testCases {
		a := Int16(tc.x)
		b := Int16(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Int16(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestInt16_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Int16
		z, r Int16
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
			t.Errorf("Int16(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Int16(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Int16(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Int16(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt16_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Int16
		q, r Int16
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
			t.Errorf("Int16(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Int16(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Int16(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Int16(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestInt16_And(t *testing.T) {
	testCases := []struct {
		x    Int16
		y    Int16
		want Int16
	}{
		{0, 0, 0},
		{1, 1, 1},
		{-1, -1, -1},
		{127, 255, 127},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Int16(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt16_AndNot(t *testing.T) {
	testCases := []struct {
		x    Int16
		y    Int16
		want Int16
	}{
		{0, 0, 0},
		{1, 1, 0},
		{-1, -1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Int16(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt16_Or(t *testing.T) {
	testCases := []struct {
		x    Int16
		y    Int16
		want Int16
	}{
		{0, 0, 0},
		{1, 1, 1},
		{-1, -1, -1},
		{127, 255, 255},
		{-128, -255, -127},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Int16(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt16_Xor(t *testing.T) {
	testCases := []struct {
		x    Int16
		y    Int16
		want Int16
	}{
		{0, 0, 0},
		{1, 1, 0},
		{-1, -1, 0},
		{127, 255, 128},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Int16(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt16_Not(t *testing.T) {
	testCases := []struct {
		x    Int16
		want Int16
	}{
		{0, -1},
		{1, -2},
		{-1, 0},
		{127, -128},
		{-128, 127},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Int16(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt16_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int16
		i    uint
		want Int16
	}{
		{0, 1, 0},
		{1, 1, 2},
		{-1, 1, -2},
		{127, 1, 254},
		{-128, 1, -256},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int16(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt16_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int16
		i    uint
		want Int16
	}{
		{0, 1, 0},
		{1, 1, 0},
		{-1, 1, -1},
		{127, 1, 63},
		{-128, 1, -64},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int16(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt16_Sign(t *testing.T) {
	testCases := []struct {
		x    Int16
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{math.MaxInt16, 1},
		{math.MinInt16, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int16(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt16_Neg(t *testing.T) {
	testCases := []struct {
		x    Int16
		want Int16
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{math.MaxInt16, -math.MaxInt16},
		{math.MinInt16, math.MinInt16},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int16(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt16_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int16
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxInt16, math.MinInt16, 1},
		{math.MinInt16, math.MaxInt16, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int16(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestInt16_Text(t *testing.T) {
	var b big.Int
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int16(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt16_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int16(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt16_AppendText(t *testing.T) {
	var buf []byte
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int16(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestInt16_String(t *testing.T) {
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int16(%d).String() = %q, want %q", i, got, want)
		}
	}
}

func TestInt16_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int16
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
			-32768,
			"-32768",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
