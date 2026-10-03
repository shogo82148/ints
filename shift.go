package ints

import "math/bits"

// zeroMask returns all ones if x == 0, and zero otherwise, in constant time.
func zeroMask(x uint) uint64 {
	return uint64((x|-x)>>(bits.UintSize-1)) - 1
}

// lsh512 shifts a left by i bits in constant time.
// It is a barrel shifter: it moves whole words by 1, 2, 4, ... positions
// under masks derived from the bits of i, and then shifts the remaining bits.
func lsh512(a [8]uint64, i uint) [8]uint64 {
	u0, u1, u2, u3, u4, u5, u6, u7 := a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]
	var m uint64

	// move words by 1 if bit 6 of i is set
	m = -(uint64(i>>6) & 1)
	u0 = u1&m | u0&^m
	u1 = u2&m | u1&^m
	u2 = u3&m | u2&^m
	u3 = u4&m | u3&^m
	u4 = u5&m | u4&^m
	u5 = u6&m | u5&^m
	u6 = u7&m | u6&^m
	u7 &^= m

	// move words by 2 if bit 7 of i is set
	m = -(uint64(i>>7) & 1)
	u0 = u2&m | u0&^m
	u1 = u3&m | u1&^m
	u2 = u4&m | u2&^m
	u3 = u5&m | u3&^m
	u4 = u6&m | u4&^m
	u5 = u7&m | u5&^m
	u6 &^= m
	u7 &^= m

	// move words by 4 if bit 8 of i is set
	m = -(uint64(i>>8) & 1)
	u0 = u4&m | u0&^m
	u1 = u5&m | u1&^m
	u2 = u6&m | u2&^m
	u3 = u7&m | u3&^m
	u4 &^= m
	u5 &^= m
	u6 &^= m
	u7 &^= m

	// shift the remaining bits
	b := i & 63
	c := (63 - b) & 63 // == 64 - b, minus one to keep it below 64
	u0 = u0<<b | u1>>1>>c
	u1 = u1<<b | u2>>1>>c
	u2 = u2<<b | u3>>1>>c
	u3 = u3<<b | u4>>1>>c
	u4 = u4<<b | u5>>1>>c
	u5 = u5<<b | u6>>1>>c
	u6 = u6<<b | u7>>1>>c
	u7 <<= b

	// shifts of 512 bits or more result in zero
	z := zeroMask(i >> 9)
	return [8]uint64{u0 & z, u1 & z, u2 & z, u3 & z, u4 & z, u5 & z, u6 & z, u7 & z}
}

// rsh512 shifts a right by i bits in constant time, filling the vacated bits with fill,
// which must be zero (logical shift) or all ones (arithmetic shift of a negative value).
// See [lsh512] for the algorithm.
func rsh512(a [8]uint64, i uint, fill uint64) [8]uint64 {
	u0, u1, u2, u3, u4, u5, u6, u7 := a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]
	var m uint64

	// move words by 1 if bit 6 of i is set
	m = -(uint64(i>>6) & 1)
	u7 = u6&m | u7&^m
	u6 = u5&m | u6&^m
	u5 = u4&m | u5&^m
	u4 = u3&m | u4&^m
	u3 = u2&m | u3&^m
	u2 = u1&m | u2&^m
	u1 = u0&m | u1&^m
	u0 = fill&m | u0&^m

	// move words by 2 if bit 7 of i is set
	m = -(uint64(i>>7) & 1)
	u7 = u5&m | u7&^m
	u6 = u4&m | u6&^m
	u5 = u3&m | u5&^m
	u4 = u2&m | u4&^m
	u3 = u1&m | u3&^m
	u2 = u0&m | u2&^m
	u1 = fill&m | u1&^m
	u0 = fill&m | u0&^m

	// move words by 4 if bit 8 of i is set
	m = -(uint64(i>>8) & 1)
	u7 = u3&m | u7&^m
	u6 = u2&m | u6&^m
	u5 = u1&m | u5&^m
	u4 = u0&m | u4&^m
	u3 = fill&m | u3&^m
	u2 = fill&m | u2&^m
	u1 = fill&m | u1&^m
	u0 = fill&m | u0&^m

	// shift the remaining bits
	b := i & 63
	c := (63 - b) & 63 // == 64 - b, minus one to keep it below 64
	u7 = u7>>b | u6<<1<<c
	u6 = u6>>b | u5<<1<<c
	u5 = u5>>b | u4<<1<<c
	u4 = u4>>b | u3<<1<<c
	u3 = u3>>b | u2<<1<<c
	u2 = u2>>b | u1<<1<<c
	u1 = u1>>b | u0<<1<<c
	u0 = u0>>b | fill<<1<<c

	// shifts of 512 bits or more result in fill
	z := zeroMask(i >> 9)
	return [8]uint64{u0&z | fill&^z, u1&z | fill&^z, u2&z | fill&^z, u3&z | fill&^z, u4&z | fill&^z, u5&z | fill&^z, u6&z | fill&^z, u7&z | fill&^z}
}

// lsh1024 shifts a left by i bits in constant time.
// It is a barrel shifter: it moves whole words by 1, 2, 4, ... positions
// under masks derived from the bits of i, and then shifts the remaining bits.
func lsh1024(a [16]uint64, i uint) [16]uint64 {
	u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 := a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15]
	var m uint64

	// move words by 1 if bit 6 of i is set
	m = -(uint64(i>>6) & 1)
	u0 = u1&m | u0&^m
	u1 = u2&m | u1&^m
	u2 = u3&m | u2&^m
	u3 = u4&m | u3&^m
	u4 = u5&m | u4&^m
	u5 = u6&m | u5&^m
	u6 = u7&m | u6&^m
	u7 = u8&m | u7&^m
	u8 = u9&m | u8&^m
	u9 = u10&m | u9&^m
	u10 = u11&m | u10&^m
	u11 = u12&m | u11&^m
	u12 = u13&m | u12&^m
	u13 = u14&m | u13&^m
	u14 = u15&m | u14&^m
	u15 &^= m

	// move words by 2 if bit 7 of i is set
	m = -(uint64(i>>7) & 1)
	u0 = u2&m | u0&^m
	u1 = u3&m | u1&^m
	u2 = u4&m | u2&^m
	u3 = u5&m | u3&^m
	u4 = u6&m | u4&^m
	u5 = u7&m | u5&^m
	u6 = u8&m | u6&^m
	u7 = u9&m | u7&^m
	u8 = u10&m | u8&^m
	u9 = u11&m | u9&^m
	u10 = u12&m | u10&^m
	u11 = u13&m | u11&^m
	u12 = u14&m | u12&^m
	u13 = u15&m | u13&^m
	u14 &^= m
	u15 &^= m

	// move words by 4 if bit 8 of i is set
	m = -(uint64(i>>8) & 1)
	u0 = u4&m | u0&^m
	u1 = u5&m | u1&^m
	u2 = u6&m | u2&^m
	u3 = u7&m | u3&^m
	u4 = u8&m | u4&^m
	u5 = u9&m | u5&^m
	u6 = u10&m | u6&^m
	u7 = u11&m | u7&^m
	u8 = u12&m | u8&^m
	u9 = u13&m | u9&^m
	u10 = u14&m | u10&^m
	u11 = u15&m | u11&^m
	u12 &^= m
	u13 &^= m
	u14 &^= m
	u15 &^= m

	// move words by 8 if bit 9 of i is set
	m = -(uint64(i>>9) & 1)
	u0 = u8&m | u0&^m
	u1 = u9&m | u1&^m
	u2 = u10&m | u2&^m
	u3 = u11&m | u3&^m
	u4 = u12&m | u4&^m
	u5 = u13&m | u5&^m
	u6 = u14&m | u6&^m
	u7 = u15&m | u7&^m
	u8 &^= m
	u9 &^= m
	u10 &^= m
	u11 &^= m
	u12 &^= m
	u13 &^= m
	u14 &^= m
	u15 &^= m

	// shift the remaining bits
	b := i & 63
	c := (63 - b) & 63 // == 64 - b, minus one to keep it below 64
	u0 = u0<<b | u1>>1>>c
	u1 = u1<<b | u2>>1>>c
	u2 = u2<<b | u3>>1>>c
	u3 = u3<<b | u4>>1>>c
	u4 = u4<<b | u5>>1>>c
	u5 = u5<<b | u6>>1>>c
	u6 = u6<<b | u7>>1>>c
	u7 = u7<<b | u8>>1>>c
	u8 = u8<<b | u9>>1>>c
	u9 = u9<<b | u10>>1>>c
	u10 = u10<<b | u11>>1>>c
	u11 = u11<<b | u12>>1>>c
	u12 = u12<<b | u13>>1>>c
	u13 = u13<<b | u14>>1>>c
	u14 = u14<<b | u15>>1>>c
	u15 <<= b

	// shifts of 1024 bits or more result in zero
	z := zeroMask(i >> 10)
	return [16]uint64{u0 & z, u1 & z, u2 & z, u3 & z, u4 & z, u5 & z, u6 & z, u7 & z, u8 & z, u9 & z, u10 & z, u11 & z, u12 & z, u13 & z, u14 & z, u15 & z}
}

// rsh1024 shifts a right by i bits in constant time, filling the vacated bits with fill,
// which must be zero (logical shift) or all ones (arithmetic shift of a negative value).
// See [lsh1024] for the algorithm.
func rsh1024(a [16]uint64, i uint, fill uint64) [16]uint64 {
	u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15 := a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15]
	var m uint64

	// move words by 1 if bit 6 of i is set
	m = -(uint64(i>>6) & 1)
	u15 = u14&m | u15&^m
	u14 = u13&m | u14&^m
	u13 = u12&m | u13&^m
	u12 = u11&m | u12&^m
	u11 = u10&m | u11&^m
	u10 = u9&m | u10&^m
	u9 = u8&m | u9&^m
	u8 = u7&m | u8&^m
	u7 = u6&m | u7&^m
	u6 = u5&m | u6&^m
	u5 = u4&m | u5&^m
	u4 = u3&m | u4&^m
	u3 = u2&m | u3&^m
	u2 = u1&m | u2&^m
	u1 = u0&m | u1&^m
	u0 = fill&m | u0&^m

	// move words by 2 if bit 7 of i is set
	m = -(uint64(i>>7) & 1)
	u15 = u13&m | u15&^m
	u14 = u12&m | u14&^m
	u13 = u11&m | u13&^m
	u12 = u10&m | u12&^m
	u11 = u9&m | u11&^m
	u10 = u8&m | u10&^m
	u9 = u7&m | u9&^m
	u8 = u6&m | u8&^m
	u7 = u5&m | u7&^m
	u6 = u4&m | u6&^m
	u5 = u3&m | u5&^m
	u4 = u2&m | u4&^m
	u3 = u1&m | u3&^m
	u2 = u0&m | u2&^m
	u1 = fill&m | u1&^m
	u0 = fill&m | u0&^m

	// move words by 4 if bit 8 of i is set
	m = -(uint64(i>>8) & 1)
	u15 = u11&m | u15&^m
	u14 = u10&m | u14&^m
	u13 = u9&m | u13&^m
	u12 = u8&m | u12&^m
	u11 = u7&m | u11&^m
	u10 = u6&m | u10&^m
	u9 = u5&m | u9&^m
	u8 = u4&m | u8&^m
	u7 = u3&m | u7&^m
	u6 = u2&m | u6&^m
	u5 = u1&m | u5&^m
	u4 = u0&m | u4&^m
	u3 = fill&m | u3&^m
	u2 = fill&m | u2&^m
	u1 = fill&m | u1&^m
	u0 = fill&m | u0&^m

	// move words by 8 if bit 9 of i is set
	m = -(uint64(i>>9) & 1)
	u15 = u7&m | u15&^m
	u14 = u6&m | u14&^m
	u13 = u5&m | u13&^m
	u12 = u4&m | u12&^m
	u11 = u3&m | u11&^m
	u10 = u2&m | u10&^m
	u9 = u1&m | u9&^m
	u8 = u0&m | u8&^m
	u7 = fill&m | u7&^m
	u6 = fill&m | u6&^m
	u5 = fill&m | u5&^m
	u4 = fill&m | u4&^m
	u3 = fill&m | u3&^m
	u2 = fill&m | u2&^m
	u1 = fill&m | u1&^m
	u0 = fill&m | u0&^m

	// shift the remaining bits
	b := i & 63
	c := (63 - b) & 63 // == 64 - b, minus one to keep it below 64
	u15 = u15>>b | u14<<1<<c
	u14 = u14>>b | u13<<1<<c
	u13 = u13>>b | u12<<1<<c
	u12 = u12>>b | u11<<1<<c
	u11 = u11>>b | u10<<1<<c
	u10 = u10>>b | u9<<1<<c
	u9 = u9>>b | u8<<1<<c
	u8 = u8>>b | u7<<1<<c
	u7 = u7>>b | u6<<1<<c
	u6 = u6>>b | u5<<1<<c
	u5 = u5>>b | u4<<1<<c
	u4 = u4>>b | u3<<1<<c
	u3 = u3>>b | u2<<1<<c
	u2 = u2>>b | u1<<1<<c
	u1 = u1>>b | u0<<1<<c
	u0 = u0>>b | fill<<1<<c

	// shifts of 1024 bits or more result in fill
	z := zeroMask(i >> 10)
	return [16]uint64{u0&z | fill&^z, u1&z | fill&^z, u2&z | fill&^z, u3&z | fill&^z, u4&z | fill&^z, u5&z | fill&^z, u6&z | fill&^z, u7&z | fill&^z, u8&z | fill&^z, u9&z | fill&^z, u10&z | fill&^z, u11&z | fill&^z, u12&z | fill&^z, u13&z | fill&^z, u14&z | fill&^z, u15&z | fill&^z}
}
