package ints

// Int8 returns a itself.
func (a Int8) Int8() Int8 {
	return a
}

// Int16 converts a to an Int16.
func (a Int8) Int16() Int16 {
	return Int16(a)
}

// Int32 converts a to an Int32.
func (a Int8) Int32() Int32 {
	return Int32(a)
}

// Int64 converts a to an Int64.
func (a Int8) Int64() Int64 {
	return Int64(a)
}

// Int128 converts a to an Int128.
func (a Int8) Int128() Int128 {
	return Int128{uint64(a >> 7), uint64(a)}
}

// Int256 converts a to an Int256.
func (a Int8) Int256() Int256 {
	return Int256{
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a),
	}
}

// Int512 converts a to an Int512.
func (a Int8) Int512() Int512 {
	return Int512{
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a),
	}
}

// Int1024 converts a to an Int1024.
func (a Int8) Int1024() Int1024 {
	return Int1024{
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a >> 7),
		uint64(a),
	}
}

// Int8 converts a to an Int8.
func (a Int16) Int8() Int8 {
	return Int8(a)
}

// Int16 returns a itself.
func (a Int16) Int16() Int16 {
	return a
}

// Int32 converts a to an Int32.
func (a Int16) Int32() Int32 {
	return Int32(a)
}

// Int64 converts a to an Int64.
func (a Int16) Int64() Int64 {
	return Int64(a)
}

// Int128 converts a to an Int128.
func (a Int16) Int128() Int128 {
	return Int128{uint64(a >> 15), uint64(a)}
}

// Int256 converts a to an Int256.
func (a Int16) Int256() Int256 {
	return Int256{
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a),
	}
}

// Int512 converts a to an Int512.
func (a Int16) Int512() Int512 {
	return Int512{
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a),
	}
}

// Int1024 converts a to an Int1024.
func (a Int16) Int1024() Int1024 {
	return Int1024{
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a >> 15),
		uint64(a),
	}
}

// Int8 converts a to an Int8.
func (a Int32) Int8() Int8 {
	return Int8(a)
}

// Int16 converts a to an Int16.
func (a Int32) Int16() Int16 {
	return Int16(a)
}

// Int32 returns a itself.
func (a Int32) Int32() Int32 {
	return a
}

// Int64 converts a to an Int64.
func (a Int32) Int64() Int64 {
	return Int64(a)
}

// Int128 converts a to an Int128.
func (a Int32) Int128() Int128 {
	return Int128{uint64(a >> 31), uint64(a)}
}

// Int256 converts a to an Int256.
func (a Int32) Int256() Int256 {
	return Int256{
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a),
	}
}

// Int512 converts a to an Int512.
func (a Int32) Int512() Int512 {
	return Int512{
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a),
	}
}

// Int1024 converts a to an Int1024.
func (a Int32) Int1024() Int1024 {
	return Int1024{
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a >> 31),
		uint64(a),
	}
}

// Int8 converts a to an Int8.
func (a Int64) Int8() Int8 {
	return Int8(a)
}

// Int16 converts a to an Int16.
func (a Int64) Int16() Int16 {
	return Int16(a)
}

// Int32 converts a to an Int32.
func (a Int64) Int32() Int32 {
	return Int32(a)
}

// Int64 returns a itself.
func (a Int64) Int64() Int64 {
	return a
}

// Int128 converts a to an Int128.
func (a Int64) Int128() Int128 {
	return Int128{uint64(a >> 63), uint64(a)}
}

// Int256 converts a to an Int256.
func (a Int64) Int256() Int256 {
	return Int256{
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a),
	}
}

// Int512 converts a to an Int512.
func (a Int64) Int512() Int512 {
	return Int512{
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a),
	}
}

// Int1024 converts a to an Int1024.
func (a Int64) Int1024() Int1024 {
	return Int1024{
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a >> 63),
		uint64(a),
	}
}

// Int8 converts a to an Int8.
func (a Int128) Int8() Int8 {
	return Int8(a[1])
}

// Int16 converts a to an Int16.
func (a Int128) Int16() Int16 {
	return Int16(a[1])
}

// Int32 converts a to an Int32.
func (a Int128) Int32() Int32 {
	return Int32(a[1])
}

// Int64 converts a to an Int64.
func (a Int128) Int64() Int64 {
	return Int64(a[1])
}

// Int128 returns a itself.
func (a Int128) Int128() Int128 {
	return a
}

// Int256 converts a to an Int256.
func (a Int128) Int256() Int256 {
	sign := uint64(int64(a[0]) >> 63)
	return Int256{
		sign,
		sign,
		a[0],
		a[1],
	}
}

// Int512 converts a to an Int512.
func (a Int128) Int512() Int512 {
	sign := uint64(int64(a[0]) >> 63)
	return Int512{
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		a[0],
		a[1],
	}
}

// Int1024 converts a to an Int1024.
func (a Int128) Int1024() Int1024 {
	sign := uint64(int64(a[0]) >> 63)
	return Int1024{
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		a[0],
		a[1],
	}
}

// Int8 converts a to an Int8.
func (a Int256) Int8() Int8 {
	return Int8(a[3])
}

// Int16 converts a to an Int16.
func (a Int256) Int16() Int16 {
	return Int16(a[3])
}

// Int32 converts a to an Int32.
func (a Int256) Int32() Int32 {
	return Int32(a[3])
}

// Int64 converts a to an Int64.
func (a Int256) Int64() Int64 {
	return Int64(a[3])
}

// Int128 converts a to an Int128.
func (a Int256) Int128() Int128 {
	return Int128{
		a[2],
		a[3],
	}
}

// Int256 returns a itself.
func (a Int256) Int256() Int256 {
	return a
}

// Int512 converts a to an Int512.
func (a Int256) Int512() Int512 {
	sign := uint64(int64(a[0]) >> 63)
	return Int512{
		sign,
		sign,
		sign,
		sign,
		a[0],
		a[1],
		a[2],
		a[3],
	}
}

// Int1024 converts a to an Int1024.
func (a Int256) Int1024() Int1024 {
	sign := uint64(int64(a[0]) >> 63)
	return Int1024{
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		a[0],
		a[1],
		a[2],
		a[3],
	}
}

// Int8 converts a to an Int8.
func (a Int512) Int8() Int8 {
	return Int8(a[7])
}

// Int16 converts a to an Int16.
func (a Int512) Int16() Int16 {
	return Int16(a[7])
}

// Int32 converts a to an Int32.
func (a Int512) Int32() Int32 {
	return Int32(a[7])
}

// Int64 converts a to an Int64.
func (a Int512) Int64() Int64 {
	return Int64(a[7])
}

// Int128 converts a to an Int128.
func (a Int512) Int128() Int128 {
	return Int128{
		a[6],
		a[7],
	}
}

// Int256 converts a to an Int256.
func (a Int512) Int256() Int256 {
	return Int256{
		a[4],
		a[5],
		a[6],
		a[7],
	}
}

// Int512 returns a itself.
func (a Int512) Int512() Int512 {
	return a
}

// Int1024 converts a to an Int1024.
func (a Int512) Int1024() Int1024 {
	sign := uint64(int64(a[0]) >> 63)
	return Int1024{
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		sign,
		a[0],
		a[1],
		a[2],
		a[3],
		a[4],
		a[5],
		a[6],
		a[7],
	}
}

// Int8 converts a to an Int8.
func (a Int1024) Int8() Int8 {
	return Int8(a[15])
}

// Int16 converts a to an Int16.
func (a Int1024) Int16() Int16 {
	return Int16(a[15])
}

// Int32 converts a to an Int32.
func (a Int1024) Int32() Int32 {
	return Int32(a[15])
}

// Int64 converts a to an Int64.
func (a Int1024) Int64() Int64 {
	return Int64(a[15])
}

// Int128 converts a to an Int128.
func (a Int1024) Int128() Int128 {
	return Int128{
		a[14],
		a[15],
	}
}

// Int256 converts a to an Int256.
func (a Int1024) Int256() Int256 {
	return Int256{
		a[12],
		a[13],
		a[14],
		a[15],
	}
}

// Int512 converts a to an Int512.
func (a Int1024) Int512() Int512 {
	return Int512{
		a[8],
		a[9],
		a[10],
		a[11],
		a[12],
		a[13],
		a[14],
		a[15],
	}
}

// Int1024 returns a itself.
func (a Int1024) Int1024() Int1024 {
	return a
}
