package ints

import (
	"math"
	"testing"
)

func TestInt8_Int8(t *testing.T) {
	testCases := []struct {
		a    Int8
		want Int8
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int8(%d).Int8() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int16(t *testing.T) {
	testCases := []struct {
		a    Int8
		want Int16
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int8(%d).Int16() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int32(t *testing.T) {
	testCases := []struct {
		a    Int8
		want Int32
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int8(%d).Int32() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int64(t *testing.T) {
	testCases := []struct {
		a    Int8
		want Int64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int8(%d).Int64() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int8
		want Int128
	}{
		{0, Int128{0, 0}},
		{1, Int128{0, 1}},
		{-1, Int128{M - 1, M - 1}},
		{127, Int128{0, 127}},
		{-128, Int128{M - 1, M - 128}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int8(%d).Int128() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int8
		want Int256
	}{
		{0, Int256{0, 0, 0, 0}},
		{1, Int256{0, 0, 0, 1}},
		{-1, Int256{M - 1, M - 1, M - 1, M - 1}},
		{127, Int256{0, 0, 0, 127}},
		{-128, Int256{M - 1, M - 1, M - 1, M - 128}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int8(%d).Int256() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int8
		want Int512
	}{
		{0, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{-1, Int512{M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1}},
		{127, Int512{0, 0, 0, 0, 0, 0, 0, 127}},
		{-128, Int512{M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 128}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int8(%d).Int512() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt8_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int8
		want Int1024
	}{
		{
			0,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			-1,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
		{
			127,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 127,
			},
		},
		{
			-128,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 128,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int8(%d).Int1024() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int8(t *testing.T) {
	testCases := []struct {
		a    Int16
		want Int8
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int16(%d).Int8() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int16(t *testing.T) {
	testCases := []struct {
		a    Int16
		want Int16
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int16(%d).Int16() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int32(t *testing.T) {
	testCases := []struct {
		a    Int16
		want Int32
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int16(%d).Int32() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int64(t *testing.T) {
	testCases := []struct {
		a    Int16
		want Int64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int16(%d).Int64() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int16
		want Int128
	}{
		{0, Int128{0, 0}},
		{1, Int128{0, 1}},
		{-1, Int128{M - 1, M - 1}},
		{math.MaxInt16, Int128{0, math.MaxInt16}},
		{math.MinInt16, Int128{M - 1, M + math.MinInt16}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int16(%d).Int128() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int16
		want Int256
	}{
		{0, Int256{0, 0, 0, 0}},
		{1, Int256{0, 0, 0, 1}},
		{-1, Int256{M - 1, M - 1, M - 1, M - 1}},
		{math.MaxInt16, Int256{0, 0, 0, math.MaxInt16}},
		{math.MinInt16, Int256{M - 1, M - 1, M - 1, M + math.MinInt16}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int16(%d).Int256() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int16
		want Int512
	}{
		{0, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{-1, Int512{M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1}},
		{math.MaxInt16, Int512{0, 0, 0, 0, 0, 0, 0, math.MaxInt16}},
		{math.MinInt16, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M + math.MinInt16,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int16(%d).Int512() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt16_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int16
		want Int1024
	}{
		{
			0,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			-1,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
		{
			math.MaxInt16,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, math.MaxInt16,
			},
		},
		{
			math.MinInt16,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M + math.MinInt16,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int16(%d).Int1024() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int8(t *testing.T) {
	testCases := []struct {
		a    Int32
		want Int8
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int32(%d).Int8() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int16(t *testing.T) {
	testCases := []struct {
		a    Int32
		want Int16
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
		{0x10000, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int32(%d).Int16() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int32(t *testing.T) {
	testCases := []struct {
		a    Int32
		want Int32
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int32(%d).Int32() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int64(t *testing.T) {
	testCases := []struct {
		a    Int32
		want Int64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int32(%d).Int64() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int32
		want Int128
	}{
		{0, Int128{0, 0}},
		{1, Int128{0, 1}},
		{-1, Int128{M - 1, M - 1}},
		{math.MaxInt32, Int128{0, math.MaxInt32}},
		{math.MinInt32, Int128{M - 1, M + math.MinInt32}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int32(%d).Int128() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int32
		want Int256
	}{
		{0, Int256{0, 0, 0, 0}},
		{1, Int256{0, 0, 0, 1}},
		{-1, Int256{M - 1, M - 1, M - 1, M - 1}},
		{math.MaxInt32, Int256{0, 0, 0, math.MaxInt32}},
		{math.MinInt32, Int256{M - 1, M - 1, M - 1, M + math.MinInt32}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int32(%d).Int256() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int32
		want Int512
	}{
		{0, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{-1, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}},
		{math.MaxInt32, Int512{0, 0, 0, 0, 0, 0, 0, math.MaxInt32}},
		{math.MinInt32, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M + math.MinInt32,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int32(%d).Int512() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt32_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int32
		want Int1024
	}{
		{
			0,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			-1,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
		{
			math.MaxInt32,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, math.MaxInt32,
			},
		},
		{
			math.MinInt32,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M + math.MinInt32,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int32(%d).Int1024() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int8(t *testing.T) {
	testCases := []struct {
		a    Int64
		want Int8
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int64(%d).Int8() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int16(t *testing.T) {
	testCases := []struct {
		a    Int64
		want Int16
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
		{0x10000, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int64(%d).Int16() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int32(t *testing.T) {
	testCases := []struct {
		a    Int64
		want Int32
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int64(%d).Int32() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int64(t *testing.T) {
	testCases := []struct {
		a    Int64
		want Int64
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0x100, 0x100},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int64(%d).Int64() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int64
		want Int128
	}{
		{0, Int128{0, 0}},
		{1, Int128{0, 1}},
		{-1, Int128{M - 1, M - 1}},
		{math.MaxInt64, Int128{0, math.MaxInt64}},
		{math.MinInt64, Int128{M - 1, M + math.MinInt64}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int64(%d).Int128() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int64
		want Int256
	}{
		{0, Int256{0, 0, 0, 0}},
		{1, Int256{0, 0, 0, 1}},
		{-1, Int256{M - 1, M - 1, M - 1, M - 1}},
		{math.MaxInt64, Int256{0, 0, 0, math.MaxInt64}},
		{math.MinInt64, Int256{M - 1, M - 1, M - 1, M + math.MinInt64}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int64(%d).Int256() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int64
		want Int512
	}{
		{0, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{-1, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}},
		{math.MaxInt64, Int512{0, 0, 0, 0, 0, 0, 0, math.MaxInt64}},
		{math.MinInt64, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M + math.MinInt64,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int64(%d).Int512() = %d, want %d", tc.a, got, tc.want)
		}
	}
}

func TestInt64_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int64
		want Int1024
	}{
		{
			0,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			-1,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
		{
			math.MaxInt64,
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, math.MaxInt64,
			},
		},
		{
			math.MinInt64,
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M + math.MinInt64,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int64(%#016x).Int1024() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int8(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int8
	}{
		{Int128{0, 0}, 0},
		{Int128{0, 1}, 1},
		{Int128{M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int8() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int16(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int16
	}{
		{Int128{0, 0}, 0},
		{Int128{0, 1}, 1},
		{Int128{M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int16() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int32(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int32
	}{
		{Int128{0, 0}, 0},
		{Int128{0, 1}, 1},
		{Int128{M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int32() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int64(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int64
	}{
		{Int128{0, 0}, 0},
		{Int128{0, 1}, 1},
		{Int128{M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int64() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int128
	}{
		{Int128{0, 0}, Int128{0, 0}},
		{Int128{0, 1}, Int128{0, 1}},
		{Int128{M - 1, M - 1}, Int128{M - 1, M - 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int256
	}{
		{Int128{0, 0}, Int256{0, 0, 0, 0}},
		{Int128{0, 1}, Int256{0, 0, 0, 1}},
		{Int128{M - 1, M - 1}, Int256{M - 1, M - 1, M - 1, M - 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int256() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int512
	}{
		{Int128{0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int128{0, 1}, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{
			Int128{M - 1, M - 1},
			Int512{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int512() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt128_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int128
		want Int1024
	}{
		{
			Int128{0, 0},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Int128{0, 1},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			Int128{M - 1, M - 1},
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int128(%#032x).Int1024() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int8(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int8
	}{
		{Int256{0, 0, 0, 0}, 0},
		{Int256{0, 0, 0, 1}, 1},
		{Int256{M - 1, M - 1, M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int8() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int16(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int16
	}{
		{Int256{0, 0, 0, 0}, 0},
		{Int256{0, 0, 0, 1}, 1},
		{Int256{M - 1, M - 1, M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int16() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int32(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int32
	}{
		{Int256{0, 0, 0, 0}, 0},
		{Int256{0, 0, 0, 1}, 1},
		{Int256{M - 1, M - 1, M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int32() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int64(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int64
	}{
		{Int256{0, 0, 0, 0}, 0},
		{Int256{0, 0, 0, 1}, 1},
		{Int256{M - 1, M - 1, M - 1, M - 1}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int64() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int128
	}{
		{Int256{0, 0, 0, 0}, Int128{0, 0}},
		{Int256{0, 0, 0, 1}, Int128{0, 1}},
		{Int256{M - 1, M - 1, M - 1, M - 1}, Int128{M - 1, M - 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int256
	}{
		{Int256{0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int256{0, 0, 0, 1}, Int256{0, 0, 0, 1}},
		{Int256{M - 1, M - 1, M - 1, M - 1}, Int256{M - 1, M - 1, M - 1, M - 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int256() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int512
	}{
		{Int256{0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int256{0, 0, 0, 1}, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{Int256{M - 1, M - 1, M - 1, M - 1}, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int512() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt256_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int256
		want Int1024
	}{
		{
			Int256{0, 0, 0, 0},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Int256{0, 0, 0, 1},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			Int256{M - 1, M - 1, M - 1, M - 1},
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int256(%#032x).Int1024() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int8(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int8
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int8() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int16(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int16
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int16() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int32(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int32
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int32() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int64(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int64
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int64() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int128
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int128{0, 0}},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, Int128{0, 1}},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, Int128{M - 1, M - 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int128() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int256
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int256{0, 0, 0, 0}},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, Int256{0, 0, 0, 1}},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, Int256{
			M - 1, M - 1, M - 1, M - 1,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int256() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int512
	}{
		{Int512{0, 0, 0, 0, 0, 0, 0, 0}, Int512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Int512{0, 0, 0, 0, 0, 0, 0, 1}, Int512{0, 0, 0, 0, 0, 0, 0, 1}},
		{Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}, Int512{
			M - 1, M - 1, M - 1, M - 1,
			M - 1, M - 1, M - 1, M - 1,
		}},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt512_Int1024(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int512
		want Int1024
	}{
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Int512{0, 0, 0, 0, 0, 0, 0, 1},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			Int512{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int512(%#0128x).Int1024() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int8(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int8
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			0,
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
			},
			-1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int8()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int16(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int16
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			0,
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
			},
			-1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int16()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int32(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int32
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			0,
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
			},
			-1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int32()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int64(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int64
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			0,
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
			},
			-1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int64()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int128(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int128
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			Int128{0, 0},
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Int128{0, 1},
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1, M - 1,
			},
			Int128{M - 1, M - 1},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int128()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int256(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int256
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			Int256{0, 0, 0, 0},
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Int256{0, 0, 0, 1},
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
			Int256{M - 1, M - 1, M - 1, M - 1},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int256()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int512(t *testing.T) {
	const M = 0x1_00000000_00000000
	testCases := []struct {
		a    Int1024
		want Int512
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			Int512{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Int512{0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			Int1024{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
			Int512{
				M - 1, M - 1, M - 1, M - 1,
				M - 1, M - 1, M - 1, M - 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int512()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestInt1024_Int1024(t *testing.T) {
	testCases := []struct {
		a    Int1024
		want Int1024
	}{
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Int1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Int1024()
		if got != tc.want {
			t.Errorf("Int1024(%#0256x).Int1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint8
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint16
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint32
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint64
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint128
	}{
		{0, Uint128{0, 0}},
		{1, Uint128{0, 1}},
		{0xFF, Uint128{0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint256
	}{
		{0, Uint256{0, 0, 0, 0}},
		{1, Uint256{0, 0, 0, 1}},
		{0xFF, Uint256{0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint512
	}{
		{0, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
		{0xFF, Uint512{0, 0, 0, 0, 0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint8_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint8
		want Uint1024
	}{
		{
			0,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
		{
			0xff,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0xff,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint8(%#02x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint8
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint16
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint32
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint64
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint128
	}{
		{0, Uint128{0, 0}},
		{1, Uint128{0, 1}},
		{0xFF, Uint128{0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint256
	}{
		{0, Uint256{0, 0, 0, 0}},
		{1, Uint256{0, 0, 0, 1}},
		{0xFF, Uint256{0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint512
	}{
		{0, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
		{0xFF, Uint512{0, 0, 0, 0, 0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint16_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint16
		want Uint1024
	}{
		{
			0,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint16(%#04x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint8
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint16
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint32
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint64
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint128
	}{
		{0, Uint128{0, 0}},
		{1, Uint128{0, 1}},
		{0xFF, Uint128{0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint256
	}{
		{0, Uint256{0, 0, 0, 0}},
		{1, Uint256{0, 0, 0, 1}},
		{0xFF, Uint256{0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint512
	}{
		{0, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
		{0xFF, Uint512{0, 0, 0, 0, 0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint32_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint32
		want Uint1024
	}{
		{
			0,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint32(%#08x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint8
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint16
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint32
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint64
	}{
		{0, 0},
		{1, 1},
		{0xFF, 0xFF},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint128
	}{
		{0, Uint128{0, 0}},
		{1, Uint128{0, 1}},
		{0xFF, Uint128{0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint256
	}{
		{0, Uint256{0, 0, 0, 0}},
		{1, Uint256{0, 0, 0, 1}},
		{0xFF, Uint256{0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint512
	}{
		{0, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{1, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
		{0xFF, Uint512{0, 0, 0, 0, 0, 0, 0, 0xFF}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint64_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint64
		want Uint1024
	}{
		{
			0,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			1,
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint64(%#016x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint8
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{0, 0xFF}, 0xFF},
		{Uint128{1, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint16
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{0, 0xFF}, 0xFF},
		{Uint128{1, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint32
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{0, 0xFF}, 0xFF},
		{Uint128{1, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint64
	}{
		{Uint128{0, 0}, 0},
		{Uint128{0, 1}, 1},
		{Uint128{0, 0xFF}, 0xFF},
		{Uint128{1, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint128
	}{
		{Uint128{0, 0}, Uint128{0, 0}},
		{Uint128{0, 1}, Uint128{0, 1}},
		{Uint128{0xFF, 0}, Uint128{0xFF, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint256
	}{
		{Uint128{0, 0}, Uint256{0, 0, 0, 0}},
		{Uint128{0, 1}, Uint256{0, 0, 0, 1}},
		{Uint128{0xFF, 0}, Uint256{0, 0, 0xFF, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint512
	}{
		{Uint128{0, 0}, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Uint128{0, 1}, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
		{Uint128{0xFF, 0}, Uint512{0, 0, 0, 0, 0, 0, 0xFF, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint128_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint128
		want Uint1024
	}{
		{
			Uint128{0, 0},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Uint128{0, 1},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint128(%#032x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint8
	}{
		{Uint256{0, 0, 0, 0}, 0},
		{Uint256{0, 0, 0, 1}, 1},
		{Uint256{0, 0, 0xFF, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint16
	}{
		{Uint256{0, 0, 0, 0}, 0},
		{Uint256{0, 0, 0, 1}, 1},
		{Uint256{0, 0, 0xFF, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint32
	}{
		{Uint256{0, 0, 0, 0}, 0},
		{Uint256{0, 0, 0, 1}, 1},
		{Uint256{0, 0, 0xFF, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint64
	}{
		{Uint256{0, 0, 0, 0}, 0},
		{Uint256{0, 0, 0, 1}, 1},
		{Uint256{0, 0, 0xFF, 0}, 0},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint128
	}{
		{Uint256{0, 0, 0, 0}, Uint128{0, 0}},
		{Uint256{0, 0, 0, 1}, Uint128{0, 1}},
		{Uint256{0, 0, 0xFF, 0}, Uint128{0xFF, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint256
	}{
		{Uint256{0, 0, 0, 0}, Uint256{0, 0, 0, 0}},
		{Uint256{0, 0, 0, 1}, Uint256{0, 0, 0, 1}},
		{Uint256{0xFF, 0, 0xFF, 0}, Uint256{0xFF, 0, 0xFF, 0}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint512
	}{
		{Uint256{0, 0, 0, 0}, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Uint256{0, 0, 0, 1}, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint256_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint256
		want Uint1024
	}{
		{
			Uint256{0, 0, 0, 0},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Uint256{0, 0, 0, 1},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint256(%#064x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint8
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint16
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint32
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint64
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, 0},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, 1},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint128
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, Uint128{0, 0}},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, Uint128{0, 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint256
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, Uint256{0, 0, 0, 0}},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, Uint256{0, 0, 0, 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint512
	}{
		{Uint512{0, 0, 0, 0, 0, 0, 0, 0}, Uint512{0, 0, 0, 0, 0, 0, 0, 0}},
		{Uint512{0, 0, 0, 0, 0, 0, 0, 1}, Uint512{0, 0, 0, 0, 0, 0, 0, 1}},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint512_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint512
		want Uint1024
	}{
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 0},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		{
			Uint512{0, 0, 0, 0, 0, 0, 0, 1},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint512(%#0128x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint8(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint8
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint8()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint8() = %#02x, want %#02x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint16(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint16
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint16()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint16() = %#04x, want %#04x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint32(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint32
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint32()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint32() = %#08x, want %#08x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint64(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint64
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			1,
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint64()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint64() = %#016x, want %#016x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint128(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint128
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Uint128{0, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint128()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint128() = %#032x, want %#032x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint256(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint256
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Uint256{0, 0, 0, 1},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint256()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint256() = %#064x, want %#064x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint512(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint512
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Uint512{0, 0, 0, 0, 0, 0, 0, 0x01},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint512()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint512() = %#0128x, want %#0128x", tc.a, got, tc.want)
		}
	}
}

func TestUint1024_Uint1024(t *testing.T) {
	testCases := []struct {
		a    Uint1024
		want Uint1024
	}{
		{
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
			Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 1,
			},
		},
	}

	for _, tc := range testCases {
		got := tc.a.Uint1024()
		if got != tc.want {
			t.Errorf("Uint1024(%#0256x).Uint1024() = %#0256x, want %#0256x", tc.a, got, tc.want)
		}
	}
}
