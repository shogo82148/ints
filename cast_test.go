package ints

import "testing"

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
