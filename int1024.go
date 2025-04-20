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
func (a Int1024) Mul(b Int1024) Int1024 {
	neg := false
	if a.Sign() < 0 {
		a = a.Neg()
		neg = true
	}
	if b.Sign() < 0 {
		b = b.Neg()
		neg = !neg
	}

	c := Int1024(Uint1024(a).Mul(Uint1024(b)))

	if neg {
		return c.Neg()
	}
	return c
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Lsh(i uint) Int1024 {
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
	n15 := uint(i - 512)
	n16 := uint(512 - i)
	n17 := uint(i - 576)
	n18 := uint(576 - i)
	n19 := uint(i - 640)
	n20 := uint(640 - i)
	n21 := uint(i - 704)
	n22 := uint(704 - i)
	n23 := uint(i - 768)
	n24 := uint(768 - i)
	n25 := uint(i - 832)
	n26 := uint(832 - i)
	n27 := uint(i - 896)
	n28 := uint(896 - i)
	n29 := uint(i - 960)
	n30 := uint(960 - i)

	return Int1024{
		a[0]<<i | a[1]<<n1 | a[1]>>n2 | a[2]<<n3 | a[2]>>n4 | a[3]<<n5 | a[3]>>n6 | a[4]<<n7 | a[4]>>n8 |
			a[5]<<n9 | a[5]>>n10 | a[6]<<n11 | a[6]>>n12 | a[7]<<n13 | a[7]>>n14 | a[8]<<n15 | a[8]>>n16 |
			a[9]<<n17 | a[9]>>n18 | a[10]<<n19 | a[10]>>n20 | a[11]<<n21 | a[11]>>n22 | a[12]<<n23 | a[12]>>n24 |
			a[13]<<n25 | a[13]>>n26 | a[14]<<n27 | a[14]>>n28 | a[15]<<n29 | a[15]>>n30,
		a[1]<<i | a[2]<<n1 | a[2]>>n2 | a[3]<<n3 | a[3]>>n4 | a[4]<<n5 | a[4]>>n6 | a[5]<<n7 | a[5]>>n8 |
			a[6]<<n9 | a[6]>>n10 | a[7]<<n11 | a[7]>>n12 | a[8]<<n13 | a[8]>>n14 | a[9]<<n15 | a[9]>>n16 |
			a[10]<<n17 | a[10]>>n18 | a[11]<<n19 | a[11]>>n20 | a[12]<<n21 | a[12]>>n22 | a[13]<<n23 | a[13]>>n24 |
			a[14]<<n25 | a[14]>>n26 | a[15]<<n27 | a[15]>>n28,
		a[2]<<i | a[3]<<n1 | a[3]>>n2 | a[4]<<n3 | a[4]>>n4 | a[5]<<n5 | a[5]>>n6 | a[6]<<n7 | a[6]>>n8 |
			a[7]<<n9 | a[7]>>n10 | a[8]<<n11 | a[8]>>n12 | a[9]<<n13 | a[9]>>n14 | a[10]<<n15 | a[10]>>n16 |
			a[11]<<n17 | a[11]>>n18 | a[12]<<n19 | a[12]>>n20 | a[13]<<n21 | a[13]>>n22 | a[14]<<n23 | a[14]>>n24 |
			a[15]<<n25 | a[15]>>n26,
		a[3]<<i | a[4]<<n1 | a[4]>>n2 | a[5]<<n3 | a[5]>>n4 | a[6]<<n5 | a[6]>>n6 | a[7]<<n7 | a[7]>>n8 |
			a[8]<<n9 | a[8]>>n10 | a[9]<<n11 | a[9]>>n12 | a[10]<<n13 | a[10]>>n14 | a[11]<<n15 | a[11]>>n16 |
			a[12]<<n17 | a[12]>>n18 | a[13]<<n19 | a[13]>>n20 | a[14]<<n21 | a[14]>>n22 | a[15]<<n23 | a[15]>>n24,
		a[4]<<i | a[5]<<n1 | a[5]>>n2 | a[6]<<n3 | a[6]>>n4 | a[7]<<n5 | a[7]>>n6 | a[8]<<n7 | a[8]>>n8 |
			a[9]<<n9 | a[9]>>n10 | a[10]<<n11 | a[10]>>n12 | a[11]<<n13 | a[11]>>n14 | a[12]<<n15 | a[12]>>n16 |
			a[13]<<n17 | a[13]>>n18 | a[14]<<n19 | a[14]>>n20 | a[15]<<n21 | a[15]>>n22,
		a[5]<<i | a[6]<<n1 | a[6]>>n2 | a[7]<<n3 | a[7]>>n4 | a[8]<<n5 | a[8]>>n6 | a[9]<<n7 | a[9]>>n8 |
			a[10]<<n9 | a[10]>>n10 | a[11]<<n11 | a[11]>>n12 | a[12]<<n13 | a[12]>>n14 | a[13]<<n15 | a[13]>>n16 |
			a[14]<<n17 | a[14]>>n18 | a[15]<<n19 | a[15]>>n20,
		a[6]<<i | a[7]<<n1 | a[7]>>n2 | a[8]<<n3 | a[8]>>n4 | a[9]<<n5 | a[9]>>n6 | a[10]<<n7 | a[10]>>n8 |
			a[11]<<n9 | a[11]>>n10 | a[12]<<n11 | a[12]>>n12 | a[13]<<n13 | a[13]>>n14 | a[14]<<n15 | a[14]>>n16 |
			a[15]<<n17 | a[15]>>n18,
		a[7]<<i | a[8]<<n1 | a[8]>>n2 | a[9]<<n3 | a[9]>>n4 | a[10]<<n5 | a[10]>>n6 | a[11]<<n7 | a[11]>>n8 |
			a[12]<<n9 | a[12]>>n10 | a[13]<<n11 | a[13]>>n12 | a[14]<<n13 | a[14]>>n14 | a[15]<<n15 | a[15]>>n16,
		a[8]<<i | a[9]<<n1 | a[9]>>n2 | a[10]<<n3 | a[10]>>n4 | a[11]<<n5 | a[11]>>n6 | a[12]<<n7 | a[12]>>n8 |
			a[13]<<n9 | a[13]>>n10 | a[14]<<n11 | a[14]>>n12 | a[15]<<n13 | a[15]>>n14,
		a[9]<<i | a[10]<<n1 | a[10]>>n2 | a[11]<<n3 | a[11]>>n4 | a[12]<<n5 | a[12]>>n6 | a[13]<<n7 | a[13]>>n8 |
			a[14]<<n9 | a[14]>>n10 | a[15]<<n11 | a[15]>>n12,
		a[10]<<i | a[11]<<n1 | a[11]>>n2 | a[12]<<n3 | a[12]>>n4 | a[13]<<n5 | a[13]>>n6 | a[14]<<n7 | a[14]>>n8 |
			a[15]<<n9 | a[15]>>n10,
		a[11]<<i | a[12]<<n1 | a[12]>>n2 | a[13]<<n3 | a[13]>>n4 | a[14]<<n5 | a[14]>>n6 | a[15]<<n7 | a[15]>>n8,
		a[12]<<i | a[13]<<n1 | a[13]>>n2 | a[14]<<n3 | a[14]>>n4 | a[15]<<n5 | a[15]>>n6,
		a[13]<<i | a[14]<<n1 | a[14]>>n2 | a[15]<<n3 | a[15]>>n4,
		a[14]<<i | a[15]<<n1 | a[15]>>n2,
		a[15] << i,
	}
}

// Rsh returns the arithmetic right shift a>>i, preserving the sign bit.
//
// This function's execution time does not depend on the inputs.
func (a Int1024) Rsh(i uint) Int1024 {
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
	n15, v15 := bits.Sub(i, 512, 0)
	n16 := uint(512 - i)
	n17, v17 := bits.Sub(i, 576, 0)
	n18 := uint(576 - i)
	n19, v19 := bits.Sub(i, 640, 0)
	n20 := uint(640 - i)
	n21, v21 := bits.Sub(i, 704, 0)
	n22 := uint(704 - i)
	n23, v23 := bits.Sub(i, 768, 0)
	n24 := uint(768 - i)
	n25, v25 := bits.Sub(i, 832, 0)
	n26 := uint(832 - i)
	n27, v27 := bits.Sub(i, 896, 0)
	n28 := uint(896 - i)
	n29, v29 := bits.Sub(i, 960, 0)
	n30 := uint(960 - i)

	mask1 := uint64(int(v1) - 1)
	mask3 := uint64(int(v3) - 1)
	mask5 := uint64(int(v5) - 1)
	mask7 := uint64(int(v7) - 1)
	mask9 := uint64(int(v9) - 1)
	mask11 := uint64(int(v11) - 1)
	mask13 := uint64(int(v13) - 1)
	mask15 := uint64(int(v15) - 1)
	mask17 := uint64(int(v17) - 1)
	mask19 := uint64(int(v19) - 1)
	mask21 := uint64(int(v21) - 1)
	mask23 := uint64(int(v23) - 1)
	mask25 := uint64(int(v25) - 1)
	mask27 := uint64(int(v27) - 1)
	mask29 := uint64(int(v29) - 1)

	return Int1024{
		uint64(int64(a[0]) >> i),
		a[1]>>i | mask1&uint64(int64(a[0])>>n1) | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | mask3&uint64(int64(a[0])>>n3) | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | mask5&uint64(int64(a[0])>>n5) | a[0]<<n6,
		a[4]>>i | a[3]>>n1 | a[3]<<n2 | a[2]>>n3 | a[2]<<n4 | a[1]>>n5 | a[1]<<n6 | mask7&uint64(int64(a[0])>>n7) | a[0]<<n8,
		a[5]>>i | a[4]>>n1 | a[4]<<n2 | a[3]>>n3 | a[3]<<n4 | a[2]>>n5 | a[2]<<n6 | a[1]>>n7 | a[1]<<n8 |
			mask9&uint64(int64(a[0])>>n9) | a[0]<<n10,
		a[6]>>i | a[5]>>n1 | a[5]<<n2 | a[4]>>n3 | a[4]<<n4 | a[3]>>n5 | a[3]<<n6 | a[2]>>n7 | a[2]<<n8 |
			a[1]>>n9 | a[1]<<n10 | mask11&uint64(int64(a[0])>>n11) | a[0]<<n12,
		a[7]>>i | a[6]>>n1 | a[6]<<n2 | a[5]>>n3 | a[5]<<n4 | a[4]>>n5 | a[4]<<n6 | a[3]>>n7 | a[3]<<n8 |
			a[2]>>n9 | a[2]<<n10 | a[1]>>n11 | a[1]<<n12 | mask13&uint64(int64(a[0])>>n13) | a[0]<<n14,
		a[8]>>i | a[7]>>n1 | a[7]<<n2 | a[6]>>n3 | a[6]<<n4 | a[5]>>n5 | a[5]<<n6 | a[4]>>n7 | a[4]<<n8 |
			a[3]>>n9 | a[3]<<n10 | a[2]>>n11 | a[2]<<n12 | a[1]>>n13 | a[1]<<n14 | mask15&uint64(int64(a[0])>>n15) | a[0]<<n16,
		a[9]>>i | a[8]>>n1 | a[8]<<n2 | a[7]>>n3 | a[7]<<n4 | a[6]>>n5 | a[6]<<n6 | a[5]>>n7 | a[5]<<n8 |
			a[4]>>n9 | a[4]<<n10 | a[3]>>n11 | a[3]<<n12 | a[2]>>n13 | a[2]<<n14 | a[1]>>n15 | a[1]<<n16 |
			mask17&uint64(int64(a[0])>>n17) | a[0]<<n18,
		a[10]>>i | a[9]>>n1 | a[9]<<n2 | a[8]>>n3 | a[8]<<n4 | a[7]>>n5 | a[7]<<n6 | a[6]>>n7 | a[6]<<n8 |
			a[5]>>n9 | a[5]<<n10 | a[4]>>n11 | a[4]<<n12 | a[3]>>n13 | a[3]<<n14 | a[2]>>n15 | a[2]<<n16 |
			a[1]>>n17 | a[1]<<n18 | mask19&uint64(int64(a[0])>>n19) | a[0]<<n20,
		a[11]>>i | a[10]>>n1 | a[10]<<n2 | a[9]>>n3 | a[9]<<n4 | a[8]>>n5 | a[8]<<n6 | a[7]>>n7 | a[7]<<n8 |
			a[6]>>n9 | a[6]<<n10 | a[5]>>n11 | a[5]<<n12 | a[4]>>n13 | a[4]<<n14 | a[3]>>n15 | a[3]<<n16 |
			a[2]>>n17 | a[2]<<n18 | a[1]>>n19 | a[1]<<n20 | mask21&uint64(int64(a[0])>>n21) | a[0]<<n22,
		a[12]>>i | a[11]>>n1 | a[11]<<n2 | a[10]>>n3 | a[10]<<n4 | a[9]>>n5 | a[9]<<n6 | a[8]>>n7 | a[8]<<n8 |
			a[7]>>n9 | a[7]<<n10 | a[6]>>n11 | a[6]<<n12 | a[5]>>n13 | a[5]<<n14 | a[4]>>n15 | a[4]<<n16 |
			a[3]>>n17 | a[3]<<n18 | a[2]>>n19 | a[2]<<n20 | a[1]>>n21 | a[1]<<n22 | mask23&uint64(int64(a[0])>>n23) | a[0]<<n24,
		a[13]>>i | a[12]>>n1 | a[12]<<n2 | a[11]>>n3 | a[11]<<n4 | a[10]>>n5 | a[10]<<n6 | a[9]>>n7 | a[9]<<n8 |
			a[8]>>n9 | a[8]<<n10 | a[7]>>n11 | a[7]<<n12 | a[6]>>n13 | a[6]<<n14 | a[5]>>n15 | a[5]<<n16 |
			a[4]>>n17 | a[4]<<n18 | a[3]>>n19 | a[3]<<n20 | a[2]>>n21 | a[2]<<n22 | a[1]>>n23 | a[1]<<n24 |
			mask25&uint64(int64(a[0])>>n25) | a[0]<<n26,
		a[14]>>i | a[13]>>n1 | a[13]<<n2 | a[12]>>n3 | a[12]<<n4 | a[11]>>n5 | a[11]<<n6 | a[10]>>n7 | a[10]<<n8 |
			a[9]>>n9 | a[9]<<n10 | a[8]>>n11 | a[8]<<n12 | a[7]>>n13 | a[7]<<n14 | a[6]>>n15 | a[6]<<n16 |
			a[5]>>n17 | a[5]<<n18 | a[4]>>n19 | a[4]<<n20 | a[3]>>n21 | a[3]<<n22 | a[2]>>n23 | a[2]<<n24 |
			a[1]>>n25 | a[1]<<n26 | mask27&uint64(int64(a[0])>>n27) | a[0]<<n28,
		a[15]>>i | a[14]>>n1 | a[14]<<n2 | a[13]>>n3 | a[13]<<n4 | a[12]>>n5 | a[12]<<n6 | a[11]>>n7 | a[11]<<n8 |
			a[10]>>n9 | a[10]<<n10 | a[9]>>n11 | a[9]<<n12 | a[8]>>n13 | a[8]<<n14 | a[7]>>n15 | a[7]<<n16 |
			a[6]>>n17 | a[6]<<n18 | a[5]>>n19 | a[5]<<n20 | a[4]>>n21 | a[4]<<n22 | a[3]>>n23 | a[3]<<n24 |
			a[2]>>n25 | a[2]<<n26 | a[1]>>n27 | a[1]<<n28 | mask29&uint64(int64(a[0])>>n29) | a[0]<<n30,
	}
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
	sign := 1
	if int64(a[0]) < 0 {
		sign = -1
	}
	if ret := cmp.Compare(a[1], b[1]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[2], b[2]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[3], b[3]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[4], b[4]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[5], b[5]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[6], b[6]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[7], b[7]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[8], b[8]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[9], b[9]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[10], b[10]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[11], b[11]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[12], b[12]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[13], b[13]); ret != 0 {
		return ret * sign
	}
	if ret := cmp.Compare(a[14], b[14]); ret != 0 {
		return ret * sign
	}
	return cmp.Compare(a[15], b[15]) * sign
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
