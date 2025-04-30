package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func int512ToBigInt(a Int512) *big.Int {
	b := new(big.Int)
	b.SetInt64(int64(a[0]))
	for i := 1; i < len(a); i++ {
		b.Lsh(b, 64)
		b.Add(b, new(big.Int).SetUint64(a[i]))
	}
	return b
}

func FuzzInt512_Add(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// -1
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		// MaxInt512
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)
	f.Add(
		// MinInt512
		uint64(1<<63), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// -1
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	base := new(big.Int).Lsh(big.NewInt(1), 512-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Int512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := int512ToBigInt(a.Add(b))

		ba := int512ToBigInt(a)
		bb := int512ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int512(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt512_Sub(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// MaxInt512
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// -1
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)
	f.Add(
		// MinInt512
		uint64(1<<63), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	base := new(big.Int).Lsh(big.NewInt(1), 512-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Int512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := int512ToBigInt(a.Sub(b))

		ba := int512ToBigInt(a)
		bb := int512ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int512(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt512_Mul(f *testing.F) {
	f.Add(
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
	)
	f.Add(
		// MaxInt512
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
		// -1
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64),
	)
	f.Add(
		// MinInt512
		uint64(1<<63), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		// 1
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1),
	)

	base := new(big.Int).Lsh(big.NewInt(1), 512-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Int512{v0, v1, v2, v3, v4, v5, v6, v7}
		got := int512ToBigInt(a.Mul(b))

		ba := int512ToBigInt(a)
		bb := int512ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int512(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt512_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10), // 10
	)

	base := new(big.Int).Lsh(big.NewInt(1), 512-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Int512{v0, v1, v2, v3, v4, v5, v6, v7}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := int512ToBigInt(q)
		gotR := int512ToBigInt(r)

		ba := int512ToBigInt(a)
		bb := int512ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Int512(%s).DivMod(%s) = (%d, %d), want (%d, %d)", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = int512ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Int512(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = int512ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Int512(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func FuzzInt512_QuoRem(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0),
		uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(10), // 10
	)

	base := new(big.Int).Lsh(big.NewInt(1), 512-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 512)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7, v0, v1, v2, v3, v4, v5, v6, v7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		b := Int512{v0, v1, v2, v3, v4, v5, v6, v7}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.QuoRem(b)
		gotQ := int512ToBigInt(q)
		gotR := int512ToBigInt(r)

		ba := int512ToBigInt(a)
		bb := int512ToBigInt(b)
		wantQ, wantR := new(big.Int).QuoRem(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Int512(%s).QuoRem(%s) = (%d, %d), want (%d, %d)", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Quo(b)
		gotQ = int512ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Int512(%s).Quo(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Rem(b)
		gotR = int512ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Int512(%s).Rem(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func TestInt512_And(t *testing.T) {
	testCases := []struct {
		a    Int512
		b    Int512
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{1, 1, 1, 1, 1, 1, 1, 1}, Int512{2, 2, 2, 2, 2, 2, 2, 2}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.And(tc.b)
		if got != tc.want {
			t.Errorf("Int512(%d).And(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestInt512_AndNot(t *testing.T) {
	testCases := []struct {
		a    Int512
		b    Int512
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{1, 1, 1, 1, 1, 1, 1, 1}, Int512{2, 2, 2, 2, 2, 2, 2, 2}, Int512{1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, tc := range testCases {
		got := tc.a.AndNot(tc.b)
		if got != tc.want {
			t.Errorf("Int512(%d).AndNot(%d) = %d", tc.a, tc.b, got)
		}
	}
}

func TestInt512_Or(t *testing.T) {
	testCases := []struct {
		a    Int512
		b    Int512
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{1, 1, 1, 1, 1, 1, 1, 1}, Int512{2, 2, 2, 2, 2, 2, 2, 2}, Int512{3, 3, 3, 3, 3, 3, 3, 3}},
	}

	for _, tc := range testCases {
		got := tc.a.Or(tc.b)
		if got != tc.want {
			t.Errorf("Int512(%d).Or(%d) = %d", tc.a, tc.b, got)
		}
	}
}

func TestInt512_Xor(t *testing.T) {
	testCases := []struct {
		a    Int512
		b    Int512
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{1, 1, 1, 1, 1, 1, 1, 1}, Int512{2, 2, 2, 2, 2, 2, 2, 2}, Int512{3, 3, 3, 3, 3, 3, 3, 3}},
	}

	for _, tc := range testCases {
		got := tc.a.Xor(tc.b)
		if got != tc.want {
			t.Errorf("Int512(%d).Xor(%d) = %d", tc.a, tc.b, got)
		}
	}
}

func TestInt512_Not(t *testing.T) {
	testCases := []struct {
		x    Int512
		want Int512
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			Int512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Int512(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt512_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int512
		i    uint
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, 0, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{1, 1, 1, 1, 1, 1, 1, 1}, 1, Int512{2, 2, 2, 2, 2, 2, 2, 2}},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int512(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt512_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int512
		i    uint
		want Int512
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			1,
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Int512{2, 2, 2, 2, 2, 2, 2, 2},
			1,
			Int512{1, 1, 1, 1, 1, 1, 1, 1},
		},

		// sign extension
		{
			Int512{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0},
			1,
			Int512{0xc0000000_00000000, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Int512{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0},
			64,
			Int512{0xffffffff_ffffffff, 0x80000000_00000000, 0, 0, 0, 0, 0, 0},
		},
		{
			Int512{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0},
			128,
			Int512{0xffffffff_ffffffff, 0xffffffff_ffffffff, 0x80000000_00000000, 0, 0, 0, 0, 0},
		},
		{
			Int512{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0},
			192,
			Int512{0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff, 0x80000000_00000000, 0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int512(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt512_Neg(t *testing.T) {
	testCases := []struct {
		x    Int512
		want Int512
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 1},
			Int512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int512(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt512_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int512
		want int
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 1},
			Int512{0, 0, 0, 0, 0, 0, 0, 2},
			-1,
		},
		{
			Int512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
			Int512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64 - 1},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int512(%s).Cmp(%s) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzInt512_Text(f *testing.F) {
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

		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		got := a.Text(base)

		b := int512ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Int512(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzInt512_Append(f *testing.F) {
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

		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		buf := a.Append(nil, base)
		got := string(buf)

		b := int512ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Int512(%s).Append(buf, %d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func TestInt512_Sign(t *testing.T) {
	testCases := []struct {
		x    Int512
		want int
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			0,
		},
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 1},
			1,
		},
		{
			Int512{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
			-1,
		},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int512(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func FuzzInt512_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := int512ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Int512(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzInt512_String(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), uint64(0), uint64(0), uint64(1), uint64(0))

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, u4, u5, u6, u7 uint64) {
		a := Int512{u0, u1, u2, u3, u4, u5, u6, u7}
		got := a.String()

		b := int512ToBigInt(a)
		want := b.String()

		if got != want {
			t.Errorf("Int512(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkInt512_Text2(b *testing.B) {
	a := Int512{1 << 63, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkInt512_Text10(b *testing.B) {
	a := Int512{1 << 63, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkInt512_Text62(b *testing.B) {
	a := Int512{1 << 63, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkInt512_String(b *testing.B) {
	a := Int512{1 << 63, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestInt512_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int512
		want   string
	}{
		// decimal
		{
			"%d",
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			"0",
		},
		{
			"%d",
			Int512{0x80000000_00000000, 0, 0, 0, 0, 0, 0, 0},
			"-6703903964971298549787012499102923063739682910296196688861780721860882015036773488400937149083451713845015929093243025426876941405973284973216824503042048",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
