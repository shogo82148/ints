package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func int256ToBigInt(a Int256) *big.Int {
	var b, c big.Int
	b.SetInt64(int64(a[0]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[1]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[2]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[3]))
	return &b
}

func FuzzInt256_Add(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxInt256
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63), uint64(0), uint64(0), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 256-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Int256{u0, u1, u2, u3}
		b := Int256{v0, v1, v2, v3}
		got := int256ToBigInt(a.Add(b))

		ba := int256ToBigInt(a)
		bb := int256ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int256(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt256_Sub(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxInt256
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63), uint64(0), uint64(0), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 256-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Int256{u0, u1, u2, u3}
		b := Int256{v0, v1, v2, v3}
		got := int256ToBigInt(a.Sub(b))

		ba := int256ToBigInt(a)
		bb := int256ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int256(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt256_Mul(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxInt256
	)
	f.Add(
		uint64(1<<63), uint64(0), uint64(0), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 256-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Int256{u0, u1, u2, u3}
		b := Int256{v0, v1, v2, v3}
		got := int256ToBigInt(a.Mul(b))

		ba := int256ToBigInt(a)
		bb := int256ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int256(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt256_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxInt256
	)
	f.Add(
		uint64(1<<63), uint64(0), uint64(0), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 256-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Int256{u0, u1, u2, u3}
		b := Int256{v0, v1, v2, v3}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := int256ToBigInt(q)
		gotR := int256ToBigInt(r)

		ba := int256ToBigInt(a)
		bb := int256ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Int256(%s).DivMod(%s) = (%d, %d), want (%d, %d)", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = int256ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Int256(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = int256ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Int256(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func FuzzInt256_QuoRem(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(1<<63-1), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxInt256
	)
	f.Add(
		uint64(1<<63), uint64(0), uint64(0), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 256-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Int256{u0, u1, u2, u3}
		b := Int256{v0, v1, v2, v3}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.QuoRem(b)
		gotQ := int256ToBigInt(q)
		gotR := int256ToBigInt(r)

		ba := int256ToBigInt(a)
		bb := int256ToBigInt(b)
		wantQ, wantR := new(big.Int).QuoRem(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Int256(%s).QuoRem(%s) = (%d, %d), want (%d, %d)", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Quo(b)
		gotQ = int256ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Int256(%s).Quo(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Rem(b)
		gotR = int256ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Int256(%s).Rem(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func TestInt256_And(t *testing.T) {
	testCases := []struct {
		x    Int256
		y    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Int256(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt256_AndNot(t *testing.T) {
	testCases := []struct {
		x    Int256
		y    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}, Int256{0, 0, 0, 0}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Int256(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt256_Or(t *testing.T) {
	testCases := []struct {
		x    Int256
		y    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Int256(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt256_Xor(t *testing.T) {
	testCases := []struct {
		x    Int256
		y    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{1, 1, 1, 1}, Int256{1, 1, 1, 1}, Int256{0, 0, 0, 0}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Int256(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt256_Not(t *testing.T) {
	testCases := []struct {
		x    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}},
		{Int256{1, 1, 1, 1}, Int256{math.MaxUint64 - 1, math.MaxUint64 - 1, math.MaxUint64 - 1, math.MaxUint64 - 1}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{0, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Int256(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt256_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int256
		i    uint
		want Int256
	}{
		{Int256{0, 0, 0, 0}, 0, Int256{0, 0, 0, 0}},
		{Int256{1, 0, 0, 0}, 1, Int256{2, 0, 0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int256(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt256_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int256
		i    uint
		want Int256
	}{
		{Int256{0, 0, 0, 0}, 0, Int256{0, 0, 0, 0}},
		{Int256{1, 0, 0, 0}, 1, Int256{0, 0x80000000_00000000, 0, 0}},

		// sign extension
		{Int256{0x80000000_00000000, 0, 0, 0}, 1, Int256{0xc0000000_00000000, 0, 0, 0}},
		{Int256{0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff}, 1, Int256{0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff}},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int256(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt256_Sign(t *testing.T) {
	testCases := []struct {
		x    Int256
		want int
	}{
		{Int256{0, 0, 0, 0}, 0},
		{Int256{1, 0, 0, 0}, 1},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int256(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt256_Neg(t *testing.T) {
	testCases := []struct {
		x    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{0, 0, 0, 1}, Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}},
		{Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Int256{0, 0, 0, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int256(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt256_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int256
		want int
	}{
		{
			Int256{0, 0, 0, 0},
			Int256{0, 0, 0, 0},
			0,
		},
		{
			Int256{0, 0, 0, 1},
			Int256{0, 0, 0, 2},
			-1,
		},
		{
			Int256{0, 0, 0, 1},
			Int256{0, 0, 0, 0},
			1,
		},
		{
			Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
			Int256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64 - 1},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int256(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzInt256_Text(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Int256{u0, u1, u2, u3}
		got := a.Text(base)

		b := int256ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Int256(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzInt256_Append(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Int256{u0, u1, u2, u3}
		buf := a.Append(nil, base)
		got := string(buf)

		b := int256ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Int256(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzInt256_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64) {
		a := Int256{u0, u1, u2, u3}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := int256ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Int256(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzInt256_String(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Add(uint64(1<<63), uint64(0), uint64(0), uint64(0))
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64) {
		a := Int256{u0, u1, u2, u3}
		got := a.String()

		b := int256ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Int256(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkInt256_Text2(b *testing.B) {
	a := Int256{1 << 63, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkInt256_Text10(b *testing.B) {
	a := Int256{1 << 63, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkInt256_Text62(b *testing.B) {
	a := Int256{1 << 63, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkInt256_String(b *testing.B) {
	a := Int256{1 << 63, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestInt256_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int256
		want   string
	}{
		// decimal
		{
			"%d",
			Int256{0, 0, 0, 0},
			"0",
		},
		{
			"%d",
			Int256{0x80000000_00000000, 0, 0, 0},
			"-57896044618658097711785492504343953926634992332820282019728792003956564819968",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
