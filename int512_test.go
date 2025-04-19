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
		runtime.KeepAlive(a.Text(2))
	}
}

func BenchmarkInt512_String(b *testing.B) {
	a := Int512{1 << 63, 0, 0, 0, 0, 0, 0, 0}
	for b.Loop() {
		runtime.KeepAlive(a.String())
	}
}
