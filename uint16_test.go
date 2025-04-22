package ints

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestUint16_Add(t *testing.T) {
	testCases := []struct {
		x, y uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		a := Uint16(tc.x)
		b := Uint16(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Uint16(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint16_Sub(t *testing.T) {
	testCases := []struct {
		x, y uint16
		want Uint16
	}{
		{0, 0, 0},
		{0, 1, 0xffff},
		{math.MaxUint16, 1, Uint16(math.MaxUint16 - 1)},
	}

	for _, tc := range testCases {
		a := Uint16(tc.x)
		b := Uint16(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Uint16(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint16_Mul(t *testing.T) {
	testCases := []struct {
		x, y uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		a := Uint16(tc.x)
		b := Uint16(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Uint16(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint16_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Uint16
		z, r Uint16
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Uint16(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Uint16(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Uint16(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Uint16(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint16_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Uint16
		q, r Uint16
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Uint16(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Uint16(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Uint16(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Uint16(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint16_And(t *testing.T) {
	testCases := []struct {
		x    Uint16
		y    Uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint16(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint16_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint16
		y    Uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint16(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint16_Or(t *testing.T) {
	testCases := []struct {
		x    Uint16
		y    Uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint16(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint16_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint16
		y    Uint16
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint16(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint16_Not(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want Uint16
	}{
		{0, math.MaxUint16},
		{1, math.MaxUint16 - 1},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint16(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}
func TestUint16_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint16
		i    uint
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint16(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint16_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint16
		i    uint
		want Uint16
	}{
		{0, 0, 0},
		{1, 0, 1},
		{2, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint16(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint16_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want int
	}{
		{0, 16},
		{1, 15},
		{2, 14},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint16(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint16_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want int
	}{
		{0, 0},
		{1, 1},
		{math.MaxUint16, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint16(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint16_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want Uint16
	}{
		{0, 0},
		{1, math.MaxUint16},
		{math.MaxUint16, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint16(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint16_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint16
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxUint16, math.MaxUint16, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint16(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestUint16_Text(t *testing.T) {
	var b big.Int
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint16(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint16_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint16(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint16_AppendText(t *testing.T) {
	var buf []byte
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint16(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestUint16_String(t *testing.T) {
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint16(%d).String() = %q, want %q", i, got, want)
		}
	}
}

func TestUint16_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint16
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
