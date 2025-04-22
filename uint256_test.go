package ints

import (
	"fmt"
	"math"
	"math/big"
	"runtime"
	"testing"
)

func uint256ToBigInt(a Uint256) *big.Int {
	var b, c big.Int
	b.SetUint64(a[0])
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[1]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[2]))
	b.Lsh(&b, 64)
	b.Add(&b, c.SetUint64(a[3]))
	return &b
}

func FuzzUint256_Add(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxUint256
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		b := Uint256{v0, v1, v2, v3}
		got := uint256ToBigInt(a.Add(b))

		ba := uint256ToBigInt(a)
		bb := uint256ToBigInt(b)
		want := new(big.Int).Add(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint256(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint256_Add(b *testing.B) {
	x := Uint256{0x80000000_00000000, 0, 0, 0}
	y := Uint256{0x7f000000_00000000, 0xffffffff_ffffffff, 0xffffffff_ffffffff, 0xffffffff_ffffffff}

	b.Run("Uint256", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Add(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint256ToBigInt(x)
		yy := uint256ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Add(xx, yy)
		}
	})
}

func FuzzUint256_Sub(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		b := Uint256{v0, v1, v2, v3}
		got := uint256ToBigInt(a.Sub(b))

		ba := uint256ToBigInt(a)
		bb := uint256ToBigInt(b)
		want := new(big.Int).Sub(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint256(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint256_Mul(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
		uint64(0), uint64(0), uint64(0), uint64(0), // 0
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxUint256
		uint64(0), uint64(0), uint64(0), uint64(1), // 1
	)
	f.Add(
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxUint256
		uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), // MaxUint256
	)

	mod := new(big.Int).Lsh(big.NewInt(1), 256)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		b := Uint256{v0, v1, v2, v3}
		got := uint256ToBigInt(a.Mul(b))

		ba := uint256ToBigInt(a)
		bb := uint256ToBigInt(b)
		want := new(big.Int).Mul(ba, bb)
		want = want.Mod(want, mod)

		if got.Cmp(want) != 0 {
			t.Errorf("Uint256(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func BenchmarkUint256_Mul(b *testing.B) {
	const M = 0xffffffff_ffffffff
	x := Uint256{0, 0, M, M}
	y := Uint256{0, 0, M, M}

	b.Run("Uint256", func(b *testing.B) {
		for b.Loop() {
			runtime.KeepAlive(x.Mul(y))
		}
	})

	b.Run("BigInt", func(b *testing.B) {
		xx := uint256ToBigInt(x)
		yy := uint256ToBigInt(y)
		zz := new(big.Int)
		for b.Loop() {
			zz.Mul(xx, yy)
		}
	})
}

func FuzzUint256_DivMod(f *testing.F) {
	f.Add(
		uint64(0), uint64(0), uint64(0), uint64(127),
		uint64(0), uint64(0), uint64(0), uint64(10),
	)
	f.Add(
		uint64(0), uint64(1), uint64(0), uint64(1),
		uint64(0), uint64(0), uint64(0), uint64(10),
	)
	f.Add(
		uint64(127), uint64(0), uint64(0), uint64(0),
		uint64(10), uint64(0), uint64(0), uint64(0),
	)
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3, v0, v1, v2, v3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		b := Uint256{v0, v1, v2, v3}
		if b.IsZero() {
			t.Skip("division by zero")
		}
		q, r := a.DivMod(b)
		gotQ := uint256ToBigInt(q)
		gotR := uint256ToBigInt(r)

		ba := uint256ToBigInt(a)
		bb := uint256ToBigInt(b)
		wantQ, wantR := new(big.Int).DivMod(ba, bb, new(big.Int))

		if gotQ.Cmp(wantQ) != 0 || gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint256(%s).DivMod(%s) = %d, %d, want %d, %d", a, b, gotQ, gotR, wantQ, wantR)
		}

		q = a.Div(b)
		gotQ = uint256ToBigInt(q)
		if gotQ.Cmp(wantQ) != 0 {
			t.Errorf("Uint256(%s).Div(%s) = %d, want %d", a, b, gotQ, wantQ)
		}

		r = a.Mod(b)
		gotR = uint256ToBigInt(r)
		if gotR.Cmp(wantR) != 0 {
			t.Errorf("Uint256(%s).Mod(%s) = %d, want %d", a, b, gotR, wantR)
		}
	})
}

func TestUint256_And(t *testing.T) {
	testCases := []struct {
		x    Uint256
		y    Uint256
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.And(tc.y)
		if got != tc.want {
			t.Errorf("Uint256(%d).And(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint256_AndNot(t *testing.T) {
	testCases := []struct {
		x    Uint256
		y    Uint256
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
			Uint256{0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.AndNot(tc.y)
		if got != tc.want {
			t.Errorf("Uint256(%d).AndNot(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint256_Or(t *testing.T) {
	testCases := []struct {
		x    Uint256
		y    Uint256
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Or(tc.y)
		if got != tc.want {
			t.Errorf("Uint256(%d).Or(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint256_Xor(t *testing.T) {
	testCases := []struct {
		x    Uint256
		y    Uint256
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{1, 1, 1, 1},
			Uint256{0, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Xor(tc.y)
		if got != tc.want {
			t.Errorf("Uint256(%d).Xor(%d) = %d, want %d", tc.x, tc.y, got, tc.want)
		}
	}
}

func TestUint256_Not(t *testing.T) {
	testCases := []struct {
		x    Uint256
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64},
		},
		{
			Uint256{1, 1, 1, 1},
			Uint256{math.MaxUint64 - 1, math.MaxUint64 - 1, math.MaxUint64 - 1, math.MaxUint64 - 1},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Not()
		if got != tc.want {
			t.Errorf("Uint256(%d).Not() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint256_Lsh(t *testing.T) {
	testCases := []struct {
		x    Uint256
		i    uint
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			0,
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			1,
			Uint256{2, 2, 2, 2},
		},
		{
			Uint256{1, 1, 1, 1},
			64,
			Uint256{1, 1, 1, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			65,
			Uint256{2, 2, 2, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			128,
			Uint256{1, 1, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			129,
			Uint256{2, 2, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			192,
			Uint256{1, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			193,
			Uint256{2, 0, 0, 0},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint256(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint256_Rsh(t *testing.T) {
	testCases := []struct {
		x    Uint256
		i    uint
		want Uint256
	}{
		{
			Uint256{0, 0, 0, 0},
			0,
			Uint256{0, 0, 0, 0},
		},
		{
			Uint256{1, 1, 1, 1},
			1,
			Uint256{0, 0x80000000_00000000, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint256{1, 1, 1, 1},
			64,
			Uint256{0, 1, 1, 1},
		},
		{
			Uint256{1, 1, 1, 1},
			65,
			Uint256{0, 0, 0x80000000_00000000, 0x80000000_00000000},
		},
		{
			Uint256{1, 1, 1, 1},
			128,
			Uint256{0, 0, 1, 1},
		},
		{
			Uint256{1, 1, 1, 1},
			129,
			Uint256{0, 0, 0, 0x80000000_00000000},
		},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Uint256(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestUint256_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint256
		want int
	}{
		{Uint256{0, 0, 0, 0}, 0},
		{Uint256{0, 0, 0, 1}, 1},
		{Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint256(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint256_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint256
		want Uint256
	}{
		{Uint256{0, 0, 0, 0}, Uint256{0, 0, 0, 0}},
		{Uint256{0, 0, 0, 1}, Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}},
		{Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Uint256{0, 0, 0, 1}},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint256(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint256_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Uint256
		want int
	}{
		{Uint256{0, 0, 0, 0}, Uint256{0, 0, 0, 0}, 0},
		{Uint256{1, 0, 0, 0}, Uint256{0, 0, 0, 0}, 1},
		{Uint256{0, 0, 0, 1}, Uint256{1, 0, 0, 0}, -1},
		{Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Uint256(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func FuzzUint256_Text(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint256{u0, u1, u2, u3}
		got := a.Text(base)

		b := uint256ToBigInt(a)
		want := b.Text(base)

		if got != want {
			t.Errorf("Uint256(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint256_Append(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 10)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 10)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 10)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 10)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 2)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 2)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 2)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 2)
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0), 62)
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0), 62)
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0), 62)
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), 62)

	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64, base int) {
		if base < 2 || base > 62 {
			t.Skip("base out of range")
		}

		a := Uint256{u0, u1, u2, u3}
		buf := a.Append(nil, base)
		got := string(buf)

		b := uint256ToBigInt(a)
		want := b.Text(base)

		if string(got) != want {
			t.Errorf("Uint256(%s).Text(%d) = %q, want %q", b.String(), base, got, want)
		}
	})
}

func FuzzUint256_AppendText(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		buf, err := a.AppendText(nil)
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)

		b := uint256ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint256(%s).AppendText(buf) = %q, want %q", b.String(), got, want)
		}
	})
}

func FuzzUint256_String(f *testing.F) {
	f.Add(uint64(0), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(0), uint64(0), uint64(1), uint64(0))
	f.Add(uint64(0), uint64(1), uint64(0), uint64(0))
	f.Add(uint64(1), uint64(0), uint64(0), uint64(0))
	f.Add(uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64), uint64(math.MaxUint64))
	f.Fuzz(func(t *testing.T, u0, u1, u2, u3 uint64) {
		a := Uint256{u0, u1, u2, u3}
		got := a.String()

		b := uint256ToBigInt(a)
		want := b.String()

		if string(got) != want {
			t.Errorf("Uint256(%s).String() = %q, want %q", b.String(), got, want)
		}
	})
}

func BenchmarkUint256_Text2(b *testing.B) {
	a := Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkUint256_Text10(b *testing.B) {
	a := Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(10))
	}
}

func BenchmarkUint256_Text62(b *testing.B) {
	a := Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.Text(62))
	}
}

func BenchmarkUint256_String(b *testing.B) {
	a := Uint256{math.MaxUint64, math.MaxUint64, math.MaxUint64, math.MaxUint64}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}

func TestUint256_Format(t *testing.T) {
	tests := []struct {
		format string
		value  Uint256
		want   string
	}{
		// decimal
		{
			"%d",
			Uint256{0, 0, 0, 0},
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
