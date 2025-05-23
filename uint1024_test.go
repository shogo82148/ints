package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func uint1024ToBigInt(a Uint1024) *big.Int {
	b := new(big.Int)
	b.SetUint64(a[0])
	for i := 1; i < len(a); i++ {
		b.Lsh(b, 64)
		b.Add(b, new(big.Int).SetUint64(a[i]))
	}
	return b
}

func FuzzUint1024_Add(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// MaxUint1024
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 1024)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		b := Uint1024{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
		got := uint1024ToBigInt(a.Add(b))

		ba := uint1024ToBigInt(a)
		bb := uint1024ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint1024(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint1024_Add(b *testing.B) {
	const M = 0xffffffff_ffffffff
	x := Uint1024{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	y := Uint1024{0x7f000000_00000000, M, M, M, M, M, M, M, M, M, M, M, M, M, M, M}

	b.Run("Uint1024", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Add(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint1024ToBigInt(x)
		yy := uint1024ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Add(xx, yy)
		}
	})
}

func FuzzUint1024_Sub(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 1024)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		b := Uint1024{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
		got := uint1024ToBigInt(a.Sub(b))

		ba := uint1024ToBigInt(a)
		bb := uint1024ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint1024(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint1024_Mul(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// MaxUint1024
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// MaxUint1024
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 1024)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		b := Uint1024{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
		got := uint1024ToBigInt(a.Mul(b))

		ba := uint1024ToBigInt(a)
		bb := uint1024ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint1024(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint1024_Mul(b *testing.B) {
	const M = 0xffffffff_ffffffff
	x := Uint1024{0, 0, 0, 0, 0, 0, 0, 0, M, M, M, M, M, M, M, M}
	y := Uint1024{0, 0, 0, 0, 0, 0, 0, 0, M, M, M, M, M, M, M, M}

	b.Run("Uint1024", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Mul(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint1024ToBigInt(x)
		yy := uint1024ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Mul(xx, yy)
		}
	})
}

func FuzzUint1024_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(127),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10),
	)

	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10),
	)

	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(2),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		b := Uint1024{v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := uint1024ToBigInt(q)
		gotR := uint1024ToBigInt(r)

		ba := uint1024ToBigInt(a)
		bb := uint1024ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint1024(%s).DivMod(%s) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = uint1024ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint1024(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = uint1024ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint1024(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func BenchmarkUint1024_DivMod(b *testing.B) {
	const M = 0xffffffff_ffffffff
	x := Uint1024{M, M, M, M, M, M, M, M, M, M, M, M, M, M, M, M}
	y := Uint1024{1, M, M, M, M, M, M, M, M, M, M, M, M, M, M, M}

	b.Run("Uint1024", func(b *testing.B) {
		for b.Loop() {
			x.DivMod(y)
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint1024ToBigInt(x)
		yy := uint1024ToBigInt(y)
		q := new(big.Int)
		r := new(big.Int)
		for b.Loop() {
			q.DivMod(xx, yy, r)
		}
	})
}

func TestUint1024_And(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		y    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint1024(%s).And(%s) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint1024_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		y    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint1024(%s).AndNot(%s) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint1024_Or(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		y    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint1024(%s).Or(%s) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint1024_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		y    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint1024(%s).Xor(%s) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint1024_Not(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint1024(%s).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint1024_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		i    uint
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			0,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			64,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			65,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			128,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			129,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			192,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			193,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			256,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			257,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			320,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			321,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			384,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			385,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			448,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			449,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			512,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			513,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			576,
			Uint1024{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			577,
			Uint1024{2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			640,
			Uint1024{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			641,
			Uint1024{2, 2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			704,
			Uint1024{1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			705,
			Uint1024{2, 2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			768,
			Uint1024{1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			769,
			Uint1024{2, 2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			832,
			Uint1024{1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			833,
			Uint1024{2, 2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			896,
			Uint1024{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			897,
			Uint1024{2, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			960,
			Uint1024{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			961,
			Uint1024{2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint1024(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint1024_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		i    uint
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			0,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			1,
			Uint1024{0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			64,
			Uint1024{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			65,
			Uint1024{0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			128,
			Uint1024{0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			129,
			Uint1024{0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			192,
			Uint1024{0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			193,
			Uint1024{0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			256,
			Uint1024{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			257,
			Uint1024{0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			320,
			Uint1024{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			321,
			Uint1024{0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			384,
			Uint1024{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			385,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			448,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			449,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			512,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			513,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			576,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			577,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			640,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			641,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			704,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			705,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			768,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			769,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			832,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			833,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			896,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			897,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x80000000_00000000},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			960,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			Uint1024{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			961,
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint1024(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint1024_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want int
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			1024,
		},
		{
			Uint1024{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			63,
		},
		{
			Uint1024{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			127,
		},
		{
			Uint1024{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			191,
		},
		{
			Uint1024{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			255,
		},
		{
			Uint1024{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			319,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			383,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			447,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			511,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			575,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			639,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			703,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			767,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			831,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			895,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			959,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			1023,
		},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint1024(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint1024_TrailingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want int
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			1024,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			0,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			64,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			128,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			192,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			256,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			320,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			384,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			448,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			512,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			576,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			640,
		},
		{
			Uint1024{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			704,
		},
		{
			Uint1024{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			768,
		},
		{
			Uint1024{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			832,
		},
		{
			Uint1024{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			896,
		},
		{
			Uint1024{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			960,
		},
	}

	for _, tc := range testCases {
		got := tc.x.TrailingZeros()
		if got != tc.want {
			t.Errorf("Uint1024(%d).TrailingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint1024_BitLen(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want int
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			1,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			65,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			129,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
			193,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			257,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			321,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0},
			385,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
			449,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			513,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			577,
		},
		{
			Uint1024{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			641,
		},
		{
			Uint1024{0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			705,
		},
		{
			Uint1024{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			769,
		},
		{
			Uint1024{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			833,
		},
		{
			Uint1024{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			897,
		},
		{
			Uint1024{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			961,
		},
	}

	for _, tc := range testCases {
		got := tc.x.BitLen()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).BitLen() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint1024_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want int
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			// 1
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			1,
		},
		{
			// MaxUint1024
			Uint1024{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
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

func TestUint1024_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint1024
		want Uint1024
	}{
		{
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			// 1
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},

			// MaxUint1024
			Uint1024{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint1024(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint1024_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint1024
		want int
	}{
		{
			// 0
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			// 0
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			// 0
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			// 1
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			-1,
		},
		{
			// 1
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			// 0
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint1024(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzUint1024_Text(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		10,
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		2,
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		62,
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		got := a.Text(base)

		b := uint1024ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint1024(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint1024_Append(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		10,
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint1024ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint1024(%s).Append(buf, %d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint1024_AppendText(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint1024ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint1024(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzUint1024_String(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64) {
		a := Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		got := a.String()

		b := uint1024ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Uint1024(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkUint1024_Text2(b *testing.B) {
	a := Uint1024{
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
	}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint1024_Text10(b *testing.B) {
	a := Uint1024{
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
	}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkUint1024_Text62(b *testing.B) {
	a := Uint1024{
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
	}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkUint1024_String(b *testing.B) {
	a := Uint1024{
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
		math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64,
	}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestUint1024_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint1024
		want   string
	}{
		// decimal
		{
			"%d",
			Uint1024{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
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
