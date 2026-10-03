package ints

import (
	"math/big"
	"math/rand/v2"
	"runtime"
	"testing"
)

// shiftAmounts returns every shift amount up to bits+64, plus some huge ones.
func shiftAmounts(bits uint) []uint {
	var s []uint
	for i := uint(0); i <= bits+64; i++ {
		s = append(s, i)
	}
	return append(s, 1<<20, 1<<32-1, 1<<63+5, ^uint(0))
}

// wrapBig reduces x modulo 2**bits, interpreted as signed if signed is true.
func wrapBig(x *big.Int, bits uint, signed bool) *big.Int {
	mod := new(big.Int).Lsh(big.NewInt(1), bits)
	x = new(big.Int).Mod(x, mod)
	if signed && x.Bit(int(bits)-1) == 1 {
		x.Sub(x, mod)
	}
	return x
}

func TestShift_CompareBigInt(t *testing.T) {
	r := rand.New(rand.NewPCG(3, 4))
	for range 4 {
		var u512 Uint512
		var u1024 Uint1024
		for j := range u512 {
			u512[j] = r.Uint64()
		}
		for j := range u1024 {
			u1024[j] = r.Uint64()
		}
		// cover both signs
		for _, neg := range []bool{false, true} {
			if neg {
				u512[0] |= 1 << 63
				u1024[0] |= 1 << 63
			} else {
				u512[0] &^= 1 << 63
				u1024[0] &^= 1 << 63
			}

			bu512, bi512 := uint512ToBigInt(u512), int512ToBigInt(Int512(u512))
			for _, i := range shiftAmounts(512) {
				// big.Int.Lsh with a huge shift amount would exhaust memory; the result is zero anyway.
				k := min(i, 1024)
				if got, want := uint512ToBigInt(u512.Lsh(i)), wrapBig(new(big.Int).Lsh(bu512, k), 512, false); got.Cmp(want) != 0 {
					t.Fatalf("Uint512(%x).Lsh(%d) = %x, want %x", bu512, i, got, want)
				}
				if got, want := int512ToBigInt(Int512(u512).Lsh(i)), wrapBig(new(big.Int).Lsh(bi512, k), 512, true); got.Cmp(want) != 0 {
					t.Fatalf("Int512(%x).Lsh(%d) = %x, want %x", bi512, i, got, want)
				}
				if got, want := uint512ToBigInt(u512.Rsh(i)), new(big.Int).Rsh(bu512, k); got.Cmp(want) != 0 {
					t.Fatalf("Uint512(%x).Rsh(%d) = %x, want %x", bu512, i, got, want)
				}
				if got, want := int512ToBigInt(Int512(u512).Rsh(i)), new(big.Int).Rsh(bi512, k); got.Cmp(want) != 0 {
					t.Fatalf("Int512(%x).Rsh(%d) = %x, want %x", bi512, i, got, want)
				}
			}

			bu1024, bi1024 := uint1024ToBigInt(u1024), int1024ToBigInt(Int1024(u1024))
			for _, i := range shiftAmounts(1024) {
				k := min(i, 2048)
				if got, want := uint1024ToBigInt(u1024.Lsh(i)), wrapBig(new(big.Int).Lsh(bu1024, k), 1024, false); got.Cmp(want) != 0 {
					t.Fatalf("Uint1024(%x).Lsh(%d) = %x, want %x", bu1024, i, got, want)
				}
				if got, want := int1024ToBigInt(Int1024(u1024).Lsh(i)), wrapBig(new(big.Int).Lsh(bi1024, k), 1024, true); got.Cmp(want) != 0 {
					t.Fatalf("Int1024(%x).Lsh(%d) = %x, want %x", bi1024, i, got, want)
				}
				if got, want := uint1024ToBigInt(u1024.Rsh(i)), new(big.Int).Rsh(bu1024, k); got.Cmp(want) != 0 {
					t.Fatalf("Uint1024(%x).Rsh(%d) = %x, want %x", bu1024, i, got, want)
				}
				if got, want := int1024ToBigInt(Int1024(u1024).Rsh(i)), new(big.Int).Rsh(bi1024, k); got.Cmp(want) != 0 {
					t.Fatalf("Int1024(%x).Rsh(%d) = %x, want %x", bi1024, i, got, want)
				}
			}
		}
	}
}

var shiftSink uint

func BenchmarkUint512_Lsh(b *testing.B) {
	x := Uint512{}.Not().Rsh(3)
	for b.Loop() {
		runtime.KeepAlive(x.Lsh(shiftSink&255 + 3))
	}
}

func BenchmarkUint512_Rsh(b *testing.B) {
	x := Uint512{}.Not().Rsh(3)
	for b.Loop() {
		runtime.KeepAlive(x.Rsh(shiftSink&255 + 3))
	}
}

func BenchmarkUint1024_Lsh(b *testing.B) {
	x := Uint1024{}.Not().Rsh(3)
	for b.Loop() {
		runtime.KeepAlive(x.Lsh(shiftSink&511 + 3))
	}
}

func BenchmarkUint1024_Rsh(b *testing.B) {
	x := Uint1024{}.Not().Rsh(3)
	for b.Loop() {
		runtime.KeepAlive(x.Rsh(shiftSink&511 + 3))
	}
}

func BenchmarkInt1024_Rsh(b *testing.B) {
	x := Int1024(Uint1024{}.Not().Rsh(3)).Neg()
	for b.Loop() {
		runtime.KeepAlive(x.Rsh(shiftSink&511 + 3))
	}
}
