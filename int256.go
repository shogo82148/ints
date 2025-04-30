package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Int256 is a type that represents an 256-bit signed integer.
type Int256 [4]uint64

// IsZero returns true if a is zero.
func (a Int256) IsZero() bool {
	var zero Int256
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Int256) Add(b Int256) Int256 {
	u3, carry := bits.Add64(a[3], b[3], 0)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Int256{u0, u1, u2, u3}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Int256) Sub(b Int256) Int256 {
	u3, borrow := bits.Sub64(a[3], b[3], 0)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Int256{u0, u1, u2, u3}
}

// Mul returns the product a*b.
func (a Int256) Mul(b Int256) Int256 {
	neg := false
	if a.Sign() < 0 {
		a = a.Neg()
		neg = true
	}
	if b.Sign() < 0 {
		b = b.Neg()
		neg = !neg
	}

	c := Int256(Uint256(a).Mul(Uint256(b)))

	if neg {
		return c.Neg()
	}
	return c
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Int256.DivMod] for more details.
func (a Int256) Div(b Int256) Int256 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Int256.DivMod] for more details.
func (a Int256) Mod(b Int256) Int256 {
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
// See [Int256.QuoRem] for T-division and modulus (like Go).
func (a Int256) DivMod(b Int256) (Int256, Int256) {
	q, r := a.QuoRem(b)
	if r.Sign() < 0 {
		if b.Sign() > 0 {
			r = r.Add(b)
			q = q.Sub(Int256{0, 0, 0, 1})
		} else {
			r = r.Sub(b)
			q = q.Add(Int256{0, 0, 0, 1})
		}
	}
	return q, r
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Int256.QuoRem] for more details.
func (a Int256) Quo(b Int256) Int256 {
	q, _ := a.QuoRem(b)
	return q
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Int256.QuoRem] for more details.
func (a Int256) Rem(b Int256) Int256 {
	_, r := a.QuoRem(b)
	return r
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Int256.DivMod] for Euclidean division and modulus (unlike Go).
func (a Int256) QuoRem(b Int256) (Int256, Int256) {
	var negA, negB bool
	if a.Sign() < 0 {
		negA = true
		a = a.Neg()
	}
	if b.Sign() < 0 {
		negB = true
		b = b.Neg()
	}

	q, r := Uint256(a).DivMod(Uint256(b))
	if negA != negB {
		q = q.Neg()
	}
	if negA {
		r = r.Neg()
	}
	return Int256(q), Int256(r)
}

// And returns the bitwise AND of a and b.
func (a Int256) And(b Int256) Int256 {
	return Int256{
		a[0] & b[0],
		a[1] & b[1],
		a[2] & b[2],
		a[3] & b[3],
	}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Int256) AndNot(b Int256) Int256 {
	return Int256{
		a[0] &^ b[0],
		a[1] &^ b[1],
		a[2] &^ b[2],
		a[3] &^ b[3],
	}
}

// Or returns the bitwise OR of a and b.
func (a Int256) Or(b Int256) Int256 {
	return Int256{
		a[0] | b[0],
		a[1] | b[1],
		a[2] | b[2],
		a[3] | b[3],
	}
}

// Xor returns the bitwise XOR of a and b.
func (a Int256) Xor(b Int256) Int256 {
	return Int256{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
	}
}

// Not returns the bitwise NOT of a.
func (a Int256) Not() Int256 {
	return Int256{
		^a[0],
		^a[1],
		^a[2],
		^a[3],
	}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Int256) Lsh(i uint) Int256 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)

	return Int256{
		a[0]<<i | a[1]<<n1 | a[1]>>n2 | a[2]<<n3 | a[2]>>n4 | a[3]<<n5 | a[3]>>n6,
		a[1]<<i | a[2]<<n1 | a[2]>>n2 | a[3]<<n3 | a[3]>>n4,
		a[2]<<i | a[3]<<n1 | a[3]>>n2,
		a[3] << i,
	}
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int256) Rsh(i uint) Int256 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1, v1 := bits.Sub(i, 64, 0)
	n2 := uint(64 - i)
	n3, v3 := bits.Sub(i, 128, 0)
	n4 := uint(128 - i)
	n5, v5 := bits.Sub(i, 192, 0)
	n6 := uint(192 - i)

	mask1 := uint64(int(v1) - 1)
	mask3 := uint64(int(v3) - 1)
	mask5 := uint64(int(v5) - 1)

	return Int256{
		uint64(int64(a[0]) >> i),
		a[1]>>i | mask1&uint64(int64(a[0])>>n1) | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | mask3&uint64(int64(a[0])>>n3) | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | mask5&uint64(int64(a[0])>>n5) | a[0]<<n6,
	}
}

// Sign returns the sign of a.
// It returns 1 if a > 0, -1 if a < 0, and 0 if a == 0.
func (a Int256) Sign() int {
	var zero Int256
	switch {
	case a == zero:
		return 0
	case int64(a[0]) < 0:
		return -1
	default:
		return 1
	}
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Int256) Neg() Int256 {
	u3, borrow := bits.Sub64(0, a[3], 0)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Int256{u0, u1, u2, u3}
}

func (a Int256) Cmp(b Int256) int {
	if ret := cmp.Compare(int64(a[0]), int64(b[0])); ret != 0 {
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
func (a Int256) Text(base int) string {
	_, s := formatBits256(nil, a[0], a[1], a[2], a[3], base, int64(a[0]) < 0, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Int256) Append(dst []byte, base int) []byte {
	d, _ := formatBits256(dst, a[0], a[1], a[2], a[3], base, int64(a[0]) < 0, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Int256) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits256(dst, a[0], a[1], a[2], a[3], 10, int64(a[0]) < 0, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Int256) String() string {
	_, s := formatBits256(nil, a[0], a[1], a[2], a[3], 10, int64(a[0]) < 0, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Int256) Format(s fmt.State, verb rune) {
	sign := a.Sign()
	b := Uint256(a)
	if sign < 0 {
		b = b.Neg()
	}
	format(s, verb, sign, b)
}
