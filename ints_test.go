package ints

import (
	"fmt"
	"math/big"
	"math/rand/v2"
	"testing"
)

// randLimbs fills u with random words, zeroing a random number of the most significant words
// so that every length of the number is exercised.
func randLimbs(r *rand.Rand, u []uint64) {
	zeros := r.IntN(len(u) + 1)
	for i := range u {
		if i < zeros {
			u[i] = 0
		} else {
			u[i] = r.Uint64()
		}
	}
}

func TestText_RandomAllBases(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	check := func(t *testing.T, name string, base int, got string, want *big.Int) {
		t.Helper()
		if w := want.Text(base); got != w {
			t.Fatalf("%s.Text(%d) = %q, want %q", name, base, got, w)
		}
	}
	for range 200 {
		for base := 2; base <= len(digits); base++ {
			var u128 Uint128
			randLimbs(r, u128[:])
			check(t, "Uint128", base, u128.Text(base), uint128ToBigInt(u128))
			check(t, "Int128", base, Int128(u128).Text(base), int128ToBigInt(Int128(u128)))

			var u256 Uint256
			randLimbs(r, u256[:])
			check(t, "Uint256", base, u256.Text(base), uint256ToBigInt(u256))
			check(t, "Int256", base, Int256(u256).Text(base), int256ToBigInt(Int256(u256)))

			var u512 Uint512
			randLimbs(r, u512[:])
			check(t, "Uint512", base, u512.Text(base), uint512ToBigInt(u512))
			check(t, "Int512", base, Int512(u512).Text(base), int512ToBigInt(Int512(u512)))

			var u1024 Uint1024
			randLimbs(r, u1024[:])
			check(t, "Uint1024", base, u1024.Text(base), uint1024ToBigInt(u1024))
			check(t, "Int1024", base, Int1024(u1024).Text(base), int1024ToBigInt(Int1024(u1024)))
		}
	}
}

func TestText_PowersOfTwoAllBases(t *testing.T) {
	one := Uint1024{15: 1}
	for k := uint(0); k < 1024; k++ {
		p := one.Lsh(k)
		for _, v := range []Uint1024{p, p.Sub(one)} {
			want := uint1024ToBigInt(v)
			for base := 2; base <= len(digits); base++ {
				if got, w := v.Text(base), want.Text(base); got != w {
					t.Fatalf("Uint1024(%s).Text(%d) = %q, want %q", want, base, got, w)
				}
				if k < 512 {
					var u Uint512
					copy(u[:], v[8:])
					if got, w := u.Text(base), want.Text(base); got != w {
						t.Fatalf("Uint512(%s).Text(%d) = %q, want %q", want, base, got, w)
					}
				}
				if k < 128 {
					u := Uint128{v[14], v[15]}
					if got, w := u.Text(base), want.Text(base); got != w {
						t.Fatalf("Uint128(%s).Text(%d) = %q, want %q", want, base, got, w)
					}
				}
			}
		}
	}
}

func TestFormat_CompareBigInt(t *testing.T) {
	formats := []string{
		"%v", "%d", "%b", "%o", "%O", "%x", "%X", "%s",
		"%+d", "% d", "%#b", "%#o", "%#x", "%#X", "%+#x",
		"%5d", "%-5d|", "%05d", "%+05d", "%#08x",
		"%100d", "%-100d|", "%0100d", "%+0100x", "%200b", "%0200b",
	}
	values := []Int256{
		{},
		{0, 0, 0, 1},
		{0, 0, 0, 0xabcdef},
		Int256{0, 0, 0, 0xabcdef}.Neg(),
		Int256(Uint256{}.Not().Rsh(1)),
		Int256(Uint256{}.Not().Rsh(1)).Neg(),
	}
	for _, v := range values {
		want := int256ToBigInt(v)
		for _, f := range formats {
			if f == "%#o" && v.IsZero() {
				// math/big prints "00", but the built-in integers print "0" as we do.
				continue
			}
			if got, w := fmt.Sprintf(f, v), fmt.Sprintf(f, want); got != w {
				t.Errorf("Sprintf(%q, Int256(%s)) = %q, want %q", f, want, got, w)
			}
			if v.Sign() >= 0 {
				u := Uint256(v)
				if got, w := fmt.Sprintf(f, u), fmt.Sprintf(f, want); got != w {
					t.Errorf("Sprintf(%q, Uint256(%s)) = %q, want %q", f, want, got, w)
				}
			}
		}
	}
}

// TestDivMod_SmallDivisors checks the reciprocal-based division paths,
// which are used for one- and two-word divisors.
func TestDivMod_SmallDivisors(t *testing.T) {
	r := rand.New(rand.NewPCG(5, 6))
	divisorWord := func() uint64 {
		switch r.IntN(4) {
		case 0:
			return r.Uint64() | 1<<63 // already normalized
		case 1:
			return ^uint64(0)
		case 2:
			return r.Uint64() >> r.UintN(63)
		default:
			return r.Uint64()
		}
	}
	for range 2000 {
		var a Uint1024
		randLimbs(r, a[:])
		var b Uint1024
		b[15] = divisorWord()
		if r.IntN(2) == 0 {
			b[14] = divisorWord()
		}
		if b.IsZero() {
			continue
		}
		ba, bb := uint1024ToBigInt(a), uint1024ToBigInt(b)
		wantQ, wantR := new(big.Int).QuoRem(ba, bb, new(big.Int))

		q, rem := a.DivMod(b)
		if uint1024ToBigInt(q).Cmp(wantQ) != 0 || uint1024ToBigInt(rem).Cmp(wantR) != 0 {
			t.Fatalf("Uint1024(%x).DivMod(%x) = %x, %x, want %x, %x", ba, bb, uint1024ToBigInt(q), uint1024ToBigInt(rem), wantQ, wantR)
		}

		var a512, b512 Uint512
		copy(a512[:], a[8:])
		copy(b512[:], b[8:])
		ba, bb = uint512ToBigInt(a512), uint512ToBigInt(b512)
		wantQ, wantR = new(big.Int).QuoRem(ba, bb, new(big.Int))
		q512, rem512 := a512.DivMod(b512)
		if uint512ToBigInt(q512).Cmp(wantQ) != 0 || uint512ToBigInt(rem512).Cmp(wantR) != 0 {
			t.Fatalf("Uint512(%x).DivMod(%x) = %x, %x, want %x, %x", ba, bb, uint512ToBigInt(q512), uint512ToBigInt(rem512), wantQ, wantR)
		}

		var a256, b256 Uint256
		copy(a256[:], a[12:])
		copy(b256[:], b[12:])
		ba, bb = uint256ToBigInt(a256), uint256ToBigInt(b256)
		wantQ, wantR = new(big.Int).QuoRem(ba, bb, new(big.Int))
		q256, rem256 := a256.DivMod(b256)
		if uint256ToBigInt(q256).Cmp(wantQ) != 0 || uint256ToBigInt(rem256).Cmp(wantR) != 0 {
			t.Fatalf("Uint256(%x).DivMod(%x) = %x, %x, want %x, %x", ba, bb, uint256ToBigInt(q256), uint256ToBigInt(rem256), wantQ, wantR)
		}
	}
}
