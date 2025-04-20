package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Int128 is a type that represents an 128-bit signed integer.
type Int128 [2]uint64

// IsZero returns true if a is zero.
func (a Int128) IsZero() bool {
	var zero Int128
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Int128) Add(b Int128) Int128 {
	u1, carry := bits.Add64(a[1], b[1], 0)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Int128{u0, u1}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Int128) Sub(b Int128) Int128 {
	u1, borrow := bits.Sub64(a[1], b[1], 0)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Int128{u0, u1}
}

// Mul returns the product a*b.
func (a Int128) Mul(b Int128) Int128 {
	neg := false
	if a.Sign() < 0 {
		neg = !neg
		a = a.Neg()
	}
	if b.Sign() < 0 {
		neg = !neg
		b = b.Neg()
	}

	c := Int128(Uint128(a).Mul(Uint128(b)))

	if neg {
		c = c.Neg()
	}
	return c
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Int128.DivMod] for more details.
func (a Int128) Div(b Int128) Int128 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Int128.DivMod] for more details.
func (a Int128) Mod(b Int128) Int128 {
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
// See [Int128.QuoRem] for T-division and modulus (like Go).
func (a Int128) DivMod(b Int128) (Int128, Int128) {
	q, r := a.QuoRem(b)
	if r.Sign() < 0 {
		if b.Sign() > 0 {
			r = r.Add(b)
			q = q.Sub(Int128{0, 1})
		} else {
			r = r.Sub(b)
			q = q.Add(Int128{0, 1})
		}
	}
	return q, r
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Int128.QuoRem] for more details.
func (a Int128) Quo(b Int128) Int128 {
	q, _ := a.QuoRem(b)
	return q
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Int128.QuoRem] for more details.
func (a Int128) Rem(b Int128) Int128 {
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
// See [Int128.DivMod] for Euclidean division and modulus (unlike Go).
func (a Int128) QuoRem(b Int128) (Int128, Int128) {
	var negA, negB bool
	if a.Sign() < 0 {
		negA = true
		a = a.Neg()
	}
	if b.Sign() < 0 {
		negB = true
		b = b.Neg()
	}

	q, r := Uint128(a).DivMod(Uint128(b))
	if negA != negB {
		q = q.Neg()
	}
	if negA {
		r = r.Neg()
	}
	return Int128(q), Int128(r)
}

// And returns the bitwise AND of a and b.
func (a Int128) And(b Int128) Int128 {
	return Int128{a[0] & b[0], a[1] & b[1]}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Int128) AndNot(b Int128) Int128 {
	return Int128{a[0] &^ b[0], a[1] &^ b[1]}
}

// Or returns the bitwise OR of a and b.
func (a Int128) Or(b Int128) Int128 {
	return Int128{a[0] | b[0], a[1] | b[1]}
}

// Xor returns the bitwise XOR of a and b.
func (a Int128) Xor(b Int128) Int128 {
	return Int128{a[0] ^ b[0], a[1] ^ b[1]}
}

// Not returns the bitwise NOT of a.
func (a Int128) Not() Int128 {
	return Int128{^a[0], ^a[1]}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Int128) Lsh(i uint) Int128 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n := uint(i - 64)
	m := uint(64 - i)

	return Int128{
		a[0]<<i | a[1]<<n | a[1]>>m,
		a[1] << i,
	}
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int128) Rsh(i uint) Int128 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n, v := bits.Sub(i, 64, 0)
	m := uint(64 - i)
	mask := uint64(int(v) - 1)

	return Int128{
		uint64(int64(a[0]) >> i),
		mask&uint64(int64(a[0])>>n) | a[0]<<m | a[1]>>i,
	}
}

// Sign returns the sign of a.
// It returns 1 if a > 0, -1 if a < 0, and 0 if a == 0.
func (a Int128) Sign() int {
	var zero Int128
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
func (a Int128) Neg() Int128 {
	u1, borrow := bits.Sub64(0, a[1], 0)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Int128{u0, u1}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Int128) Cmp(b Int128) int {
	if ret := cmp.Compare(int64(a[0]), int64(b[0])); ret != 0 {
		return ret
	}
	sign := 1
	if int64(a[0]) < 0 {
		sign = -1
	}
	return cmp.Compare(a[1], b[1]) * sign
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Int128) Text(base int) string {
	_, s := formatBits128(nil, a[0], a[1], base, int64(a[0]) < 0, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Int128) Append(dst []byte, base int) []byte {
	d, _ := formatBits128(dst, a[0], a[1], base, int64(a[0]) < 0, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Int128) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits128(dst, a[0], a[1], 10, int64(a[0]) < 0, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Int128) String() string {
	_, s := formatBits128(nil, a[0], a[1], 10, int64(a[0]) < 0, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Int128) Format(s fmt.State, verb rune) {
	sign := a.Sign()
	b := Uint128(a)
	if sign < 0 {
		b = b.Neg()
	}
	format(s, verb, sign, b)
}
