package ints

import (
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
