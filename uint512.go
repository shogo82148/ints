package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint512 is a type that represents an 512-bit unsigned integer.
type Uint512 [8]uint64

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Add(b Uint512) Uint512 {
	u7, carry := bits.Add64(a[7], b[7], 0)
	u6, carry := bits.Add64(a[6], b[6], carry)
	u5, carry := bits.Add64(a[5], b[5], carry)
	u4, carry := bits.Add64(a[4], b[4], carry)
	u3, carry := bits.Add64(a[3], b[3], carry)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Sub(b Uint512) Uint512 {
	u7, borrow := bits.Sub64(a[7], b[7], 0)
	u6, borrow := bits.Sub64(a[6], b[6], borrow)
	u5, borrow := bits.Sub64(a[5], b[5], borrow)
	u4, borrow := bits.Sub64(a[4], b[4], borrow)
	u3, borrow := bits.Sub64(a[3], b[3], borrow)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Mul returns the product a*b.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Mul(b Uint512) Uint512 {
	//                                 a0  a1  a2  a3  a4  a5  a6  a7
	//                               x b0  b1  b2  b3  b4  b5  b6  b7
	//                               --------------------------------
	//                                                        h77 l77 - 1.
	//                                                    h67 l67
	//                                                h57 l57
	//                                            h47 l47
	//                                        h37 l37
	//                                    h27 l27
	//                                h17 l17
	//                            h07 l07
	//                                                    h76 l76     - 2.
	//                                                h66 l66
	//                                            h56 l56
	//                                        h46 l46
	//                                    h36 l36
	//                                h26 l26
	//                            h16 l16
	//                        h06 l06
	//                                                h75 l75         - 3.
	//                                            h65 l65
	//                                        h55 l55
	//                                    h45 l45
	//                                h35 l35
	//                            h25 l25
	//                        h15 l15
	//                    h05 l05
	//                                            h74 l74             - 4.
	//                                        h64 l64
	//                                    h54 l54
	//                                h44 l44
	//                            h34 l34
	//                        h24 l24
	//                    h14 l14
	//                h04 l04
	//                                        h73 l73                 - 5.
	//                                    h63 l63
	//                                h53 l53
	//                            h43 l43
	//                        h33 l33
	//                    h23 l23
	//                h13 l13
	//            h03 l03
	//                                    h72 l72                     - 6.
	//                                h62 l62
	//                            h52 l52
	//                        h42 l42
	//                    h32 l32
	//                h22 l22
	//            h12 l12
	//        h02 l02
	//                                h71 l71                         - 7.
	//                            h61 l61
	//                        h51 l51
	//                    h41 l41
	//                h31 l31
	//            h21 l21
	//        h11 l11
	//    h01 l01
	//                            h70 l70                             - 8.
	//                        h60 l60
	//                    h50 l50
	//                h40 l40
	//            h30 l30
	//        h20 l20
	//    h10 l10
	// h00 l00
	// --------------------------------------------------------------
	//                                 u0  u1  u2  u3  u4  u5  u6  u7
	h77, l77 := bits.Mul64(a[7], b[7])
	h67, l67 := bits.Mul64(a[6], b[7])
	h57, l57 := bits.Mul64(a[5], b[7])
	h47, l47 := bits.Mul64(a[4], b[7])
	h37, l37 := bits.Mul64(a[3], b[7])
	h27, l27 := bits.Mul64(a[2], b[7])
	h17, l17 := bits.Mul64(a[1], b[7])
	_, l07 := bits.Mul64(a[0], b[7])

	h76, l76 := bits.Mul64(a[7], b[6])
	h66, l66 := bits.Mul64(a[6], b[6])
	h56, l56 := bits.Mul64(a[5], b[6])
	h46, l46 := bits.Mul64(a[4], b[6])
	h36, l36 := bits.Mul64(a[3], b[6])
	h26, l26 := bits.Mul64(a[2], b[6])
	_, l16 := bits.Mul64(a[1], b[6])
	// h06, l06 := bits.Mul64(a[0], b[6])

	h75, l75 := bits.Mul64(a[7], b[5])
	h65, l65 := bits.Mul64(a[6], b[5])
	h55, l55 := bits.Mul64(a[5], b[5])
	h45, l45 := bits.Mul64(a[4], b[5])
	h35, l35 := bits.Mul64(a[3], b[5])
	_, l25 := bits.Mul64(a[2], b[5])
	// h15, l15 := bits.Mul64(a[1], b[5])
	// h05, l05 := bits.Mul64(a[0], b[5])

	h74, l74 := bits.Mul64(a[7], b[4])
	h64, l64 := bits.Mul64(a[6], b[4])
	h54, l54 := bits.Mul64(a[5], b[4])
	h44, l44 := bits.Mul64(a[4], b[4])
	_, l34 := bits.Mul64(a[3], b[4])
	// h24, l24 := bits.Mul64(a[2], b[4])
	// h14, l14 := bits.Mul64(a[1], b[4])
	// h04, l04 := bits.Mul64(a[0], b[4])

	h73, l73 := bits.Mul64(a[7], b[3])
	h63, l63 := bits.Mul64(a[6], b[3])
	h53, l53 := bits.Mul64(a[5], b[3])
	_, l43 := bits.Mul64(a[4], b[3])
	// h33, l33 := bits.Mul64(a[3], b[3])
	// h23, l23 := bits.Mul64(a[2], b[3])
	// h13, l13 := bits.Mul64(a[1], b[3])
	// h03, l03 := bits.Mul64(a[0], b[3])

	h72, l72 := bits.Mul64(a[7], b[2])
	h62, l62 := bits.Mul64(a[6], b[2])
	_, l52 := bits.Mul64(a[5], b[2])
	// h42, l42 := bits.Mul64(a[4], b[2])
	// h32, l32 := bits.Mul64(a[3], b[2])
	// h22, l22 := bits.Mul64(a[2], b[2])
	// h12, l12 := bits.Mul64(a[1], b[2])
	// h02, l02 := bits.Mul64(a[0], b[2])

	h71, l71 := bits.Mul64(a[7], b[1])
	_, l61 := bits.Mul64(a[6], b[1])
	// h51, l51 := bits.Mul64(a[5], b[1])
	// h41, l41 := bits.Mul64(a[4], b[1])
	// h31, l31 := bits.Mul64(a[3], b[1])
	// h21, l21 := bits.Mul64(a[2], b[1])
	// h11, l11 := bits.Mul64(a[1], b[1])
	// h01, l01 := bits.Mul64(a[0], b[1])

	_, l70 := bits.Mul64(a[7], b[0])
	// h60, l60 := bits.Mul64(a[6], b[0])
	// h50, l50 := bits.Mul64(a[5], b[0])
	// h40, l40 := bits.Mul64(a[4], b[0])
	// h30, l30 := bits.Mul64(a[3], b[0])
	// h20, l20 := bits.Mul64(a[2], b[0])
	// h10, l10 := bits.Mul64(a[1], b[0])
	// h00, l00 := bits.Mul64(a[0], b[0])

	var u0, u1, u2, u3, u4, u5, u6, u7, carry uint64
	// 1.
	u7 = l77
	u6 = l67
	u5 = l57
	u4 = l47
	u3 = l37
	u2 = l27
	u1 = l17
	u0 = l07
	u6, carry = bits.Add64(u6, h77, 0)
	u5, carry = bits.Add64(u5, h67, carry)
	u4, carry = bits.Add64(u4, h57, carry)
	u3, carry = bits.Add64(u3, h47, carry)
	u2, carry = bits.Add64(u2, h37, carry)
	u1, carry = bits.Add64(u1, h27, carry)
	u0, _ = bits.Add64(u0, h17, carry)
	// 2.
	u6, carry = bits.Add64(u6, l76, 0)
	u5, carry = bits.Add64(u5, l66, carry)
	u4, carry = bits.Add64(u4, l56, carry)
	u3, carry = bits.Add64(u3, l46, carry)
	u2, carry = bits.Add64(u2, l36, carry)
	u1, carry = bits.Add64(u1, l26, carry)
	u0, _ = bits.Add64(u0, l16, carry)
	u5, carry = bits.Add64(u5, h76, 0)
	u4, carry = bits.Add64(u4, h66, carry)
	u3, carry = bits.Add64(u3, h56, carry)
	u2, carry = bits.Add64(u2, h46, carry)
	u1, carry = bits.Add64(u1, h36, carry)
	u0, _ = bits.Add64(u0, h26, carry)
	// 3.
	u5, carry = bits.Add64(u5, l75, 0)
	u4, carry = bits.Add64(u4, l65, carry)
	u3, carry = bits.Add64(u3, l55, carry)
	u2, carry = bits.Add64(u2, l45, carry)
	u1, carry = bits.Add64(u1, l35, carry)
	u0, _ = bits.Add64(u0, l25, carry)
	u4, carry = bits.Add64(u4, h75, 0)
	u3, carry = bits.Add64(u3, h65, carry)
	u2, carry = bits.Add64(u2, h55, carry)
	u1, carry = bits.Add64(u1, h45, carry)
	u0, _ = bits.Add64(u0, h35, carry)
	// 4.
	u4, carry = bits.Add64(u4, l74, 0)
	u3, carry = bits.Add64(u3, l64, carry)
	u2, carry = bits.Add64(u2, l54, carry)
	u1, carry = bits.Add64(u1, l44, carry)
	u0, _ = bits.Add64(u0, l34, carry)
	u3, carry = bits.Add64(u3, h74, 0)
	u2, carry = bits.Add64(u2, h64, carry)
	u1, carry = bits.Add64(u1, h54, carry)
	u0, _ = bits.Add64(u0, h44, carry)
	// 5.
	u3, carry = bits.Add64(u3, l73, 0)
	u2, carry = bits.Add64(u2, l63, carry)
	u1, carry = bits.Add64(u1, l53, carry)
	u0, _ = bits.Add64(u0, l43, carry)
	u2, carry = bits.Add64(u2, h73, 0)
	u1, carry = bits.Add64(u1, h63, carry)
	u0, _ = bits.Add64(u0, h53, carry)
	// 6.
	u2, carry = bits.Add64(u2, l72, 0)
	u1, carry = bits.Add64(u1, l62, carry)
	u0, _ = bits.Add64(u0, l52, carry)
	u1, carry = bits.Add64(u1, h72, 0)
	u0, _ = bits.Add64(u0, h62, carry)
	// 7.
	u1, carry = bits.Add64(u1, l71, 0)
	u0, _ = bits.Add64(u0, l61, carry)
	u0, _ = bits.Add64(u0, h71, 0)
	// 8.
	u0, _ = bits.Add64(u0, l70, 0)

	return Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Lsh(i uint) Uint512 {
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

	return Uint512{
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

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Rsh(i uint) Uint512 {
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

	return Uint512{
		a[0] >> i,
		a[1]>>i | a[0]>>n1 | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | a[0]>>n3 | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | a[0]>>n5 | a[0]<<n6,
		a[4]>>i | a[3]>>n1 | a[3]<<n2 | a[2]>>n3 | a[2]<<n4 | a[1]>>n5 | a[1]<<n6 | a[0]>>n7 | a[0]<<n8,
		a[5]>>i | a[4]>>n1 | a[4]<<n2 | a[3]>>n3 | a[3]<<n4 | a[2]>>n5 | a[2]<<n6 | a[1]>>n7 | a[1]<<n8 | a[0]>>n9 | a[0]<<n10,
		a[6]>>i | a[5]>>n1 | a[5]<<n2 | a[4]>>n3 | a[4]<<n4 | a[3]>>n5 | a[3]<<n6 | a[2]>>n7 | a[2]<<n8 | a[1]>>n9 | a[1]<<n10 | a[0]>>n11 | a[0]<<n12,
		a[7]>>i | a[6]>>n1 | a[6]<<n2 | a[5]>>n3 | a[5]<<n4 | a[4]>>n5 | a[4]<<n6 | a[3]>>n7 | a[3]<<n8 | a[2]>>n9 | a[2]<<n10 | a[1]>>n11 | a[1]<<n12 | a[0]>>n13 | a[0]<<n14,
	}
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint512 is unsigned.
func (a Uint512) Sign() int {
	var zero Uint512
	if a == zero {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint512) Neg() Uint512 {
	u7, borrow := bits.Sub64(0, a[7], 0)
	u6, borrow := bits.Sub64(0, a[6], borrow)
	u5, borrow := bits.Sub64(0, a[5], borrow)
	u4, borrow := bits.Sub64(0, a[4], borrow)
	u3, borrow := bits.Sub64(0, a[3], borrow)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Uint512{u0, u1, u2, u3, u4, u5, u6, u7}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint512) Cmp(b Uint512) int {
	if ret := cmp.Compare(a[0], b[0]); ret != 0 {
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
func (a Uint512) Text(base int) string {
	_, s := formatBits512(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], base, false, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint512) Append(dst []byte, base int) []byte {
	d, _ := formatBits512(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], base, false, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint512) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits512(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 10, false, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Uint512) String() string {
	_, s := formatBits512(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 10, false, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Uint512) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
