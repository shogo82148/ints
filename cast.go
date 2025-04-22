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
