package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint64 is a type that represents an 64-bit unsigned integer.
// It is an alias for the built-in uint64 type.
type Uint64 uint64

// IsZero returns true if a is zero.
func (a Uint64) IsZero() bool {
	return a == 0
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint64) Add(b Uint64) Uint64 {
	return a + b
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint64) Sub(b Uint64) Uint64 {
	return a - b
}

// Mul returns the product a*b.
func (a Uint64) Mul(b Uint64) Uint64 {
	return a * b
}

// Mul128 returns the product a*b, the result is a 128-bit integer.
func (a Uint64) Mul128(b Uint64) Uint128 {
	h, l := bits.Mul64(uint64(a), uint64(b))
	return Uint128{h, l}
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Uint64.DivMod] for more details.
func (a Uint64) Div(b Uint64) Uint64 {
	return a / b
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Uint64.DivMod] for more details.
func (a Uint64) Mod(b Uint64) Uint64 {
	return a % b
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
// See [Uint64.QuoRem] for T-division and modulus (like Go).
func (a Uint64) DivMod(b Uint64) (Uint64, Uint64) {
	return a / b, a % b
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint64.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Quo simply forwards to Div.
func (a Uint64) Quo(b Uint64) Uint64 {
	return a / b
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint64.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Rem simply forwards to Mod.
func (a Uint64) Rem(b Uint64) Uint64 {
	return a % b
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Uint64.DivMod] for Euclidean division and modulus (unlike Go).
// For unsigned integers T‑division and Euclidean division are identical,
// therefore QuoRem simply forwards to DivMod.
func (a Uint64) QuoRem(b Uint64) (Uint64, Uint64) {
	q, r := bits.Div64(0, uint64(a), uint64(b))
	return Uint64(q), Uint64(r)
}

// And returns the bitwise AND of a and b.
func (a Uint64) And(b Uint64) Uint64 {
	return a & b
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Uint64) AndNot(b Uint64) Uint64 {
	return a &^ b
}

// Or returns the bitwise OR of a and b.
func (a Uint64) Or(b Uint64) Uint64 {
	return a | b
}

// Xor returns the bitwise XOR of a and b.
func (a Uint64) Xor(b Uint64) Uint64 {
	return a ^ b
}

// Not returns the bitwise NOT of a.
func (a Uint64) Not() Uint64 {
	return ^a
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint64) Lsh(i uint) Uint64 {
	return a << i
}

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint64) Rsh(i uint) Uint64 {
	return a >> i
}

// LeadingZeros returns the number of leading zero bits in a; the result is 64 for a == 0.
func (a Uint64) LeadingZeros() int {
	return bits.LeadingZeros64(uint64(a))
}

// TrailingZeros returns the number of trailing zero bits in a; the result is 64 for a == 0.
func (a Uint64) TrailingZeros() int {
	return bits.TrailingZeros64(uint64(a))
}

// BitLen returns the number of bits required to represent a in binary; the result is 0 for a == 0.
func (a Uint64) BitLen() int {
	return bits.Len64(uint64(a))
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint64 is unsigned.
func (a Uint64) Sign() int {
	if a == 0 {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint64) Neg() Uint64 {
	return -a
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint64) Cmp(b Uint64) int {
	return cmp.Compare(a, b)
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Uint64) Text(base int) string {
	return formatUint(uint64(a), base)
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint64) Append(dst []byte, base int) []byte {
	return appendUint(dst, uint64(a), base)
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint64) AppendText(dst []byte) ([]byte, error) {
	return appendUint(dst, uint64(a), 10), nil
}

// String returns the string representation of a in base 10.
func (a Uint64) String() string {
	return formatUint(uint64(a), 10)
}

// Format implements [fmt.Formatter].
func (a Uint64) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
