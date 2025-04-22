package ints

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestUint32_Add(t *testing.T) {
	testCases := []struct {
		x, y uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 0, 1},
	}

	for _, tc := range testCases {
		a := Uint32(tc.x)
		b := Uint32(tc.y)
		got := a.Add(b)
		if got != tc.want {
			t.Errorf("Uint32(%d).Add(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint32_Sub(t *testing.T) {
	testCases := []struct {
		x, y uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 0},
		{1, 0, 1},
		{0, 1, math.MaxUint32},
	}

	for _, tc := range testCases {
		a := Uint32(tc.x)
		b := Uint32(tc.y)
		got := a.Sub(b)
		if got != tc.want {
			t.Errorf("Uint32(%d).Sub(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint32_Mul(t *testing.T) {
	testCases := []struct {
		x, y uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		a := Uint32(tc.x)
		b := Uint32(tc.y)
		got := a.Mul(b)
		if got != tc.want {
			t.Errorf("Uint32(%d).Mul(%d) = %d, want %d", a, b, got, tc.want)
		}
	}
}

func TestUint32_DivMod(t *testing.T) {
	testCases := []struct {
		x, y Uint32
		z, r Uint32
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		z, r := tc.x.DivMod(tc.y)
		if z != tc.z {
			t.Errorf("Uint32(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		if r != tc.r {
			t.Errorf("Uint32(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		z = tc.x.Div(tc.y)
		if z != tc.z {
			t.Errorf("Uint32(%d).Div(%d) = %d, want %d", tc.x, tc.y, z, tc.z)
		}
		r = tc.x.Mod(tc.y)
		if r != tc.r {
			t.Errorf("Uint32(%d).Mod(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint32_QuoRem(t *testing.T) {
	testCases := []struct {
		x, y Uint32
		q, r Uint32
	}{
		{0, 1, 0, 0},
		{100, 10, 10, 0},
		{127, 10, 12, 7},
	}

	for _, tc := range testCases {
		q, r := tc.x.QuoRem(tc.y)
		if q != tc.q {
			t.Errorf("Uint32(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		if r != tc.r {
			t.Errorf("Uint32(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}

		q = tc.x.Quo(tc.y)
		if q != tc.q {
			t.Errorf("Uint32(%d).Quo(%d) = %d, want %d", tc.x, tc.y, q, tc.q)
		}
		r = tc.x.Rem(tc.y)
		if r != tc.r {
			t.Errorf("Uint32(%d).Rem(%d) = %d, want %d", tc.x, tc.y, r, tc.r)
		}
	}
}

func TestUint32_And(t *testing.T) {
	testCases := []struct {
		x    Uint32
		y    Uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint32(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint32_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint32
		y    Uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint32(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint32_Or(t *testing.T) {
	testCases := []struct {
		x    Uint32
		y    Uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint32(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint32_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint32
		y    Uint32
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint32(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint32_Not(t *testing.T) {
	testCases := []struct {
		x    Uint32
		want Uint32
	}{
		{0, math.MaxUint32},
		{1, math.MaxUint32 - 1},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint32(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint32_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint32
		i    uint
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 2},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint32(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint32_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint32
		i    uint
		want Uint32
	}{
		{0, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint32(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint32_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint32
		want int
	}{
		{0, 32},
		{1, 31},
		{2, 30},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint32(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint32_TrailingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint32
		want int
	}{
		{0, 32},
		{1, 0},
		{2, 1},
	}

	for _, tc := range testCases {
		got := tc.x.TrailingZeros()
		if got != tc.want {
			t.Errorf("Uint32(%d).TrailingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint32_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint32
		want int
	}{
		{0, 0},
		{1, 1},
		{math.MaxUint32, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint32(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint32_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint32
		want Uint32
	}{
		{0, 0},
		{1, math.MaxUint32},
		{math.MaxUint32, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint32(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint32_Cmp(t *testing.T) {
	testCases := []struct {
		x    Uint32
		y    Uint32
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxUint32, math.MaxUint32, 0},
	}

	for _, tc := range testCases {
		got := tc.x.Cmp(tc.y)
		if got != tc.want {
			t.Errorf("Uint32(%d).Cmp(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func FuzzUint32_Text(f *testing.F) {
	f.Add(uint32(0), 10)
	f.Add(uint32(0), 62)
	f.Add(uint32(math.MaxUint32), 2)
	f.Add(uint32(math.MaxUint32), 10)
	f.Add(uint32(math.MaxUint32), 62)
	f.Fuzz(func(t *testing.T, x uint32, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint32(x)
		b := big.NewInt(int64(x))
		got := a.Text(base)
		want := b.Text(base)
		if got != want {
			t.Errorf("Uint32(%d).Text(%d) = %q, want %q", x, base, got, want)
		}
	})
}

func FuzzUint32_Append(f *testing.F) {
	f.Add(uint32(0), 10)
	f.Add(uint32(0), 62)
	f.Add(uint32(math.MaxUint32), 2)
	f.Add(uint32(math.MaxUint32), 10)
	f.Add(uint32(math.MaxUint32), 62)
	f.Fuzz(func(t *testing.T, x uint32, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint32(x)
		b := big.NewInt(int64(x))
		got := a.Append(nil, base)
		want := b.Text(base)
		if string(got) != want {
			t.Errorf("Uint32(%d).Append(buf, %d) = %q, want %q", x, base, string(got), want)
		}
	})
}

func FuzzUint32_AppendText(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, x uint32) {
		a := Uint32(x)
		b := big.NewInt(int64(x))
		got, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint32(%d).AppendText(buf) = %q, want %q", x, string(got), want)
		}
	})
}

func FuzzUint32_String(f *testing.F) {
	f.Add(uint32(0))
	f.Add(uint32(math.MaxUint32))
	f.Fuzz(func(t *testing.T, x uint32) {
		a := Uint32(x)
		b := big.NewInt(int64(x))
		got := a.String()
		want := b.String()
		if string(got) != want {
			t.Errorf("Uint32(%d).String() = %q, want %q", x, got, want)
		}
	})
}

func TestUint32_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint32
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
