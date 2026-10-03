package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Int512 is a type that represents an 512-bit signed integer.
type Int512 [8]uint64

// IsZero returns true if a is zero.
func (a Int512) IsZero() bool {
	var zero Int512
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Add(b Int512) Int512 {
	u7, carry := bits.Add64(a[7], b[7], 0)
	u6, carry := bits.Add64(a[6], b[6], carry)
	u5, carry := bits.Add64(a[5], b[5], carry)
	u4, carry := bits.Add64(a[4], b[4], carry)
	u3, carry := bits.Add64(a[3], b[3], carry)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Int512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Sub(b Int512) Int512 {
	u7, borrow := bits.Sub64(a[7], b[7], 0)
	u6, borrow := bits.Sub64(a[6], b[6], borrow)
	u5, borrow := bits.Sub64(a[5], b[5], borrow)
	u4, borrow := bits.Sub64(a[4], b[4], borrow)
	u3, borrow := bits.Sub64(a[3], b[3], borrow)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Int512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Mul returns the product a*b.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Mul(b Int512) Int512 {
	// In two's complement, the low 512 bits of the product are
	// the same for signed and unsigned operands.
	return Int512(Uint512(a).Mul(Uint512(b)))
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Int512.DivMod] for more details.
func (a Int512) Div(b Int512) Int512 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Int512.DivMod] for more details.
func (a Int512) Mod(b Int512) Int512 {
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
// See [Int512.QuoRem] for T-division and modulus (like Go).
func (a Int512) DivMod(b Int512) (Int512, Int512) {
	negA := int64(a[0]) < 0
	negB := int64(b[0]) < 0
	ua, ub := Uint512(a), Uint512(b)
	if negA {
		ua = ua.Neg()
	}
	if negB {
		ub = ub.Neg()
	}

	q, r := ua.DivMod(ub)
	if negA && !r.IsZero() {
		// a = -(ub*q + r) = ub*(-(q+1)) + (ub - r)
		q = q.Add(Uint512{7: 1})
		r = ub.Sub(r)
	}
	if negA != negB {
		q = q.Neg()
	}
	return Int512(q), Int512(r)
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Int512.QuoRem] for more details.
func (a Int512) Quo(b Int512) Int512 {
	q, _ := a.QuoRem(b)
	return q
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Int512.QuoRem] for more details.
func (a Int512) Rem(b Int512) Int512 {
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
// See [Int512.DivMod] for Euclidean division and modulus (unlike Go).
func (a Int512) QuoRem(b Int512) (Int512, Int512) {
	negA := int64(a[0]) < 0
	negB := int64(b[0]) < 0
	ua, ub := Uint512(a), Uint512(b)
	if negA {
		ua = ua.Neg()
	}
	if negB {
		ub = ub.Neg()
	}

	q, r := ua.DivMod(ub)
	if negA != negB {
		q = q.Neg()
	}
	if negA {
		r = r.Neg()
	}
	return Int512(q), Int512(r)
}

// And returns the bitwise AND of a and b.
func (a Int512) And(b Int512) Int512 {
	return Int512{
		a[0] & b[0],
		a[1] & b[1],
		a[2] & b[2],
		a[3] & b[3],
		a[4] & b[4],
		a[5] & b[5],
		a[6] & b[6],
		a[7] & b[7],
	}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Int512) AndNot(b Int512) Int512 {
	return Int512{
		a[0] &^ b[0],
		a[1] &^ b[1],
		a[2] &^ b[2],
		a[3] &^ b[3],
		a[4] &^ b[4],
		a[5] &^ b[5],
		a[6] &^ b[6],
		a[7] &^ b[7],
	}
}

// Or returns the bitwise OR of a and b.
func (a Int512) Or(b Int512) Int512 {
	return Int512{
		a[0] | b[0],
		a[1] | b[1],
		a[2] | b[2],
		a[3] | b[3],
		a[4] | b[4],
		a[5] | b[5],
		a[6] | b[6],
		a[7] | b[7],
	}
}

// Xor returns the bitwise XOR of a and b.
func (a Int512) Xor(b Int512) Int512 {
	return Int512{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
		a[4] ^ b[4],
		a[5] ^ b[5],
		a[6] ^ b[6],
		a[7] ^ b[7],
	}
}

// Not returns the bitwise NOT of a.
func (a Int512) Not() Int512 {
	return Int512{
		^a[0],
		^a[1],
		^a[2],
		^a[3],
		^a[4],
		^a[5],
		^a[6],
		^a[7],
	}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Lsh(i uint) Int512 {
	return lsh512(a, i)
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Rsh(i uint) Int512 {
	sign := uint64(int64(a[0]) >> 63)
	return rsh512(a, i, sign)
}

// Sign returns the sign of a.
// It returns 1 if a > 0, -1 if a < 0, and 0 if a == 0.
func (a Int512) Sign() int {
	var zero Int512
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
func (a Int512) Neg() Int512 {
	u7, borrow := bits.Sub64(0, a[7], 0)
	u6, borrow := bits.Sub64(0, a[6], borrow)
	u5, borrow := bits.Sub64(0, a[5], borrow)
	u4, borrow := bits.Sub64(0, a[4], borrow)
	u3, borrow := bits.Sub64(0, a[3], borrow)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Int512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Int512) Cmp(b Int512) int {
	if ret := cmp.Compare(int64(a[0]), int64(b[0])); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[1], b[1]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[2], b[2]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[3], b[3]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[4], b[4]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[5], b[5]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[6], b[6]); ret != 0 {
		return ret
	}
	return cmp.Compare(a[7], b[7])
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Int512) Text(base int) string {
	_, s := formatBits512(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], base, int64(a[0]) < 0, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Int512) Append(dst []byte, base int) []byte {
	d, _ := formatBits512(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], base, int64(a[0]) < 0, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Int512) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits512(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 10, int64(a[0]) < 0, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Int512) String() string {
	_, s := formatBits512(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 10, int64(a[0]) < 0, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Int512) Format(s fmt.State, verb rune) {
	sign := a.Sign()
	b := Uint512(a)
	if sign < 0 {
		b = b.Neg()
	}
	format(s, verb, sign, b)
}
