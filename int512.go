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
func (a Int512) Mul(b Int512) Int512 {
	neg := false
	if a.Sign() < 0 {
		a = a.Neg()
		neg = true
	}
	if b.Sign() < 0 {
		b = b.Neg()
		neg = !neg
	}

	c := Int512(Uint512(a).Mul(Uint512(b)))

	if neg {
		return c.Neg()
	}
	return c
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
	q, r := a.QuoRem(b)
	if r.Sign() < 0 {
		if b.Sign() > 0 {
			r = r.Add(b)
			q = q.Sub(Int512{0, 0, 0, 0, 0, 0, 0, 1})
		} else {
			r = r.Sub(b)
			q = q.Add(Int512{0, 0, 0, 0, 0, 0, 0, 1})
		}
	}
	return q, r
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
	var negA, negB bool
	if a.Sign() < 0 {
		negA = true
		a = a.Neg()
	}
	if b.Sign() < 0 {
		negB = true
		b = b.Neg()
	}

	q, r := Uint512(a).DivMod(Uint512(b))
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
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)
	n7 := uint(i - 256)
	n8 := uint(256 - i)
	n9 := uint(i - 320)
	n10 := uint(320 - i)
	n11 := uint(i - 384)
	n12 := uint(384 - i)
	n13 := uint(i - 448)
	n14 := uint(448 - i)

	return Int512{
		a[0]<<i | a[1]<<n1 | a[1]>>n2 | a[2]<<n3 | a[2]>>n4 | a[3]<<n5 | a[3]>>n6 | a[4]<<n7 | a[4]>>n8 | a[5]<<n9 | a[5]>>n10 | a[6]<<n11 | a[6]>>n12 | a[7]<<n13 | a[7]>>n14,
		a[1]<<i | a[2]<<n1 | a[2]>>n2 | a[3]<<n3 | a[3]>>n4 | a[4]<<n5 | a[4]>>n6 | a[5]<<n7 | a[5]>>n8 | a[6]<<n9 | a[6]>>n10 | a[7]<<n11 | a[7]>>n12,
		a[2]<<i | a[3]<<n1 | a[3]>>n2 | a[4]<<n3 | a[4]>>n4 | a[5]<<n5 | a[5]>>n6 | a[6]<<n7 | a[6]>>n8 | a[7]<<n9 | a[7]>>n10,
		a[3]<<i | a[4]<<n1 | a[4]>>n2 | a[5]<<n3 | a[5]>>n4 | a[6]<<n5 | a[6]>>n6 | a[7]<<n7 | a[7]>>n8,
		a[4]<<i | a[5]<<n1 | a[5]>>n2 | a[6]<<n3 | a[6]>>n4 | a[7]<<n5 | a[7]>>n6,
		a[5]<<i | a[6]<<n1 | a[6]>>n2 | a[7]<<n3 | a[7]>>n4,
		a[6]<<i | a[7]<<n1 | a[7]>>n2,
		a[7] << i,
	}
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int512) Rsh(i uint) Int512 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1, v1 := bits.Sub(i, 64, 0)
	n2 := uint(64 - i)
	n3, v3 := bits.Sub(i, 128, 0)
	n4 := uint(128 - i)
	n5, v5 := bits.Sub(i, 192, 0)
	n6 := uint(192 - i)
	n7, v7 := bits.Sub(i, 256, 0)
	n8 := uint(256 - i)
	n9, v9 := bits.Sub(i, 320, 0)
	n10 := uint(320 - i)
	n11, v11 := bits.Sub(i, 384, 0)
	n12 := uint(384 - i)
	n13, v13 := bits.Sub(i, 448, 0)
	n14 := uint(448 - i)

	mask1 := uint64(int(v1) - 1)
	mask3 := uint64(int(v3) - 1)
	mask5 := uint64(int(v5) - 1)
	mask7 := uint64(int(v7) - 1)
	mask9 := uint64(int(v9) - 1)
	mask11 := uint64(int(v11) - 1)
	mask13 := uint64(int(v13) - 1)

	return Int512{
		uint64(int64(a[0]) >> i),
		a[1]>>i | mask1&uint64(int64(a[0])>>n1) | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | mask3&uint64(int64(a[0])>>n3) | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | mask5&uint64(int64(a[0])>>n5) | a[0]<<n6,
		a[4]>>i | a[3]>>n1 | a[3]<<n2 | a[2]>>n3 | a[2]<<n4 | a[1]>>n5 | a[1]<<n6 | mask7&uint64(int64(a[0])>>n7) | a[0]<<n8,
		a[5]>>i | a[4]>>n1 | a[4]<<n2 | a[3]>>n3 | a[3]<<n4 | a[2]>>n5 | a[2]<<n6 | a[1]>>n7 | a[1]<<n8 | mask9&uint64(int64(a[0])>>n9) | a[0]<<n10,
		a[6]>>i | a[5]>>n1 | a[5]<<n2 | a[4]>>n3 | a[4]<<n4 | a[3]>>n5 | a[3]<<n6 | a[2]>>n7 | a[2]<<n8 | a[1]>>n9 | a[1]<<n10 | mask11&uint64(int64(a[0])>>n11) | a[0]<<n12,
		a[7]>>i | a[6]>>n1 | a[6]<<n2 | a[5]>>n3 | a[5]<<n4 | a[4]>>n5 | a[4]<<n6 | a[3]>>n7 | a[3]<<n8 | a[2]>>n9 | a[2]<<n10 | a[1]>>n11 | a[1]<<n12 | mask13&uint64(int64(a[0])>>n13) | a[0]<<n14,
	}
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
