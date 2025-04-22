package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint128 is a type that represents an 128-bit unsigned integer.
type Uint128 [2]uint64

// IsZero returns true if a is zero.
func (a Uint128) IsZero() bool {
	var zero Uint128
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Add(b Uint128) Uint128 {
	u1, carry := bits.Add64(a[1], b[1], 0)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Uint128{u0, u1}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Sub(b Uint128) Uint128 {
	u1, borrow := bits.Sub64(a[1], b[1], 0)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Uint128{u0, u1}
}

// Mul returns the product a*b.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Mul(b Uint128) Uint128 {
	h, l := bits.Mul64(a[1], b[1])
	_, h1 := bits.Mul64(a[0], b[1])
	_, h2 := bits.Mul64(a[1], b[0])
	return Uint128{h + h1 + h2, l}
}

// Mul256 returns the product a*b, the result is a 256-bit integer.
func (a Uint128) Mul256(b Uint128) Uint256 {
	//              a0  a1
	//            x b0  b1
	//           ---------
	//             h11 l11 - 1.
	//         h01 l01     - 2.
	//         h10 l10     - 3.
	//     h00 l00         - 4.
	//     ---------------
	//      u0  u1  u2  u3

	h11, l11 := bits.Mul64(a[1], b[1])
	h01, l01 := bits.Mul64(a[0], b[1])
	h10, l10 := bits.Mul64(a[1], b[0])
	h00, l00 := bits.Mul64(a[0], b[0])

	var u0, u1, u2, u3, carry uint64
	// 1.
	u3 = l11
	u2 = h11

	// 2.
	u2, carry = bits.Add64(u2, l01, 0)
	u1, carry = bits.Add64(u1, h01, carry)
	u0, _ = bits.Add64(u0, 0, carry)

	// 3.
	u2, carry = bits.Add64(u2, l10, 0)
	u1, carry = bits.Add64(u1, h10, carry)
	u0, _ = bits.Add64(u0, 0, carry)

	// 4.
	u1, carry = bits.Add64(u1, l00, 0)
	u0, _ = bits.Add64(u0, h00, carry)

	return Uint256{u0, u1, u2, u3}
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Uint128.DivMod] for more details.
func (a Uint128) Div(b Uint128) Uint128 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Uint128.DivMod] for more details.
func (a Uint128) Mod(b Uint128) Uint128 {
	_, r := a.DivMod(b)
	return r
}

// DivMod returns the quotient and remainder of a/b.
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = a div b  such that
//	m = a - b*q  with 0 <= m < |b|
//
// (See Raymond T. Boute, “The Euclidean definition of the functions
// div and mod”. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See [Uint128.QuoRem] for T-division and modulus (like Go).
func (a Uint128) DivMod(b Uint128) (Uint128, Uint128) {
	if b[0] == 0 {
		// optimize for uint128 / uint64
		u0 := a[0] / b[1]
		u1, r := bits.Div64(a[0]-u0*b[1], a[1], b[1])
		return Uint128{u0, u1}, Uint128{0, r}
	}

	n := uint(bits.LeadingZeros64(b[0]))
	x := a.Rsh(1)
	y := b.Lsh(n)
	q, _ := bits.Div64(x[0], x[1], y[0])
	q >>= 63 - n
	if q > 0 {
		q--
	}

	u0, u1 := bits.Mul64(b[1], q)
	u0 += b[0] * q
	r := a.Sub(Uint128{u0, u1})
	if r.Cmp(b) >= 0 {
		q++
		r = r.Sub(b)
	}
	return Uint128{0, q}, r
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint128.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Quo simply forwards to Div.
func (a Uint128) Quo(b Uint128) Uint128 {
	return a.Div(b)
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint128.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Rem simply forwards to Mod.
func (a Uint128) Rem(b Uint128) Uint128 {
	return a.Mod(b)
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Uint128.DivMod] for Euclidean division and modulus (unlike Go).
// For unsigned integers T‑division and Euclidean division are identical,
// therefore QuoRem simply forwards to DivMod.
func (a Uint128) QuoRem(b Uint128) (Uint128, Uint128) {
	return a.DivMod(b)
}

// And returns the bitwise AND of a and b.
func (a Uint128) And(b Uint128) Uint128 {
	return Uint128{a[0] & b[0], a[1] & b[1]}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Uint128) AndNot(b Uint128) Uint128 {
	return Uint128{a[0] &^ b[0], a[1] &^ b[1]}
}

// Or returns the bitwise OR of a and b.
func (a Uint128) Or(b Uint128) Uint128 {
	return Uint128{a[0] | b[0], a[1] | b[1]}
}

// Xor returns the bitwise XOR of a and b.
func (a Uint128) Xor(b Uint128) Uint128 {
	return Uint128{a[0] ^ b[0], a[1] ^ b[1]}
}

// Not returns the bitwise NOT of a.
func (a Uint128) Not() Uint128 {
	return Uint128{^a[0], ^a[1]}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Lsh(i uint) Uint128 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n := uint(i - 64)
	m := uint(64 - i)

	return Uint128{
		a[0]<<i | a[1]<<n | a[1]>>m,
		a[1] << i,
	}
}

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Rsh(i uint) Uint128 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n := uint(i - 64)
	m := uint(64 - i)

	return Uint128{
		a[0] >> i,
		a[0]>>n | a[0]<<m | a[1]>>i,
	}
}

// LeadingZeros returns the number of leading zero bits in x; the result is 128 for x == 0.
func (a Uint128) LeadingZeros() int {
	if a[0] != 0 {
		return bits.LeadingZeros64(a[0])
	}
	return bits.LeadingZeros64(a[1]) + 64
}

// TrailingZeros returns the number of trailing zero bits in x; the result is 128 for x == 0.
func (a Uint128) TrailingZeros() int {
	if a[1] != 0 {
		return bits.TrailingZeros64(a[1])
	}
	return bits.TrailingZeros64(a[0]) + 64
}

// BitLen returns the number of bits required to represent x in binary; the result is 0 for x == 0.
func (a Uint128) BitLen() int {
	if a[0] != 0 {
		return bits.Len64(a[0]) + 64
	}
	return bits.Len64(a[1])
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint128 is unsigned.
func (a Uint128) Sign() int {
	var zero Uint128
	if a == zero {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint128) Neg() Uint128 {
	u1, borrow := bits.Sub64(0, a[1], 0)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Uint128{u0, u1}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint128) Cmp(b Uint128) int {
	if ret := cmp.Compare(a[0], b[0]); ret != 0 {
		return ret
	}
	return cmp.Compare(a[1], b[1])
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Uint128) Text(base int) string {
	_, s := formatBits128(nil, a[0], a[1], base, false, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint128) Append(dst []byte, base int) []byte {
	d, _ := formatBits128(dst, a[0], a[1], base, false, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint128) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits128(dst, a[0], a[1], 10, false, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Uint128) String() string {
	_, s := formatBits128(nil, a[0], a[1], 10, false, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Uint128) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
