package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func uint128ToBigInt(a Uint128) *big.Int {
	var b, c big.Int
	b.SetUint64(a[0])
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[1]))
	return &b
}

func FuzzUint128_Add(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(math.MaxUint64),
		uint64(0), uint64(1),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(0), uint64(1),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Add(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint128_Add(b *testing.B) {
	x := Uint128{0x80000000_00000000, 0}
	y := Uint128{0x7f000000_00000000, 0xffffffff_ffffffff}

	b.Run("Uint128", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Add(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint128ToBigInt(x)
		yy := uint128ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Add(xx, yy)
		}
	})
}

func FuzzUint128_Sub(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(1),
	)
	f.Add(
		uint64(0), uint64(0),
		uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Sub(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint128_Mul(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0),
		uint64(1), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 128)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint128ToBigInt(a.Mul(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint128_Mul(b *testing.B) {
	x := Uint128{0, 0xffffffff_ffffffff}
	y := Uint128{0, 0xffffffff_ffffffff}

	b.Run("Uint128", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Mul(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint128ToBigInt(x)
		yy := uint128ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Mul(xx, yy)
		}
	})
}

func FuzzUint128_Mul256(f *testing.F) {
	f.Add(
		uint64(0), uint64(0),
		uint64(0), uint64(0),
	)
	f.Add(
		uint64(1), uint64(0),
		uint64(1), uint64(0),
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64),
		uint64(math.MaxUint64), uint64(math.MaxUint64),
	)

	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		got := uint256ToBigInt(a.Mul256(b))

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint128(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint128_Mul256(b *testing.B) {
	x := Uint128{0xffffffff_ffffffff, 0xffffffff_ffffffff}
	y := Uint128{0xffffffff_ffffffff, 0xffffffff_ffffffff}

	b.Run("Uint128", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Mul256(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint128ToBigInt(x)
		yy := uint128ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Mul(xx, yy)
		}
	})
}

func FuzzUint128_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(127),
		uint64(0), uint64(10),
	)
	f.Add(
		uint64(127), uint64(0),
		uint64(10), uint64(0),
	)
	f.Fuzz(func(t *testing.T, u0, u1, v0, v1 uint64) {
		a := Uint128{u0, u1}
		b := Uint128{v0, v1}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := uint128ToBigInt(q)
		gotR := uint128ToBigInt(r)

		ba := uint128ToBigInt(a)
		bb := uint128ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%s).DivMod(%s) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = uint128ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint128(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = uint128ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint128(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func BenchmarkUint128_DivMod(b *testing.B) {
	const M = 0xffffffff_ffffffff
	x := Uint128{M, M}
	y := Uint128{1, M}

	b.Run("Uint128", func(b *testing.B) {
		for b.Loop() {
			x.DivMod(y)
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint128ToBigInt(x)
		yy := uint128ToBigInt(y)
		q := new(big.Int)
		r := new(big.Int)
		for b.Loop() {
			q, r = q.DivMod(xx, yy, r)
		}
	})
}

func TestUint128_And(t *testing.T) {
	testCases := []struct {
		x    Uint128
		y    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{1, 1}, Uint128{1, 1}, Uint128{1, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint128(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint128_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint128
		y    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{1, 1}, Uint128{1, 1}, Uint128{0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint128(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint128_Or(t *testing.T) {
	testCases := []struct {
		x    Uint128
		y    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{1, 1}, Uint128{1, 1}, Uint128{1, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint128(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint128_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint128
		y    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{1, 1}, Uint128{1, 1}, Uint128{0, 0}},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint128(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint128_Not(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{math.MaxUint64, math.MaxUint64}},
		{Uint128{1, 1}, Uint128{math.MaxUint64 - 1, math.MaxUint64 - 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint128(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint128
		i    uint
		want Uint128
	}{
		{Uint128{0, 0}, 0, Uint128{0, 0}},
		{Uint128{1, 1}, 1, Uint128{2, 2}},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint128(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint128_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint128
		i    uint
		want Uint128
	}{
		{Uint128{0, 0}, 0, Uint128{0, 0}},
		{Uint128{1, 1}, 1, Uint128{0, 0x80000000_00000000}},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint128(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint128_LeadingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want int
	}{
		{Uint128{0, 0}, 128},
		{Uint128{1, 0}, 63},
		{Uint128{0, 1}, 127},
	}

	for _, tc := range testCases {
		got := tc.x.LeadingZeros()
		if got != tc.want {
			t.Errorf("Uint128(%d).LeadingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_TrailingZeros(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want int
	}{
		{Uint128{0, 0}, 128},
		{Uint128{1, 0}, 64},
		{Uint128{0, 1}, 0},
	}

	for _, tc := range testCases {
		got := tc.x.TrailingZeros()
		if got != tc.want {
			t.Errorf("Uint128(%d).TrailingZeros() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_BitLen(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want int
	}{
		{Uint128{0, 0}, 0},
		{Uint128{1, 0}, 65},
		{Uint128{0, 1}, 1},
	}

	for _, tc := range testCases {
		got := tc.x.BitLen()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).BitLen() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want int
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{math.MaxUint64, math.MaxUint64}, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint128(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{0, 1}, Uint128{math.MaxUint64, math.MaxUint64}},
		{Uint128{math.MaxUint64, math.MaxUint64}, Uint128{0, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint128(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint128_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint128
		want int
	}{
		{Uint128{0, 0}, Uint128{0, 0}, 0},
		{Uint128{1, 0}, Uint128{0, 1}, 1},
		{Uint128{0, 1}, Uint128{1, 0}, -1},
		{Uint128{math.MaxUint64, math.MaxUint64}, Uint128{math.MaxUint64, math.MaxUint64}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint128(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzUint128_Text(f *testing.F) {
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

		a := Uint128{u0, u1}
		got := a.Text(base)

		b := uint128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_Append(f *testing.F) {
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

		a := Uint128{u0, u1}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint128ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Uint128{u0, u1}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func FuzzUint128_String(f *testing.F) {
	f.Add(uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1 uint64) {
		a := Uint128{u0, u1}
		got := a.String()

		b := uint128ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint128(%s).String() = %q, want %q", want, got, want)
		}
	})
}

func BenchmarkUint128_Text2(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint128_Text10(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkUint128_Text62(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkUint128_String(b *testing.B) {
	a := Uint128{math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestUint128_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint128
		want   string
	}{
		// decimal
		{
			"%d",
			Uint128{0, 0},
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
