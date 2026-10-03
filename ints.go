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
		u := [...]uint64{u0, u1}
		i = formatBitsGeneral(a[:], i, u[:], base)
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
		u := [...]uint64{u0, u1, u2, u3}
		i = formatBitsGeneral(a[:], i, u[:], base)
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
		u := [...]uint64{u0, u1, u2, u3, u4, u5, u6, u7}
		i = formatBitsGeneral(a[:], i, u[:], base)
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
		u := [...]uint64{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
		i = formatBitsGeneral(a[:], i, u[:], base)
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

// bigBase holds, for each base, the largest power of base that fits in a uint64
// and its exponent.
var bigBase = func() (t [len(digits) + 1]struct {
	bb uint64 // base**n
	n  int
}) {
	for base := 2; base <= len(digits); base++ {
		bb, n := uint64(base), 1
		for {
			hi, lo := bits.Mul64(bb, uint64(base))
			if hi != 0 {
				break
			}
			bb, n = lo, n+1
		}
		t[base].bb, t[base].n = bb, n
	}
	return
}()

// formatBitsGeneral writes the digits of the multi-word unsigned integer u
// (most significant word first) in the given base into a, ending at index i,
// and returns the index of the first digit written.
// It destroys the contents of u.
//
// Rather than dividing u by base once per digit, it divides u by bb = base**n,
// the largest power of base that fits in a uint64, and converts each
// single-word remainder into n digits.
func formatBitsGeneral(a []byte, i int, u []uint64, base int) int {
	for len(u) > 1 && u[0] == 0 {
		u = u[1:]
	}
	bb, n := bigBase[base].bb, bigBase[base].n
	for len(u) > 1 {
		// u, r = u / bb, u % bb
		var r uint64
		for j := range u {
			u[j], r = bits.Div64(r, u[j], bb)
		}
		if u[0] == 0 {
			u = u[1:]
		}

		// r has exactly n digits, including leading zeros.
		i = formatChunk(a, i, r, base, n)
	}
	return formatLastChunk(a, i, u[0], base)
}

const smallsString = "00010203040506070809" +
	"10111213141516171819" +
	"20212223242526272829" +
	"30313233343536373839" +
	"40414243444546474849" +
	"50515253545556575859" +
	"60616263646566676869" +
	"70717273747576777879" +
	"80818283848586878889" +
	"90919293949596979899"

// formatChunk writes exactly n digits of r in the given base into a, ending at index i,
// padding with leading zeros.
func formatChunk(a []byte, i int, r uint64, base, n int) int {
	if base == 10 {
		// n == 19: use constant divisors so that the compiler
		// replaces the divisions with multiplications.
		for range 9 {
			q := r / 100
			j := uint(r-q*100) * 2
			r = q
			i -= 2
			a[i+1] = smallsString[j+1]
			a[i] = smallsString[j]
		}
		i--
		a[i] = byte('0' + r)
		return i
	}

	b := uint64(base)
	for range n {
		q := r / b
		i--
		a[i] = digits[uint(r-q*b)]
		r = q
	}
	return i
}

// formatLastChunk writes the digits of u in the given base into a, ending at index i,
// without leading zeros.
func formatLastChunk(a []byte, i int, u uint64, base int) int {
	if base == 10 {
		for u >= 100 {
			q := u / 100
			j := uint(u-q*100) * 2
			u = q
			i -= 2
			a[i+1] = smallsString[j+1]
			a[i] = smallsString[j]
		}
		// u < 100
		j := uint(u) * 2
		i--
		a[i] = smallsString[j+1]
		if u >= 10 {
			i--
			a[i] = smallsString[j]
		}
		return i
	}

	b := uint64(base)
	for u >= b {
		q := u / b
		i--
		a[i] = digits[uint(u-q*b)]
		u = q
	}
	// u < base
	i--
	a[i] = digits[uint(u)]
	return i
}
