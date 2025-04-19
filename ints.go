package ints

import "strconv"

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
