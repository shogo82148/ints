package ints

import (
	"fmt"
	"io"
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
		// With only two words, this is faster than formatBitsPow2.
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

	u := [...]uint64{u0, u1, u2, u3}
	if isPowerOfTwo(base) {
		i = formatBitsPow2(a[:], i, u[:], base)
	} else {
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

	u := [...]uint64{u0, u1, u2, u3, u4, u5, u6, u7}
	if isPowerOfTwo(base) {
		i = formatBitsPow2(a[:], i, u[:], base)
	} else {
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

	u := [...]uint64{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
	if isPowerOfTwo(base) {
		i = formatBitsPow2(a[:], i, u[:], base)
	} else {
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

const (
	zeros  = "0000000000000000000000000000000000000000000000000000000000000000"
	spaces = "                                                                "
)

// writePadding writes n copies of the single repeated character in pad to s.
func writePadding(s fmt.State, pad string, n int) {
	for n > 0 {
		m := min(n, len(pad))
		io.WriteString(s, pad[:m]) //nolint:errcheck
		n -= m
	}
}

// format implements [fmt.Formatter] for v, whose absolute value is v and whose sign is sign.
// It is generic so that v is not boxed into an interface.
func format[T appender](s fmt.State, verb rune, sign int, v T) {
	var out []byte

	if verb == 'v' {
		if sign < 0 {
			out = append(out, '-')
		}
		out = v.Append(out, 10)
		s.Write(out) //nolint:errcheck
		return
	}

	var signPrefix string
	if s.Flag('+') {
		if sign >= 0 {
			signPrefix = "+"
		} else {
			signPrefix = "-"
		}
	} else if s.Flag(' ') {
		if sign >= 0 {
			signPrefix = " "
		} else {
			signPrefix = "-"
		}
	} else {
		if sign < 0 {
			signPrefix = "-"
		}
	}

	var basePrefix string
	switch verb {
	case 'b':
		out = v.Append(out, 2)
		if s.Flag('#') {
			basePrefix = "0b"
		}
	case 'o':
		out = v.Append(out, 8)
		if s.Flag('#') && !(len(out) > 0 && out[0] == '0') {
			basePrefix = "0"
		}
	case 'O':
		out = v.Append(out, 8)
		basePrefix = "0o"
	case 'd':
		out = v.Append(out, 10)
	case 'x':
		out = v.Append(out, 16)
		if s.Flag('#') {
			basePrefix = "0x"
		}
	case 'X':
		out = v.Append(out, 16)
		for i, c := range out {
			if 'a' <= c && c <= 'f' {
				out[i] = c - ('a' - 'A')
			}
		}
		if s.Flag('#') {
			basePrefix = "0X"
		}
	case 's':
		out = v.Append(out, 10)
	}

	writePrefix := func() {
		if signPrefix != "" {
			io.WriteString(s, signPrefix) //nolint:errcheck
		}
		if basePrefix != "" {
			io.WriteString(s, basePrefix) //nolint:errcheck
		}
	}

	if w, ok := s.Width(); ok {
		pad := w - len(signPrefix) - len(basePrefix) - len(out)
		if s.Flag('0') {
			writePrefix()
			writePadding(s, zeros, pad)
			s.Write(out) //nolint:errcheck
		} else if s.Flag('-') {
			writePrefix()
			s.Write(out) //nolint:errcheck
			writePadding(s, spaces, pad)
		} else {
			writePadding(s, spaces, pad)
			writePrefix()
			s.Write(out) //nolint:errcheck
		}
		return
	}

	writePrefix()
	s.Write(out) //nolint:errcheck
}

// bigBase holds, for each base, the largest power of base that fits in a uint64,
// its exponent, and what is needed to divide by it with [divWW].
var bigBase = func() (t [len(digits) + 1]struct {
	bb    uint64 // base**n
	n     int
	shift uint   // number of leading zeros of bb
	recip uint64 // reciprocal(bb << shift)
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
		shift := uint(bits.LeadingZeros64(bb))
		t[base].bb, t[base].n = bb, n
		t[base].shift, t[base].recip = shift, reciprocal(bb<<shift)
	}
	return
}()

// reciprocal returns floor((2**128 - 1) / d) - 2**64 for a normalized d (d >= 2**63).
// See [divWW].
func reciprocal(d uint64) uint64 {
	r, _ := bits.Div64(^d, ^uint64(0), d)
	return r
}

// divWW returns the quotient and remainder of (x1:x0) / d,
// where d must be normalized (d >= 2**63), x1 < d, and m = reciprocal(d).
//
// It replaces the division by a multiplication by the precomputed reciprocal,
// which is much faster than [bits.Div64] on platforms without a 128/64-bit
// division instruction. See Niels Möller and Torbjörn Granlund,
// "Improved division by invariant integers", IEEE Transactions on Computers, 2011.
func divWW(x1, x0, d, m uint64) (q, r uint64) {
	qh, ql := bits.Mul64(x1, m)
	ql, c := bits.Add64(ql, x0, 0)
	qh, _ = bits.Add64(qh, x1, c)
	qh++
	r = x0 - qh*d
	if r > ql {
		qh--
		r += d
	}
	if r >= d {
		qh++
		r -= d
	}
	return qh, r
}

// divVW sets q = u / d and returns u % d, where u and q are little-endian
// (word 0 is the least significant) and have the same length.
// For long u it divides with [divWW] instead of [bits.Div64].
func divVW(q, u []uint64, d uint64) uint64 {
	var r uint64
	if len(u) < 3 {
		// Too short to pay for computing the reciprocal.
		for j := len(u) - 1; j >= 0; j-- {
			q[j], r = bits.Div64(r, u[j], d)
		}
		return r
	}

	// Dividing (r:x) << s by d << s gives the same quotient,
	// and the remainder shifted left by s.
	s := uint(bits.LeadingZeros64(d))
	dn := d << s
	m := reciprocal(dn)
	for j := len(u) - 1; j >= 0; j-- {
		x := u[j]
		q[j], r = divWW(r<<s|x>>(64-s), x<<s, dn, m)
		r >>= s
	}
	return r
}

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
	n, shift, recip := bigBase[base].n, bigBase[base].shift, bigBase[base].recip
	d := bigBase[base].bb << shift
	for len(u) > 1 {
		// u, r = u / bb, u % bb
		// Dividing (r:u[j]) << shift by bb << shift gives the same quotient,
		// and the remainder shifted left by shift.
		var r uint64
		for j := range u {
			x := u[j]
			u[j], r = divWW(r<<shift|x>>(64-shift), x<<shift, d, recip)
			r >>= shift
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

// formatBitsPow2 writes the digits of the multi-word unsigned integer u
// (most significant word first) in the given base, which must be a power of two,
// into a, ending at index i, and returns the index of the first digit written.
//
// It walks u one word at a time from the least significant word,
// so u is not shifted as a whole once per digit.
func formatBitsPow2(a []byte, i int, u []uint64, base int) int {
	shift := uint(bits.TrailingZeros(uint(base)))
	m := uint64(base) - 1 // == 1<<shift - 1

	for len(u) > 1 && u[0] == 0 {
		u = u[1:]
	}

	// acc holds accBits bits left over from the previous word,
	// which are fewer than shift bits.
	var acc uint64
	var accBits uint
	for k := len(u) - 1; k > 0; k-- {
		w := u[k]
		avail := uint(64)
		if accBits > 0 {
			// the digit straddles two words
			i--
			a[i] = digits[(acc|w<<accBits)&m]
			used := shift - accBits
			w >>= used
			avail -= used
		}
		for ; avail >= shift; avail -= shift {
			i--
			a[i] = digits[w&m]
			w >>= shift
		}
		acc, accBits = w, avail
	}

	// the most significant word: stop when no non-zero bits remain.
	w := u[0]
	if accBits > 0 {
		i--
		a[i] = digits[(acc|w<<accBits)&m]
		w >>= shift - accBits
	}
	for w > m {
		i--
		a[i] = digits[w&m]
		w >>= shift
	}
	if w != 0 || i == len(a) {
		i--
		a[i] = digits[w]
	}
	return i
}
