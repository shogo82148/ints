package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint256 is a type that represents an 256-bit unsigned integer.
type Uint256 [4]uint64

// IsZero returns true if a is zero.
func (a Uint256) IsZero() bool {
	var zero Uint256
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint256) Add(b Uint256) Uint256 {
	u3, carry := bits.Add64(a[3], b[3], 0)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Uint256{u0, u1, u2, u3}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint256) Sub(b Uint256) Uint256 {
	u3, borrow := bits.Sub64(a[3], b[3], 0)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Uint256{u0, u1, u2, u3}
}

// Mul returns the product a*b.
func (a Uint256) Mul(b Uint256) Uint256 {
	//                  a0  a1  a2  a3
	//                x b0  b1  b2  b3
	//                ----------------
	//                         h33 l33 - 1.
	//                     h23 l23
	//                 h13 l13
	//             h03 l03
	//                     h32 l32     - 2.
	//                 h22 l22
	//             h12 l12
	//         h02 l02
	//                 h31 l31         - 3.
	//             h21 l21
	//         h11 l11
	//     h01 l01
	//             h30 l30             - 4.
	//         h20 l20
	//     h10 l10
	// h00 l00
	// -------------------------------
	//                  u0  u1  u2  u3

	h33, l33 := bits.Mul64(a[3], b[3])
	h23, l23 := bits.Mul64(a[2], b[3])
	h13, l13 := bits.Mul64(a[1], b[3])
	_, l03 := bits.Mul64(a[0], b[3])

	h32, l32 := bits.Mul64(a[3], b[2])
	h22, l22 := bits.Mul64(a[2], b[2])
	_, l12 := bits.Mul64(a[1], b[2])
	// h02, l02 := bits.Mul64(a[0], b[2])

	h31, l31 := bits.Mul64(a[3], b[1])
	_, l21 := bits.Mul64(a[2], b[1])
	// h11, l11 := bits.Mul64(a[1], b[1])
	// h01, l01 := bits.Mul64(a[0], b[1])

	_, l30 := bits.Mul64(a[3], b[0])
	// h20, l20 := bits.Mul64(a[2], b[0])
	// h10, l10 := bits.Mul64(a[1], b[0])
	// h00, l00 := bits.Mul64(a[0], b[0])

	var u0, u1, u2, u3, carry uint64
	// 1.
	u3 = l33
	u2 = l23
	u1 = l13
	u0 = l03
	u2, carry = bits.Add64(u2, h33, 0)
	u1, carry = bits.Add64(u1, h23, carry)
	u0, _ = bits.Add64(u0, h13, carry)
	// 2.
	u2, carry = bits.Add64(u2, l32, 0)
	u1, carry = bits.Add64(u1, l22, carry)
	u0, _ = bits.Add64(u0, l12, carry)
	u1, carry = bits.Add64(u1, h32, 0)
	u0, _ = bits.Add64(u0, h22, carry)
	// 3.
	u1, carry = bits.Add64(u1, l31, 0)
	u0, _ = bits.Add64(u0, l21, carry)
	u0, _ = bits.Add64(u0, h31, 0)
	// 4.
	u0, _ = bits.Add64(u0, l30, 0)
	return Uint256{u0, u1, u2, u3}
}

// Mul512 returns the product a*b, the result is a 512-bit integer.
func (a Uint256) Mul512(b Uint256) Uint512 {
	//                  a0  a1  a2  a3
	//                x b0  b1  b2  b3
	//                ----------------
	//                         h33 l33 - 1.
	//                     h23 l23
	//                 h13 l13
	//             h03 l03
	//                     h32 l32     - 2.
	//                 h22 l22
	//             h12 l12
	//         h02 l02
	//                 h31 l31         - 3.
	//             h21 l21
	//         h11 l11
	//     h01 l01
	//             h30 l30             - 4.
	//         h20 l20
	//     h10 l10
	// h00 l00
	// -------------------------------
	//  u0  u1  u2  u3  u4  u5  u6  u7

	h33, l33 := bits.Mul64(a[3], b[3])
	h23, l23 := bits.Mul64(a[2], b[3])
	h13, l13 := bits.Mul64(a[1], b[3])
	h03, l03 := bits.Mul64(a[0], b[3])

	h32, l32 := bits.Mul64(a[3], b[2])
	h22, l22 := bits.Mul64(a[2], b[2])
	h12, l12 := bits.Mul64(a[1], b[2])
	h02, l02 := bits.Mul64(a[0], b[2])

	h31, l31 := bits.Mul64(a[3], b[1])
	h21, l21 := bits.Mul64(a[2], b[1])
	h11, l11 := bits.Mul64(a[1], b[1])
	h01, l01 := bits.Mul64(a[0], b[1])

	h30, l30 := bits.Mul64(a[3], b[0])
	h20, l20 := bits.Mul64(a[2], b[0])
	h10, l10 := bits.Mul64(a[1], b[0])
	h00, l00 := bits.Mul64(a[0], b[0])

	var u0, u1, u2, u3, u4, u5, u6, u7, carry uint64
	// 1.
	u7 = l33
	u6 = l23
	u5 = l13
	u4 = l03
	u6, carry = bits.Add64(u6, h33, 0)
	u5, carry = bits.Add64(u5, h23, carry)
	u4, carry = bits.Add64(u4, h13, carry)
	u3, carry = bits.Add64(u3, h03, carry)
	u2, carry = bits.Add64(u2, 0, carry)
	u1, carry = bits.Add64(u1, 0, carry)
	u0, _ = bits.Add64(u0, 0, carry)

	// 2.
	u6, carry = bits.Add64(u6, l32, 0)
	u5, carry = bits.Add64(u5, l22, carry)
	u4, carry = bits.Add64(u4, l12, carry)
	u3, carry = bits.Add64(u3, l02, carry)
	u2, carry = bits.Add64(u2, 0, carry)
	u1, carry = bits.Add64(u1, 0, carry)
	u0, _ = bits.Add64(u0, 0, carry)
	u5, carry = bits.Add64(u5, h32, 0)
	u4, carry = bits.Add64(u4, h22, carry)
	u3, carry = bits.Add64(u3, h12, carry)
	u2, carry = bits.Add64(u2, h02, carry)
	u1, carry = bits.Add64(u1, 0, carry)
	u0, _ = bits.Add64(u0, 0, carry)

	// 3.
	u5, carry = bits.Add64(u5, l31, 0)
	u4, carry = bits.Add64(u4, l21, carry)
	u3, carry = bits.Add64(u3, l11, carry)
	u2, carry = bits.Add64(u2, l01, carry)
	u1, carry = bits.Add64(u1, 0, carry)
	u0, _ = bits.Add64(u0, 0, carry)
	u4, carry = bits.Add64(u4, h31, 0)
	u3, carry = bits.Add64(u3, h21, carry)
	u2, carry = bits.Add64(u2, h11, carry)
	u1, carry = bits.Add64(u1, h01, carry)
	u0, _ = bits.Add64(u0, 0, carry)

	// 4.
	u4, carry = bits.Add64(u4, l30, 0)
	u3, carry = bits.Add64(u3, l20, carry)
	u2, carry = bits.Add64(u2, l10, carry)
	u1, carry = bits.Add64(u1, l00, carry)
	u0, _ = bits.Add64(u0, 0, carry)
	u3, carry = bits.Add64(u3, h30, 0)
	u2, carry = bits.Add64(u2, h20, carry)
	u1, carry = bits.Add64(u1, h10, carry)
	u0, _ = bits.Add64(u0, h00, carry)

	return Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Uint256.DivMod] for more details.
func (a Uint256) Div(b Uint256) Uint256 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Uint256.DivMod] for more details.
func (a Uint256) Mod(b Uint256) Uint256 {
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
// See [Uint256.QuoRem] for T-division and modulus (like Go).
func (a Uint256) DivMod(b Uint256) (Uint256, Uint256) {
	// Convert to little-endian word slices (word 0 is least significant).
	var u, v [4]uint64
	for i := 0; i < 4; i++ {
		u[i] = a[3-i]
		v[i] = b[3-i]
	}

	// n is the number of significant words in the divisor.
	n := 4
	for n > 0 && v[n-1] == 0 {
		n--
	}
	if n == 0 {
		panic("division by zero")
	}

	// m is the number of significant words in the dividend.
	m := 4
	for m > 0 && u[m-1] == 0 {
		m--
	}

	// If the dividend has fewer significant words than the divisor,
	// the quotient is zero and the remainder is the dividend.
	if m < n {
		return Uint256{}, a
	}

	var q, r [4]uint64

	if n == 1 {
		// Single-word divisor: a simple long division suffices.
		d := v[0]
		var rem uint64
		for j := m - 1; j >= 0; j-- {
			q[j], rem = bits.Div64(rem, u[j], d)
		}
		r[0] = rem
	} else {
		divmnu256(&q, &r, &u, &v, m, n)
	}

	// Convert the results back to big-endian Uint256.
	var quo, rem Uint256
	for i := 0; i < 4; i++ {
		quo[3-i] = q[i]
		rem[3-i] = r[i]
	}
	return quo, rem
}

// divmnu256 divides the m-word dividend u by the n-word divisor v
// (both little-endian, n >= 2, v[n-1] != 0) using Knuth's Algorithm D
// and stores the quotient in q and the remainder in r.
// It mirrors [divmnu1024]; see there for a detailed description.
func divmnu256(q, r, u, v *[4]uint64, m, n int) {
	s := uint(bits.LeadingZeros64(v[n-1]))
	var vn [4]uint64
	for i := n - 1; i > 0; i-- {
		vn[i] = v[i]<<s | v[i-1]>>(64-s)
	}
	vn[0] = v[0] << s

	var un [5]uint64
	un[m] = u[m-1] >> (64 - s)
	for i := m - 1; i > 0; i-- {
		un[i] = u[i]<<s | u[i-1]>>(64-s)
	}
	un[0] = u[0] << s

	vn1 := vn[n-1]
	vn2 := vn[n-2]
	ujn := un[m]
	var qhatv [5]uint64

	for j := m - n; j >= 0; j-- {
		qhat := ^uint64(0)
		if ujn != vn1 {
			var rhat uint64
			qhat, rhat = bits.Div64(ujn, un[j+n-1], vn1)

			x1, x2 := bits.Mul64(qhat, vn2)
			ujn2 := un[j+n-2]
			for greaterThanVW(x1, x2, rhat, ujn2) {
				qhat--
				prevRhat := rhat
				rhat += vn1
				if rhat < prevRhat {
					break
				}
				if vn2 > x2 {
					x1--
				}
				x2 -= vn2
			}
		}

		qhatv[n] = mulAddVWW(qhatv[:n], vn[:n], qhat, 0)
		qhl := n + 1
		if j+qhl > m+1 && qhatv[n] == 0 {
			qhl--
		}
		c := subVV(un[j:j+qhl], un[j:j+qhl], qhatv[:qhl])
		if c != 0 {
			c := addVV(un[j:j+n], un[j:j+n], vn[:n])
			if n < qhl {
				un[j+n] += c
			}
			qhat--
		}

		ujn = un[j+n-1]
		q[j] = qhat
	}

	if s == 0 {
		for i := 0; i < n; i++ {
			r[i] = un[i]
		}
	} else {
		for i := 0; i < n-1; i++ {
			r[i] = un[i]>>s | un[i+1]<<(64-s)
		}
		r[n-1] = un[n-1] >> s
	}
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint256.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Quo simply forwards to Div.
func (a Uint256) Quo(b Uint256) Uint256 {
	return a.Div(b)
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint256.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Rem simply forwards to Mod.
func (a Uint256) Rem(b Uint256) Uint256 {
	return a.Mod(b)
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Uint256.DivMod] for Euclidean division and modulus (unlike Go).
// For unsigned integers T‑division and Euclidean division are identical,
// therefore QuoRem simply forwards to DivMod.
func (a Uint256) QuoRem(b Uint256) (Uint256, Uint256) {
	return a.DivMod(b)
}

// And returns the bitwise AND of a and b.
func (a Uint256) And(b Uint256) Uint256 {
	return Uint256{
		a[0] & b[0],
		a[1] & b[1],
		a[2] & b[2],
		a[3] & b[3],
	}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Uint256) AndNot(b Uint256) Uint256 {
	return Uint256{
		a[0] &^ b[0],
		a[1] &^ b[1],
		a[2] &^ b[2],
		a[3] &^ b[3],
	}
}

// Or returns the bitwise OR of a and b.
func (a Uint256) Or(b Uint256) Uint256 {
	return Uint256{
		a[0] | b[0],
		a[1] | b[1],
		a[2] | b[2],
		a[3] | b[3],
	}
}

// Xor returns the bitwise XOR of a and b.
func (a Uint256) Xor(b Uint256) Uint256 {
	return Uint256{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
	}
}

// Not returns the bitwise NOT of a.
func (a Uint256) Not() Uint256 {
	return Uint256{
		^a[0],
		^a[1],
		^a[2],
		^a[3],
	}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint256) Lsh(i uint) Uint256 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)

	return Uint256{
		a[0]<<i | a[1]<<n1 | a[1]>>n2 | a[2]<<n3 | a[2]>>n4 | a[3]<<n5 | a[3]>>n6,
		a[1]<<i | a[2]<<n1 | a[2]>>n2 | a[3]<<n3 | a[3]>>n4,
		a[2]<<i | a[3]<<n1 | a[3]>>n2,
		a[3] << i,
	}
}

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint256) Rsh(i uint) Uint256 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)

	return Uint256{
		a[0] >> i,
		a[1]>>i | a[0]>>n1 | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | a[0]>>n3 | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | a[0]>>n5 | a[0]<<n6,
	}
}

// LeadingZeros returns the number of leading zero bits in a; the result is 256 for a == 0.
func (a Uint256) LeadingZeros() int {
	if a[0] != 0 {
		return bits.LeadingZeros64(a[0])
	}
	if a[1] != 0 {
		return bits.LeadingZeros64(a[1]) + 64
	}
	if a[2] != 0 {
		return bits.LeadingZeros64(a[2]) + 128
	}
	return bits.LeadingZeros64(a[3]) + 192
}

// TrailingZeros returns the number of trailing zero bits in a; the result is 256 for a == 0.
func (a Uint256) TrailingZeros() int {
	if a[3] != 0 {
		return bits.TrailingZeros64(a[3])
	}
	if a[2] != 0 {
		return bits.TrailingZeros64(a[2]) + 64
	}
	if a[1] != 0 {
		return bits.TrailingZeros64(a[1]) + 128
	}
	return bits.TrailingZeros64(a[0]) + 192
}

// BitLen returns the number of bits required to represent a in binary; the result is 0 for a == 0.
func (a Uint256) BitLen() int {
	if a[0] != 0 {
		return bits.Len64(a[0]) + 192
	}
	if a[1] != 0 {
		return bits.Len64(a[1]) + 128
	}
	if a[2] != 0 {
		return bits.Len64(a[2]) + 64
	}
	return bits.Len64(a[3])
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint256 is unsigned.
func (a Uint256) Sign() int {
	var zero Uint256
	if a == zero {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint256) Neg() Uint256 {
	u3, borrow := bits.Sub64(0, a[3], 0)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Uint256{u0, u1, u2, u3}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint256) Cmp(b Uint256) int {
	if ret := cmp.Compare(a[0], b[0]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[1], b[1]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[2], b[2]); ret != 0 {
		return ret
	}
	return cmp.Compare(a[3], b[3])
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Uint256) Text(base int) string {
	_, s := formatBits256(nil, a[0], a[1], a[2], a[3], base, false, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint256) Append(dst []byte, base int) []byte {
	d, _ := formatBits256(dst, a[0], a[1], a[2], a[3], base, false, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint256) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits256(dst, a[0], a[1], a[2], a[3], 10, false, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Uint256) String() string {
	_, s := formatBits256(nil, a[0], a[1], a[2], a[3], 10, false, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Uint256) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
