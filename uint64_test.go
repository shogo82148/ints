package ints

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestUint64_Add(t *testing.T) {
	testCases := []struct {
		x, y uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		a := Uint64(tc.x)
		b := Uint64(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Uint64(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint64_Sub(t *testing.T) {
	testCases := []struct {
		x, y uint64
		want Uint64
	}{
		{0, 0, 0},
		{0, 1, math.MaxUint64},
		{math.MaxUint64, 1, Uint64(math.MaxUint64 - 1)},
	}

	for _, tc := range testCases {
		a := Uint64(tc.x)
		b := Uint64(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Uint64(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint64_Mul(t *testing.T) {
	testCases := []struct {
		x, y uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		a := Uint64(tc.x)
		b := Uint64(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Uint64(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint64_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Uint64
		z, r Uint64
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Uint64(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Uint64(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Uint64(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Uint64(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint64_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Uint64
		q, r Uint64
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Uint64(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Uint64(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Uint64(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Uint64(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint64_And(t *testing.T) {
	testCases := []struct {
		x    Uint64
		y    Uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint64(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint64_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint64
		y    Uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint64(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint64_Or(t *testing.T) {
	testCases := []struct {
		x    Uint64
		y    Uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint64(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint64_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint64
		y    Uint64
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint64(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint64_Not(t *testing.T) {
	testCases := []struct {
		x    Uint64
		want Uint64
	}{
		{0, math.MaxUint64},
		{1, math.MaxUint64 - 1},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint64(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint64_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint64
		i    uint
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint64(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint64_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint64
		i    uint
		want Uint64
	}{
		{0, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint64(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint64_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint64
		want int
	}{
		{0, 64},
		{1, 63},
		{math.MaxUint64, 0},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint64(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint64_TrailingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint64
		want int
	}{
		{0, 64},
		{1, 0},
		{math.MaxUint64, 0},
	}

	for _, tc := range testCases {
		got := tc.x.TrailingZeros()
		if got != tc.want {
			t.Errorf("Uint64(%d).TrailingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint64_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint64
		want int
	}{
		{0, 0},
		{1, 1},
		{math.MaxUint64, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint64(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint64_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint64
		want Uint64
	}{
		{0, 0},
		{1, math.MaxUint64},
		{math.MaxUint64, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint64(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint64_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint64
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxUint64, math.MaxUint64, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint64(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzUint64_Text(f *testing.F) {
	f.Add(uint64(0), 10)
	f.Add(uint64(0), 62)
	f.Add(uint64(math.MaxUint64), 2)
	f.Add(uint64(math.MaxUint64), 10)
	f.Add(uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, x uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.Text(base)
		want := b.Text(base)
		if got != want {
			t.Errorf("Uint64(%d).Text(%d) = %q, want %q", x, base, got, want)
		}
	})
}

func FuzzUint64_Append(f *testing.F) {
	f.Add(uint64(0), 10)
	f.Add(uint64(0), 62)
	f.Add(uint64(math.MaxUint64), 2)
	f.Add(uint64(math.MaxUint64), 10)
	f.Add(uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, x uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.Append(nil, base)
		want := b.Text(base)
		if string(got) != want {
			t.Errorf("Uint64(%d).Append(buf, %d) = %q, want %q", x, base, string(got), want)
		}
	})
}

func FuzzUint64_AppendText(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, x uint64) {
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint64(%d).AppendText(buf) = %q, want %q", x, string(got), want)
		}
	})
}

func FuzzUint64_String(f *testing.F) {
	f.Add(uint64(0))
	f.Add(uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, x uint64) {
		a := Uint64(x)
		var b big.Int
		b.SetUint64(x)
		got := a.String()
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint64(%d).String() = %q, want %q", x, got, want)
		}
	})
}

func TestUint64_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint64
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
