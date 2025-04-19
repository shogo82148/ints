package ints

import (
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
