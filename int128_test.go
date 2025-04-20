package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func int128ToBigInt(a Int128) *big.Int {
	var b, c big.Int
	b.SetInt64(int64(a[0]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[1]))
	return &b
}

func FuzzInt128_Add(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(math.MaxUint64), // 1<<64 - 1
		uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
		uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), // MaxInt128
		uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63), uint64(0), // MinInt256
		uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 128-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Int128{u0, u1}
		b := Int128{v0, v1}
		got := int128ToBigInt(a.Add(b))

		ba := int128ToBigInt(a)
		bb := int128ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int128(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt128_Sub(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), // MaxInt128
		uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)
	f.Add(
		uint64(1<<63), uint64(0), // MinInt256
		uint64(0), uint64(1), // 1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 128-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Int128{u0, u1}
		b := Int128{v0, v1}
		got := int128ToBigInt(a.Sub(b))

		ba := int128ToBigInt(a)
		bb := int128ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int128(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt128_Mul(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), // MaxInt128
		uint64(math.MaxUint64), uint64(math.MaxUint64), // -1
	)
	f.Add(
		uint64(1<<63), uint64(0), // MinInt256
		uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(1<<63-1), uint64(math.MaxUint64), // MaxInt128
		uint64(1<<63-1), uint64(math.MaxUint64), // MaxInt128
	)

	base := new(big.Int).Lsh(big.NewInt(1), 128-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Int128{u0, u1}
		b := Int128{v0, v1}
		got := int128ToBigInt(a.Mul(b))

		ba := int128ToBigInt(a)
		bb := int128ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Add(want, base)
		want = want.Mod(want, mod)
		want = want.Sub(want, base)

		if got.Cmp(want) != 0 {
			t.Errorf("Int128(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt128_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(127),
		uint64(0), uint64(10),
	)
	f.Add(
		uint64(127), uint64(0),
		uint64(10), uint64(0),
	)
	f.Add(
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
	)
	f.Add(
		uint64(0x80000000_00000000), uint64(0), // MinInt128
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 128-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Int128{u0, u1}
		b := Int128{v0, v1}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := int128ToBigInt(q)
		gotR := int128ToBigInt(r)

		ba := int128ToBigInt(a)
		bb := int128ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%d).DivMod(%d) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = int128ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint128(%d).Div(%d) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = int128ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%d).Mod(%d) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func FuzzInt128_QuoRem(f *testing.F) {
	f.Add(
		uint64(0), uint64(127),
		uint64(0), uint64(10),
	)
	f.Add(
		uint64(127), uint64(0),
		uint64(10), uint64(0),
	)
	f.Add(
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
	)
	f.Add(
		uint64(0x80000000_00000000), uint64(0), // MinInt128
		uint64(0xffffffff_ffffffff), uint64(0xffffffff_ffffffff), // -1
	)

	base := new(big.Int).Lsh(big.NewInt(1), 128-1)
	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Int128{u0, u1}
		b := Int128{v0, v1}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.QuoRem(b)
		gotQ := int128ToBigInt(q)
		gotR := int128ToBigInt(r)

		ba := int128ToBigInt(a)
		bb := int128ToBigInt(b)
		wantQ, wantR := new(big.Int).QuoRem(ba, bb, new(big.Int))
		wantQ = wantQ.Add(wantQ, base)
		wantQ = wantQ.Mod(wantQ, mod)
		wantQ = wantQ.Sub(wantQ, base)

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%d).QuoRem(%d) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Quo(b)
		gotQ = int128ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint128(%d).Quo(%d) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Rem(b)
		gotR = int128ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%d).Rem(%d) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func TestInt128_And(t *testing.T) {
	testCases := []struct {
		x    Int128
		y    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{1, 0}, Int128{1, 0}},
		{Int128{1, 0}, Int128{2, 0}, Int128{0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Int128(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt128_AndNot(t *testing.T) {
	testCases := []struct {
		x    Int128
		y    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{1, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{2, 0}, Int128{1, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Int128(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt128_Or(t *testing.T) {
	testCases := []struct {
		x    Int128
		y    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{1, 0}, Int128{1, 0}},
		{Int128{1, 0}, Int128{2, 0}, Int128{3, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Int128(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt128_Xor(t *testing.T) {
	testCases := []struct {
		x    Int128
		y    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{1, 0}, Int128{0, 0}},
		{Int128{1, 0}, Int128{2, 0}, Int128{3, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Int128(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestInt128_Not(t *testing.T) {
	testCases := []struct {
		x    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{math.MaxUint64, math.MaxUint64}},
		{Int128{1, 0}, Int128{math.MaxUint64 - 1, math.MaxUint64}},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Int128(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt128_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int128
		i    uint
		want Int128
	}{
		{Int128{0, 0}, 0, Int128{0, 0}},
		{Int128{1, 0}, 1, Int128{2, 0}},
		{Int128{1, 0}, 2, Int128{4, 0}},
		{Int128{1, 0}, 3, Int128{8, 0}},
		{Int128{1, 0}, 4, Int128{16, 0}},
		{Int128{0, 1}, 64, Int128{1, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int128(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt128_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int128
		i    uint
		want Int128
	}{
		{Int128{0, 0}, 0, Int128{0, 0}},
		{Int128{1, 0}, 1, Int128{0, 0x80000000_00000000}},
		{Int128{1, 0}, 2, Int128{0, 0x40000000_00000000}},
		{Int128{8, 0}, 3, Int128{1, 0}},
		{Int128{1, 0}, 64, Int128{0, 1}},

		// Sign extension
		{Int128{0x80000000_00000000, 0}, 1, Int128{0xc0000000_00000000, 0}},
		{Int128{0x80000000_00000000, 0}, 2, Int128{0xe0000000_00000000, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int128(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt128_Sign(t *testing.T) {
	testCases := []struct {
		x    Int128
		want int
	}{
		{Int128{0, 0}, 0},
		{Int128{1, 0}, 1},
		{Int128{math.MaxUint64, math.MaxUint64}, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int128(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt128_Neg(t *testing.T) {
	testCases := []struct {
		x    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}},
		{Int128{0, 1}, Int128{math.MaxUint64, math.MaxUint64}},
		{Int128{math.MaxUint64, math.MaxUint64}, Int128{0, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int128(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt128_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int128
		want int
	}{
		{Int128{0, 0}, Int128{0, 0}, 0},
		{Int128{1, 0}, Int128{0, 1}, 1},
		{Int128{0, 1}, Int128{1, 0}, -1},
		{Int128{math.MaxUint64, math.MaxUint64}, Int128{math.MaxUint64, math.MaxUint64}, 0},
		{Int128{math.MaxUint64, math.MaxUint64}, Int128{math.MaxUint64, math.MaxUint64 - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int128(%s).Cmp(%s) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzInt128_Text(f *testing.F) {
	f.Add(uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, u0, u1 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Int128{u0, u1}
		got := a.Text(base)

		b := int128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Int128(%s).Text(%d) = %q, want %q", want, base, got, want)
		}
	})
}

func FuzzInt128_Append(f *testing.F) {
	f.Add(uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), 62)
	f.Fuzz(func(t *testing.T, u0, u1 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Int128{u0, u1}
		buf := a.Append(nil, base)
		got := string(buf)

		b := int128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Int128(%s).Append(buf, %d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzInt128_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Int128{u0, u1}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := int128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Int128(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzInt128_String(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Add(uint64(1<<63), uint64(0))
	f.Add(uint64(1<<63), uint64(1))

	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Int128{u0, u1}
		got := a.String()

		b := int128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Int128(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkInt128_Text2(b *testing.B) {
	a := Int128{1 << 63, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkInt128_Text10(b *testing.B) {
	a := Int128{1 << 63, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkInt128_Text62(b *testing.B) {
	a := Int128{1 << 63, 0}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkInt128_String(b *testing.B) {
	a := Int128{1 << 63, 0}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestInt128_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Int128
		want   string
	}{
		// decimal
		{
			"%d",
			Int128{0, 0},
			"0",
		},
		{
			"%d",
			Int128{0x80000000_00000000, 0},
			"-170141183460469231731687303715884105728",
		},
	}

	for _, tt := range tests {
		got := fmt.Sprintf(tt.format, tt.value)
		if got != tt.want {
			t.Errorf("%#v: want %q, got %q", tt, tt.want, got)
		}
	}
}
