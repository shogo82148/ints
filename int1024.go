package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Int1024 is a type that represents an 1024-bit signed integer.
type Int1024 [16]uint64

// IsZero returns true if a is zero.
func (a Int1024) IsZero() bool {
	var zero Int1024
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Add(b Int1024) Int1024 {
	u15, carry := bits.Add64(a[15], b[15], 0)
	u14, carry := bits.Add64(a[14], b[14], carry)
	u13, carry := bits.Add64(a[13], b[13], carry)
	u12, carry := bits.Add64(a[12], b[12], carry)
	u11, carry := bits.Add64(a[11], b[11], carry)
	u10, carry := bits.Add64(a[10], b[10], carry)
	u9, carry := bits.Add64(a[9], b[9], carry)
	u8, carry := bits.Add64(a[8], b[8], carry)
	u7, carry := bits.Add64(a[7], b[7], carry)
	u6, carry := bits.Add64(a[6], b[6], carry)
	u5, carry := bits.Add64(a[5], b[5], carry)
	u4, carry := bits.Add64(a[4], b[4], carry)
	u3, carry := bits.Add64(a[3], b[3], carry)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Int1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Sub(b Int1024) Int1024 {
	u15, borrow := bits.Sub64(a[15], b[15], 0)
	u14, borrow := bits.Sub64(a[14], b[14], borrow)
	u13, borrow := bits.Sub64(a[13], b[13], borrow)
	u12, borrow := bits.Sub64(a[12], b[12], borrow)
	u11, borrow := bits.Sub64(a[11], b[11], borrow)
	u10, borrow := bits.Sub64(a[10], b[10], borrow)
	u9, borrow := bits.Sub64(a[9], b[9], borrow)
	u8, borrow := bits.Sub64(a[8], b[8], borrow)
	u7, borrow := bits.Sub64(a[7], b[7], borrow)
	u6, borrow := bits.Sub64(a[6], b[6], borrow)
	u5, borrow := bits.Sub64(a[5], b[5], borrow)
	u4, borrow := bits.Sub64(a[4], b[4], borrow)
	u3, borrow := bits.Sub64(a[3], b[3], borrow)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Int1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Mul returns the product a*b.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Mul(b Int1024) Int1024 {
	// In two's complement, the low 1024 bits of the product are
	// the same for signed and unsigned operands.
	return Int1024(Uint1024(a).Mul(Uint1024(b)))
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Int1024.DivMod] for more details.
func (a Int1024) Div(b Int1024) Int1024 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Int1024.DivMod] for more details.
func (a Int1024) Mod(b Int1024) Int1024 {
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
// See [Int1024.QuoRem] for T-division and modulus (like Go).
func (a Int1024) DivMod(b Int1024) (Int1024, Int1024) {
	negA := int64(a[0]) < 0
	negB := int64(b[0]) < 0
	ua, ub := Uint1024(a), Uint1024(b)
	if negA {
		ua = ua.Neg()
	}
	if negB {
		ub = ub.Neg()
	}

	q, r := ua.DivMod(ub)
	if negA && !r.IsZero() {
		// a = -(ub*q + r) = ub*(-(q+1)) + (ub - r)
		q = q.Add(Uint1024{15: 1})
		r = ub.Sub(r)
	}
	if negA != negB {
		q = q.Neg()
	}
	return Int1024(q), Int1024(r)
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Int1024.QuoRem] for more details.
func (a Int1024) Quo(b Int1024) Int1024 {
	q, _ := a.QuoRem(b)
	return q
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Int1024.QuoRem] for more details.
func (a Int1024) Rem(b Int1024) Int1024 {
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
// See [Int1024.DivMod] for Euclidean division and modulus (unlike Go).
func (a Int1024) QuoRem(b Int1024) (Int1024, Int1024) {
	negA := int64(a[0]) < 0
	negB := int64(b[0]) < 0
	ua, ub := Uint1024(a), Uint1024(b)
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
	return Int1024(q), Int1024(r)
}

// And returns the bitwise AND of a and b.
func (a Int1024) And(b Int1024) Int1024 {
	return Int1024{
		a[0] & b[0],
		a[1] & b[1],
		a[2] & b[2],
		a[3] & b[3],
		a[4] & b[4],
		a[5] & b[5],
		a[6] & b[6],
		a[7] & b[7],
		a[8] & b[8],
		a[9] & b[9],
		a[10] & b[10],
		a[11] & b[11],
		a[12] & b[12],
		a[13] & b[13],
		a[14] & b[14],
		a[15] & b[15],
	}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Int1024) AndNot(b Int1024) Int1024 {
	return Int1024{
		a[0] &^ b[0],
		a[1] &^ b[1],
		a[2] &^ b[2],
		a[3] &^ b[3],
		a[4] &^ b[4],
		a[5] &^ b[5],
		a[6] &^ b[6],
		a[7] &^ b[7],
		a[8] &^ b[8],
		a[9] &^ b[9],
		a[10] &^ b[10],
		a[11] &^ b[11],
		a[12] &^ b[12],
		a[13] &^ b[13],
		a[14] &^ b[14],
		a[15] &^ b[15],
	}
}

// Or returns the bitwise OR of a and b.
func (a Int1024) Or(b Int1024) Int1024 {
	return Int1024{
		a[0] | b[0],
		a[1] | b[1],
		a[2] | b[2],
		a[3] | b[3],
		a[4] | b[4],
		a[5] | b[5],
		a[6] | b[6],
		a[7] | b[7],
		a[8] | b[8],
		a[9] | b[9],
		a[10] | b[10],
		a[11] | b[11],
		a[12] | b[12],
		a[13] | b[13],
		a[14] | b[14],
		a[15] | b[15],
	}
}

// Xor returns the bitwise XOR of a and b.
func (a Int1024) Xor(b Int1024) Int1024 {
	return Int1024{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
		a[4] ^ b[4],
		a[5] ^ b[5],
		a[6] ^ b[6],
		a[7] ^ b[7],
		a[8] ^ b[8],
		a[9] ^ b[9],
		a[10] ^ b[10],
		a[11] ^ b[11],
		a[12] ^ b[12],
		a[13] ^ b[13],
		a[14] ^ b[14],
		a[15] ^ b[15],
	}
}

// Not returns the bitwise NOT of a.
func (a Int1024) Not() Int1024 {
	return Int1024{
		^a[0],
		^a[1],
		^a[2],
		^a[3],
		^a[4],
		^a[5],
		^a[6],
		^a[7],
		^a[8],
		^a[9],
		^a[10],
		^a[11],
		^a[12],
		^a[13],
		^a[14],
		^a[15],
	}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Lsh(i uint) Int1024 {
	return lsh1024(a, i)
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Rsh(i uint) Int1024 {
	sign := uint64(int64(a[0]) >> 63)
	return rsh1024(a, i, sign)
}

// Sign returns the sign of a.
// It returns 1 if a > 0, -1 if a < 0, and 0 if a == 0.
func (a Int1024) Sign() int {
	var zero Int1024
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
func (a Int1024) Neg() Int1024 {
	u15, borrow := bits.Sub64(0, a[15], 0)
	u14, borrow := bits.Sub64(0, a[14], borrow)
	u13, borrow := bits.Sub64(0, a[13], borrow)
	u12, borrow := bits.Sub64(0, a[12], borrow)
	u11, borrow := bits.Sub64(0, a[11], borrow)
	u10, borrow := bits.Sub64(0, a[10], borrow)
	u9, borrow := bits.Sub64(0, a[9], borrow)
	u8, borrow := bits.Sub64(0, a[8], borrow)
	u7, borrow := bits.Sub64(0, a[7], borrow)
	u6, borrow := bits.Sub64(0, a[6], borrow)
	u5, borrow := bits.Sub64(0, a[5], borrow)
	u4, borrow := bits.Sub64(0, a[4], borrow)
	u3, borrow := bits.Sub64(0, a[3], borrow)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Int1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Int1024) Cmp(b Int1024) int {
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
	if ret := cmp.Compare(a[7], b[7]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[8], b[8]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[9], b[9]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[10], b[10]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[11], b[11]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[12], b[12]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[13], b[13]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[14], b[14]); ret != 0 {
		return ret
	}
	return cmp.Compare(a[15], b[15])
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Int1024) Text(base int) string {
	_, s := formatBits1024(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], base, int64(a[0]) < 0, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Int1024) Append(dst []byte, base int) []byte {
	d, _ := formatBits1024(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], base, int64(a[0]) < 0, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Int1024) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits1024(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], 10, int64(a[0]) < 0, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Int1024) String() string {
	_, s := formatBits1024(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], 10, int64(a[0]) < 0, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Int1024) Format(s fmt.State, verb rune) {
	sign := a.Sign()
	b := Uint1024(a)
	if sign < 0 {
		b = b.Neg()
	}
	format(s, verb, sign, b)
}
