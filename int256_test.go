package ints

import (
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
