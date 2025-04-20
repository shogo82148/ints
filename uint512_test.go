package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func uint512ToBigInt(a Uint512) *big.Int {
	b := new(big.Int)
	b.SetUint64(a[0])
	for i := 1; i < len(a); i++ {
		b.Lsh(b, 64)
		b.Add(b, new(big.Int).SetUint64(a[i]))
	}
	return b
}

func FuzzUint512_Add(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// MaxUint512
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Uint512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := uint512ToBigInt(a.Add(b))

		ba := uint512ToBigInt(a)
		bb := uint512ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint512(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint512_Sub(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Uint512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := uint512ToBigInt(a.Sub(b))

		ba := uint512ToBigInt(a)
		bb := uint512ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint512(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint512_Mul(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		// MaxUint512
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// MaxUint512
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Uint512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := uint512ToBigInt(a.Mul(b))

		ba := uint512ToBigInt(a)
		bb := uint512ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint512(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint512_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(127),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10),
	)
	f.Add(
		uint64(127), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(10), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Uint512{v0, v1, v2, v3, v4, v5, v6, v7}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := uint512ToBigInt(q)
		gotR := uint512ToBigInt(r)

		ba := uint512ToBigInt(a)
		bb := uint512ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint512(%s).DivMod(%s) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = uint512ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint512(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = uint512ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint512(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func TestUint512_And(t *testing.T) {
	testCases := []struct {
		x    Uint512
		y    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint512(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint512_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint512
		y    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint512(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint512_Or(t *testing.T) {
	testCases := []struct {
		x    Uint512
		y    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint512(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint512_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint512
		y    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint512(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint512_Not(t *testing.T) {
	testCases := []struct {
		x    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint512(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint512_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint512
		i    uint
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			1,
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			1,
			Uint512{2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			Uint512{0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
			1,
			Uint512{1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			65,
			Uint512{2, 2, 2, 2, 2, 2, 2, 0},
		},
		{
			Uint512{0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
			65,
			Uint512{1, 1, 1, 1, 1, 1, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			129,
			Uint512{2, 2, 2, 2, 2, 2, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			193,
			Uint512{2, 2, 2, 2, 2, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			257,
			Uint512{2, 2, 2, 2, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			321,
			Uint512{2, 2, 2, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			385,
			Uint512{2, 2, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			449,
			Uint512{2, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint512(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint512_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint512
		i    uint
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			1,
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			1,
			Uint512{0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			64,
			Uint512{0, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			65,
			Uint512{0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			128,
			Uint512{0, 0, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			192,
			Uint512{0, 0, 0, 1, 1, 1, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			256,
			Uint512{0, 0, 0, 0, 1, 1, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			320,
			Uint512{0, 0, 0, 0, 0, 1, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			384,
			Uint512{0, 0, 0, 0, 0, 0, 1, 1},
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			448,
			Uint512{0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint512(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint512_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint512
		want int
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 1},
			1,
		},
		{
			Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint512(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint512_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint512
		want Uint512
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 1},
			Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint512(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint512_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint512
		want int
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			Uint512{1, 0, 0, 0, 0, 0, 0, 0},
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			-1,
		},
		{
			Uint512{1, 1, 1, 1, 1, 1, 1, 1},
			Uint512{1, 0, 0, 0, 0, 0, 0, 0},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint512(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzUint512_Text(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)

	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)

	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		got := a.Text(base)

		b := uint512ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint512(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzUint512_Append(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 10)

	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 2)

	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint512ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint512(%s).Append(buf, %d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint512_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint512ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint512(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzUint512_String(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64) {
		a := Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
		got := a.String()

		b := uint512ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint512(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkUint512_Text2(b *testing.B) {
	a := Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint512_Text10(b *testing.B) {
	a := Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkUint512_Text62(b *testing.B) {
	a := Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkUint512_String(b *testing.B) {
	a := Uint512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestUint512_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint512
		want   string
	}{
		// decimal
		{
			"%d",
			Uint512{0, 0},
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
