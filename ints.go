package ints

import (
	"math/bits"
	"strconv"
)

func formatInt(i int64, base int) string {
	if base < 36 {
		return strconv.FormatInt(i, base)
	}

	// For bases >= 36, implement custom formatting
	_, s := formatBits(nil, uint64(i), base, i < 0, false)
	return s
}

func appendInt(dst []byte, i int64, base int) []byte {
	if base < 36 {
		return strconv.AppendInt(dst, i, base)
	}

	// For bases >= 36, implement custom formatting
	d, _ := formatBits(dst, uint64(i), base, i < 0, true)
	return d
}

func formatUint(i uint64, base int) string {
	if base < 36 {
		return strconv.FormatUint(i, base)
	}

	// For bases >= 36, implement custom formatting
	_, s := formatBits(nil, i, base, false, false)
	return s
}

func appendUint(dst []byte, i uint64, base int) []byte {
	if base < 36 {
		return strconv.AppendUint(dst, i, base)
	}

	// For bases >= 36, implement custom formatting
	d, _ := formatBits(dst, i, base, false, true)
	return d
}

const digits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// formatBits computes the string representation of u in the given base.
// If neg is set, u is treated as negative int64 value. If append_ is
// set, the string is appended to dst and the resulting byte slice is
// returned as the first result value; otherwise the string is returned
// as the second result value.
func formatBits(dst []byte, u uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 36 <= base && base <= len(digits)

	var a [64 + 1]byte // +1 for sign to 64bit value in base 2
	i := len(a)

	if neg {
		u = -u
	}

	b := uint64(base)
	for u >= b {
		i--
		// Avoid using r = a%b in addition to q = a/b
		// since 64bit division and modulo operations
		// are calculated by runtime functions on 32bit machines.
		q := u / b
		a[i] = digits[uint(u-q*b)]
		u = q
	}
	// u < base
	i--
	a[i] = digits[uint(u)]

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}
	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func formatBits128(dst []byte, u0, u1 uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [128 + 1]byte // +1 for sign to 64bit value in base 2
	i := len(a)

	if neg {
		// u = -u = 0 - u
		var borrow uint64
		u1, borrow = bits.Sub64(0, u1, 0)
		u0, _ = bits.Add64(0, u0, borrow)
	}

	// general case
	b := uint64(base)
	for u0 != 0 {
		i--
		q := u0 / b
		var r uint64
		u1, r = bits.Div64(u0-q*b, u1, b)
		u0 = q
		a[i] = digits[uint(r)]
	}
	for u1 >= b {
		i--
		q := u1 / b
		a[i] = digits[uint(u1-q*b)]
		u1 = q
	}
	// u1 < base
	i--
	a[i] = digits[uint(u1)]

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}
	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func formatBits256(dst []byte, u0, u1, u2, u3 uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [256 + 1]byte // +1 for sign to 64bit value in base 2
	i := len(a)

	if neg {
		// u = -u = 0 - u
		var borrow uint64
		u3, borrow = bits.Sub64(0, u3, 0)
		u2, borrow = bits.Sub64(0, u2, borrow)
		u1, borrow = bits.Sub64(0, u1, borrow)
		u0, _ = bits.Sub64(0, u0, borrow)
	}

	// general case
	b := uint64(base)
	for u0 != 0 {
		i--
		q := u0 / b
		var r uint64
		u1, r = bits.Div64(u0-q*b, u1, b)
		u0 = q
		u2, r = bits.Div64(r, u2, b)
		u3, r = bits.Div64(r, u3, b)
		a[i] = digits[uint(r)]
	}
	for u1 != 0 {
		i--
		q := u1 / b
		var r uint64
		u2, r = bits.Div64(u1-q*b, u2, b)
		u1 = q
		u3, r = bits.Div64(r, u3, b)
		a[i] = digits[uint(r)]
	}
	for u2 != 0 {
		i--
		q := u2 / b
		var r uint64
		u3, r = bits.Div64(u2-q*b, u3, b)
		u2 = q
		a[i] = digits[uint(r)]
	}
	for u3 >= b {
		i--
		q := u3 / b
		a[i] = digits[uint(u3-q*b)]
		u3 = q
	}
	// u3 < base
	i--
	a[i] = digits[uint(u3)]

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}
	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func formatBits512(dst []byte, u0, u1, u2, u3, u4, u5, u6, u7 uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [512 + 1]byte // +1 for sign to 512bit value in base 2
	i := len(a)

	if neg {
		// u = -u = 0 - u
		var borrow uint64
		u7, borrow = bits.Sub64(0, u7, 0)
		u6, borrow = bits.Sub64(0, u6, borrow)
		u5, borrow = bits.Sub64(0, u5, borrow)
		u4, borrow = bits.Sub64(0, u4, borrow)
		u3, borrow = bits.Sub64(0, u3, borrow)
		u2, borrow = bits.Sub64(0, u2, borrow)
		u1, borrow = bits.Sub64(0, u1, borrow)
		u0, _ = bits.Sub64(0, u0, borrow)
	}

	// general case
	b := uint64(base)
	for u0 != 0 {
		i--
		q := u0 / b
		var r uint64
		u1, r = bits.Div64(u0-q*b, u1, b)
		u2, r = bits.Div64(r, u2, b)
		u3, r = bits.Div64(r, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u0 = q
		a[i] = digits[uint(r)]
	}
	for u1 != 0 {
		i--
		q := u1 / b
		var r uint64
		u2, r = bits.Div64(u1-q*b, u2, b)
		u3, r = bits.Div64(r, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u1 = q
		a[i] = digits[uint(r)]
	}
	for u2 != 0 {
		i--
		q := u2 / b
		var r uint64
		u3, r = bits.Div64(u2-q*b, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u2 = q
		a[i] = digits[uint(r)]
	}
	for u3 != 0 {
		i--
		q := u3 / b
		var r uint64
		u4, r = bits.Div64(u3-q*b, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u3 = q
		a[i] = digits[uint(r)]
	}
	for u4 != 0 {
		i--
		q := u4 / b
		var r uint64
		u5, r = bits.Div64(u4-q*b, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u4 = q
		a[i] = digits[uint(r)]
	}
	for u5 != 0 {
		i--
		q := u5 / b
		var r uint64
		u6, r = bits.Div64(u5-q*b, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u5 = q
		a[i] = digits[uint(r)]
	}
	for u6 != 0 {
		i--
		q := u6 / b
		var r uint64
		u7, r = bits.Div64(u6-q*b, u7, b)
		u6 = q
		a[i] = digits[uint(r)]
	}
	for u7 >= b {
		i--
		q := u7 / b
		a[i] = digits[uint(u7-q*b)]
		u7 = q
	}
	// u7 < base
	i--
	a[i] = digits[uint(u7)]

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}
	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}

func formatBits1024(dst []byte, u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 uint64, base int, neg, append_ bool) (d []byte, s string) {
	if base < 2 || base > len(digits) {
		panic("strconv: illegal AppendInt/FormatInt base")
	}
	// 2 <= base && base <= len(digits)

	var a [1024 + 1]byte // +1 for sign to 64bit value in base 2
	i := len(a)

	if neg {
		// u = -u = 0 - u
		var borrow uint64
		u15, borrow = bits.Sub64(0, u15, 0)
		u14, borrow = bits.Sub64(0, u14, borrow)
		u13, borrow = bits.Sub64(0, u13, borrow)
		u12, borrow = bits.Sub64(0, u12, borrow)
		u11, borrow = bits.Sub64(0, u11, borrow)
		u10, borrow = bits.Sub64(0, u10, borrow)
		u9, borrow = bits.Sub64(0, u9, borrow)
		u8, borrow = bits.Sub64(0, u8, borrow)
		u7, borrow = bits.Sub64(0, u7, borrow)
		u6, borrow = bits.Sub64(0, u6, borrow)
		u5, borrow = bits.Sub64(0, u5, borrow)
		u4, borrow = bits.Sub64(0, u4, borrow)
		u3, borrow = bits.Sub64(0, u3, borrow)
		u2, borrow = bits.Sub64(0, u2, borrow)
		u1, borrow = bits.Sub64(0, u1, borrow)
		u0, _ = bits.Sub64(0, u0, borrow)
	}

	// general case
	b := uint64(base)
	for u0 != 0 {
		i--
		q := u0 / b
		var r uint64
		u1, r = bits.Div64(u0-q*b, u1, b)
		u2, r = bits.Div64(r, u2, b)
		u3, r = bits.Div64(r, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u0 = q
		a[i] = digits[uint(r)]
	}
	for u1 != 0 {
		i--
		q := u1 / b
		var r uint64
		u2, r = bits.Div64(u1-q*b, u2, b)
		u3, r = bits.Div64(r, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u1 = q
		a[i] = digits[uint(r)]
	}
	for u2 != 0 {
		i--
		q := u2 / b
		var r uint64
		u3, r = bits.Div64(u2-q*b, u3, b)
		u4, r = bits.Div64(r, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u2 = q
		a[i] = digits[uint(r)]
	}
	for u3 != 0 {
		i--
		q := u3 / b
		var r uint64
		u4, r = bits.Div64(u3-q*b, u4, b)
		u5, r = bits.Div64(r, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u3 = q
		a[i] = digits[uint(r)]
	}
	for u4 != 0 {
		i--
		q := u4 / b
		var r uint64
		u5, r = bits.Div64(u4-q*b, u5, b)
		u6, r = bits.Div64(r, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u4 = q
		a[i] = digits[uint(r)]
	}
	for u5 != 0 {
		i--
		q := u5 / b
		var r uint64
		u6, r = bits.Div64(u5-q*b, u6, b)
		u7, r = bits.Div64(r, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u5 = q
		a[i] = digits[uint(r)]
	}
	for u6 != 0 {
		i--
		q := u6 / b
		var r uint64
		u7, r = bits.Div64(u6-q*b, u7, b)
		u8, r = bits.Div64(r, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u6 = q
		a[i] = digits[uint(r)]
	}
	for u7 != 0 {
		i--
		q := u7 / b
		var r uint64
		u8, r = bits.Div64(u7-q*b, u8, b)
		u9, r = bits.Div64(r, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u7 = q
		a[i] = digits[uint(r)]
	}
	for u8 != 0 {
		i--
		q := u8 / b
		var r uint64
		u9, r = bits.Div64(u8-q*b, u9, b)
		u10, r = bits.Div64(r, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u8 = q
		a[i] = digits[uint(r)]
	}
	for u9 != 0 {
		i--
		q := u9 / b
		var r uint64
		u10, r = bits.Div64(u9-q*b, u10, b)
		u11, r = bits.Div64(r, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u9 = q
		a[i] = digits[uint(r)]
	}
	for u10 != 0 {
		i--
		q := u10 / b
		var r uint64
		u11, r = bits.Div64(u10-q*b, u11, b)
		u12, r = bits.Div64(r, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u10 = q
		a[i] = digits[uint(r)]
	}
	for u11 != 0 {
		i--
		q := u11 / b
		var r uint64
		u12, r = bits.Div64(u11-q*b, u12, b)
		u13, r = bits.Div64(r, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u11 = q
		a[i] = digits[uint(r)]
	}
	for u12 != 0 {
		i--
		q := u12 / b
		var r uint64
		u13, r = bits.Div64(u12-q*b, u13, b)
		u14, r = bits.Div64(r, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u12 = q
		a[i] = digits[uint(r)]
	}
	for u13 != 0 {
		i--
		q := u13 / b
		var r uint64
		u14, r = bits.Div64(u13-q*b, u14, b)
		u15, r = bits.Div64(r, u15, b)
		u13 = q
		a[i] = digits[uint(r)]
	}
	for u14 != 0 {
		i--
		q := u14 / b
		var r uint64
		u15, r = bits.Div64(u14-q*b, u15, b)
		u14 = q
		a[i] = digits[uint(r)]
	}
	for u15 >= b {
		i--
		q := u15 / b
		a[i] = digits[uint(u15-q*b)]
		u15 = q
	}
	// u15 < base
	i--
	a[i] = digits[uint(u15)]

	// add sign, if any
	if neg {
		i--
		a[i] = '-'
	}
	if append_ {
		d = append(dst, a[i:]...)
		return
	}
	s = string(a[i:])
	return
}
