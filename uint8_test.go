package ints

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestUint8_Add(t *testing.T) {
	testCases := []struct {
		x, y uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		a := Uint8(tc.x)
		b := Uint8(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Uint8(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint8_Sub(t *testing.T) {
	testCases := []struct {
		x, y uint8
		want Uint8
	}{
		{0, 0, 0},
		{0, 1, 255},
	}

	for _, tc := range testCases {
		a := Uint8(tc.x)
		b := Uint8(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Uint8(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint8_Mul(t *testing.T) {
	testCases := []struct {
		x, y uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		a := Uint8(tc.x)
		b := Uint8(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Uint8(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint8_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Uint8
		z, r Uint8
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Uint8(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Uint8(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Uint8(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Uint8(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint8_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Uint8
		q, r Uint8
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Uint8(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Uint8(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Uint8(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Uint8(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint8_And(t *testing.T) {
	testCases := []struct {
		x    Uint8
		y    Uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint8(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint8_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint8
		y    Uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint8(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint8_Or(t *testing.T) {
	testCases := []struct {
		x    Uint8
		y    Uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint8(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint8_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint8
		y    Uint8
		want Uint8
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint8(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint8_Not(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want Uint8
	}{
		{0, 0xff},
		{1, 0xfe},
		{0xff, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint8(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint8
		i    uint
		want Uint8
	}{
		{0, 0, 0},
		{1, 1, 2},
		{1, 7, 128},
		{1, 8, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint8(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint8_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint8
		i    uint
		want Uint8
	}{
		{0, 0, 0},
		{1, 1, 0},
		{0xff, 1, 0x7f},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint8(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint8_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want int
	}{
		{0, 8},
		{1, 7},
		{0x80, 0},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint8(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_TrailingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want int
	}{
		{0, 8},
		{1, 0},
		{0x80, 7},
	}

	for _, tc := range testCases {
		got := tc.x.TrailingZeros()
		if got != tc.want {
			t.Errorf("Uint8(%d).TrailingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_BitLen(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want int
	}{
		{0, 0},
		{1, 1},
		{0x80, 8},
	}

	for _, tc := range testCases {
		got := tc.x.BitLen()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).BitLen() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want int
	}{
		{0, 0},
		{1, 1},
		{math.MaxUint8, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint8(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint8
		want Uint8
	}{
		{0, 0},
		{1, 0xff},
		{0xff, 1},
		{127, 129},
		{128, 128},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint8(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint8_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint8
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxUint8, math.MaxUint8, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint8(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestUint8_Text(t *testing.T) {
	var b big.Int
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint8(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint8_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint8(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint8_AppendText(t *testing.T) {
	var buf []byte
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint8(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestUint8_String(t *testing.T) {
	for i := range math.MaxUint8 + 1 {
		a := Uint8(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint8(%d).String() = %q, want %q", i, got, want)
		}
	}
}

func TestUint8_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint8
		want   string
	}{
		// decimal
		{
			"%d",
			0,
			"0",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
