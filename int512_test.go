package ints

import (
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
