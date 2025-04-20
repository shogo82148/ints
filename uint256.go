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

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint128 is unsigned.
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
