package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint8 is a type that represents an 8-bit unsigned integer.
// It is an alias for the built-in uint8 type.
type Uint8 uint8

// IsZero returns true if a is zero.
func (a Uint8) IsZero() bool {
	return a == 0
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint8) Add(b Uint8) Uint8 {
	return a + b
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint8) Sub(b Uint8) Uint8 {
	return a - b
}

// Mul returns the product a*b.
func (a Uint8) Mul(b Uint8) Uint8 {
	return a * b
}

// Mul16 returns the product a*b, the result is a 16-bit integer.
func (a Uint8) Mul16(b Uint8) Uint16 {
	return Uint16(a) * Uint16(b)
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Uint8.DivMod] for more details.
func (a Uint8) Div(b Uint8) Uint8 {
	return a / b
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Uint8.DivMod] for more details.
func (a Uint8) Mod(b Uint8) Uint8 {
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
// See [Uint8.QuoRem] for T-division and modulus (like Go).
func (a Uint8) DivMod(b Uint8) (Uint8, Uint8) {
	return a / b, a % b
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint8.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Quo simply forwards to Div.
func (a Uint8) Quo(b Uint8) Uint8 {
	return a / b
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint8.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Rem simply forwards to Mod.
func (a Uint8) Rem(b Uint8) Uint8 {
	return a % b
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Uint8.DivMod] for Euclidean division and modulus (unlike Go).
// For unsigned integers T‑division and Euclidean division are identical,
// therefore QuoRem simply forwards to DivMod.
func (a Uint8) QuoRem(b Uint8) (Uint8, Uint8) {
	return a / b, a % b
}

// And returns the bitwise AND of a and b.
func (a Uint8) And(b Uint8) Uint8 {
	return a & b
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Uint8) AndNot(b Uint8) Uint8 {
	return a &^ b
}

// Or returns the bitwise OR of a and b.
func (a Uint8) Or(b Uint8) Uint8 {
	return a | b
}

// Xor returns the bitwise XOR of a and b.
func (a Uint8) Xor(b Uint8) Uint8 {
	return a ^ b
}

// Not returns the bitwise NOT of a.
func (a Uint8) Not() Uint8 {
	return ^a
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint8) Lsh(i uint) Uint8 {
	return a << i
}

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint8) Rsh(i uint) Uint8 {
	return a >> i
}

// LeadingZeros returns the number of leading zero bits in a; the result is 8 for a == 0.
func (a Uint8) LeadingZeros() int {
	return bits.LeadingZeros8(uint8(a))
}

// TrailingZeros returns the number of trailing zero bits in a; the result is 8 for a == 0.
func (a Uint8) TrailingZeros() int {
	return bits.TrailingZeros8(uint8(a))
}

// BitLen returns the number of bits required to represent a in binary; the result is 0 for a == 0.
func (a Uint8) BitLen() int {
	return bits.Len8(uint8(a))
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint8 is unsigned.
func (a Uint8) Sign() int {
	if a == 0 {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint8) Neg() Uint8 {
	return -a
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint8) Cmp(b Uint8) int {
	return cmp.Compare(a, b)
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Uint8) Text(base int) string {
	return formatUint(uint64(a), base)
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint8) Append(dst []byte, base int) []byte {
	return appendUint(dst, uint64(a), base)
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint8) AppendText(dst []byte) ([]byte, error) {
	return appendUint(dst, uint64(a), 10), nil
}

// String returns the string representation of a in base 10.
func (a Uint8) String() string {
	return formatUint(uint64(a), 10)
}

// Format implements [fmt.Formatter].
func (a Uint8) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
