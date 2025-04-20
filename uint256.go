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
	if b[0] == 0 && b[1] == 0 {
		// optimize for uint256 / uint128
		q0, r0 := Uint128{a[0], a[1]}.DivMod(Uint128{b[2], b[3]})
		q1, r1 := div128(r0, Uint128{a[2], a[3]}, Uint128{b[2], b[3]})
		return Uint256{q0[0], q0[1], q1[0], q1[1]}, Uint256{0, 0, r1[0], r1[1]}
	}

	n := uint(Uint128{b[0], b[1]}.LeadingZeros())
	x := a.Rsh(1)
	y := b.Lsh(n)
	q, _ := div128(Uint128{x[0], x[1]}, Uint128{x[2], x[3]}, Uint128{y[0], y[1]})
	q = q.Rsh(127 - n)
	if q.Sign() > 0 {
		q = q.Sub(Uint128{0, 1})
	}

	u := b.Mul(Uint256{0, 0, q[0], q[1]})
	r := a.Sub(u)
	if r.Cmp(b) >= 0 {
		q = q.Add(Uint128{0, 1})
		r = r.Sub(b)
	}

	return Uint256{0, 0, q[0], q[1]}, r
}

// 128-bit of version of bits.Div64.
// https://github.com/golang/go/blob/c893e1cf821b06aa0602f7944ce52f0eb28fd7b5/src/math/bits/bits.go#L514-L568
func div128(hi, lo, y Uint128) (quo, rem Uint128) {
	if y.IsZero() {
		panic("division by zero")
	}
	if y.Cmp(hi) <= 0 {
		panic("division overflow")
	}

	// If high part is zero, we can directly return the results.
	if hi.IsZero() {
		return lo.DivMod(y)
	}

	s := uint(y.LeadingZeros())
	y = y.Lsh(s)

	two64 := Uint128{1, 0}
	yn1 := Uint128{0, y[0]}
	yn0 := Uint128{0, y[1]}
	un64 := hi.Lsh(s).Or(lo.Rsh(128 - s))
	un10 := lo.Lsh(s)
	un1 := Uint128{0, un10[0]}
	un0 := Uint128{0, un10[1]}
	q1 := un64.Div(yn1)
	rhat := un64.Sub(q1.Mul(yn1))

	for q1.Cmp(two64) >= 0 || q1.Mul(yn0).Cmp(two64.Mul(rhat).Add(un1)) > 0 {
		q1 = q1.Sub(Uint128{0, 1})
		rhat = rhat.Add(yn1)
		if rhat.Cmp(two64) >= 0 {
			break
		}
	}

	un21 := un64.Mul(two64).Add(un1).Sub(q1.Mul(y))
	q0 := un21.Div(yn1)
	rhat = un21.Sub(q0.Mul(yn1))

	for q0.Cmp(two64) >= 0 || q0.Mul(yn0).Cmp(two64.Mul(rhat).Add(un0)) > 0 {
		q0 = q0.Sub(Uint128{0, 1})
		rhat = rhat.Add(yn1)
		if rhat.Cmp(two64) >= 0 {
			break
		}
	}

	return q1.Mul(two64).Add(q0), un21.Mul(two64).Add(un0).Sub(q0.Mul(y)).Rsh(s)
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint256.QuoRem] for more details.
func (a Uint256) Quo(b Uint256) Uint256 {
	return a.Div(b)
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint256.QuoRem] for more details.
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
func (a Uint256) QuoRem(b Uint256) (Uint256, Uint256) {
	return a.DivMod(b)
}

// And returns the bitwise AND of a and b.
func (a Uint256) And(b Uint256) Uint256 {
	return Uint256{a[0] & b[0], a[1] & b[1], a[2] & b[2], a[3] & b[3]}
}

// Or returns the bitwise OR of a and b.
func (a Uint256) Or(b Uint256) Uint256 {
	return Uint256{a[0] | b[0], a[1] | b[1], a[2] | b[2], a[3] | b[3]}
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

// LeadingZeros returns the number of leading zero bits in x; the result is 256 for x == 0.
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
