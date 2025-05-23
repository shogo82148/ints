package ints

import (
	"bytes"
	"fmt"
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
		u0, _ = bits.Sub64(0, u0, borrow)
	}

	if isPowerOfTwo(base) {
		// Use shifts and masks instead of / and %.
		shift := uint(bits.TrailingZeros(uint(base)))
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u0 != 0 {
			i--
			a[i] = digits[uint(u1)&m]
			u1 = u0<<(64-shift) | u1>>shift
			u0 >>= shift
		}
		for u1 >= b {
			i--
			a[i] = digits[uint(u1)&m]
			u1 >>= shift
		}
		// u1 < base
		i--
		a[i] = digits[uint(u1)]
	} else {
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
	}

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

	if isPowerOfTwo(base) {
		// Use shifts and masks instead of / and %.
		shift := uint(bits.TrailingZeros(uint(base)))
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u0 != 0 {
			i--
			a[i] = digits[uint(u3)&m]
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 = u0<<(64-shift) | u1>>shift
			u0 >>= shift
		}
		for u1 != 0 {
			i--
			a[i] = digits[uint(u3)&m]
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 >>= shift
		}
		for u2 != 0 {
			i--
			a[i] = digits[uint(u3)&m]
			u3 = u2<<(64-shift) | u3>>shift
			u2 >>= shift
		}
		for u3 >= b {
			i--
			a[i] = digits[uint(u3)&m]
			u3 >>= shift
		}
		// u3 < base
		i--
		a[i] = digits[uint(u3)]
	} else {
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
	}

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

	if isPowerOfTwo(base) {
		// Use shifts and masks instead of / and %.
		shift := uint(bits.TrailingZeros(uint(base)))
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u0 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 = u0<<(64-shift) | u1>>shift
			u0 >>= shift
		}
		for u1 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 >>= shift
		}
		for u2 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 >>= shift
		}
		for u3 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 >>= shift
		}
		for u4 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 >>= shift
		}
		for u5 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 >>= shift
		}
		for u6 != 0 {
			i--
			a[i] = digits[uint(u7)&m]
			u7 = u6<<(64-shift) | u7>>shift
			u6 >>= shift
		}
		for u7 >= b {
			i--
			a[i] = digits[uint(u7)&m]
			u7 >>= shift
		}
		// u7 < base
		i--
		a[i] = digits[uint(u7)]
	} else {
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
	}

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

	if isPowerOfTwo(base) {
		// Use shifts and masks instead of / and %.
		shift := uint(bits.TrailingZeros(uint(base)))
		b := uint64(base)
		m := uint(base) - 1 // == 1<<shift - 1
		for u0 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 = u0<<(64-shift) | u1>>shift
			u0 >>= shift
		}
		for u1 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 = u1<<(64-shift) | u2>>shift
			u1 >>= shift
		}
		for u2 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 = u2<<(64-shift) | u3>>shift
			u2 >>= shift
		}
		for u3 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 = u3<<(64-shift) | u4>>shift
			u3 >>= shift
		}
		for u4 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 = u4<<(64-shift) | u5>>shift
			u4 >>= shift
		}
		for u5 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 = u5<<(64-shift) | u6>>shift
			u5 >>= shift
		}
		for u6 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 = u6<<(64-shift) | u7>>shift
			u6 >>= shift
		}
		for u7 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 = u7<<(64-shift) | u8>>shift
			u7 >>= shift
		}
		for u8 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 = u8<<(64-shift) | u9>>shift
			u8 >>= shift
		}
		for u9 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 = u9<<(64-shift) | u10>>shift
			u9 >>= shift
		}
		for u10 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 = u10<<(64-shift) | u11>>shift
			u10 >>= shift
		}
		for u11 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 = u11<<(64-shift) | u12>>shift
			u11 >>= shift
		}
		for u12 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 = u12<<(64-shift) | u13>>shift
			u12 >>= shift
		}
		for u13 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 = u13<<(64-shift) | u14>>shift
			u13 >>= shift
		}
		for u14 != 0 {
			i--
			a[i] = digits[uint(u15)&m]
			u15 = u14<<(64-shift) | u15>>shift
			u14 >>= shift
		}
		for u15 >= b {
			i--
			a[i] = digits[uint(u15)&m]
			u15 >>= shift
		}
		// u15 < base
		i--
		a[i] = digits[uint(u15)]
	} else {
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
	}

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

func isPowerOfTwo(x int) bool {
	return x&(x-1) == 0
}

type appender interface {
	Append(dst []byte, base int) []byte
}

func format(s fmt.State, verb rune, sign int, v appender) {
	var out []byte
	var prefix []byte

	if verb == 'v' {
		out = v.Append(out, 10)
		s.Write(out) //nolint:errcheck
		return
	}

	if s.Flag('+') {
		if sign >= 0 {
			prefix = []byte("+")
		} else {
			prefix = []byte("-")
		}
	} else if s.Flag(' ') {
		if sign >= 0 {
			prefix = []byte(" ")
		} else {
			prefix = []byte("-")
		}
	} else {
		if sign < 0 {
			prefix = []byte("-")
		}
	}

	switch verb {
	case 'b':
		out = v.Append(out, 2)
		if s.Flag('#') {
			prefix = append(prefix, "0b"...)
		}
	case 'o':
		out = v.Append(out, 8)
		if s.Flag('#') && !(len(out) > 0 && out[0] == '0') {
			prefix = append(prefix, '0')
		}
	case 'O':
		out = v.Append(out, 8)
		prefix = append(prefix, "0o"...)
	case 'd':
		out = v.Append(out, 10)
	case 'x':
		out = v.Append(out, 16)
		if s.Flag('#') {
			prefix = append(prefix, "0x"...)
		}
	case 'X':
		out = v.Append(out, 16)
		out = bytes.ToUpper(out)
		if s.Flag('#') {
			prefix = append(prefix, "0X"...)
		}
	case 's':
		out = v.Append(out, 10)
	}

	if w, ok := s.Width(); ok {
		var buf [8]byte
		if s.Flag('0') {
			if len(prefix) > 0 {
				s.Write(prefix) //nolint:errcheck
			}

			// pad with zeros
			buf[0] = '0'
			for i := len(prefix) + len(out); i < w; i++ {
				s.Write(buf[:1]) //nolint:errcheck
			}
			s.Write(out) //nolint:errcheck
		} else if s.Flag('-') {
			if len(prefix) > 0 {
				s.Write(prefix) //nolint:errcheck
			}
			s.Write(out) //nolint:errcheck

			// pad with spaces
			buf[0] = ' '
			for i := len(prefix) + len(out); i < w; i++ {
				s.Write(buf[:1]) //nolint:errcheck
			}
		} else {
			// pad with spaces
			buf[0] = ' '
			for i := len(prefix) + len(out); i < w; i++ {
				s.Write(buf[:1]) //nolint:errcheck
			}
			if len(prefix) > 0 {
				s.Write(prefix) //nolint:errcheck
			}
			s.Write(out) //nolint:errcheck
		}
		return
	}

	if len(prefix) > 0 {
		s.Write(prefix) //nolint:errcheck
	}
	s.Write(out) //nolint:errcheck
}
