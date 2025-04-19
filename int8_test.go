package ints

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

func FuzzInt8_Add(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(1), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Fuzz(func(t *testing.T, x, y int8) {
		a := Int8(x)
		b := Int8(y)
		got := a.Add(b)
		want := Int8(int8(x + y))
		if got != want {
			t.Errorf("Int8(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt8_Sub(f *testing.F) {
	f.Add(int8(0), int8(0))
	f.Add(int8(1), int8(0))
	f.Add(int8(math.MaxInt8), int8(math.MaxInt8))
	f.Add(int8(math.MinInt8), int8(math.MinInt8))
	f.Fuzz(func(t *testing.T, x, y int8) {
		a := Int8(x)
		b := Int8(y)
		got := a.Sub(b)
		want := Int8(int8(x - y))
		if got != want {
			t.Errorf("Int8(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestInt8_Sign(t *testing.T) {
	testCases := []struct {
		x    Int8
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{127, 1},
		{-128, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int8(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt8_Neg(t *testing.T) {
	testCases := []struct {
		x    Int8
		want Int8
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{127, -127},
		{-128, -128},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int8(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt8_Text(t *testing.T) {
	var b big.Int
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int8(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt8_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int8(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt8_AppendText(t *testing.T) {
	var buf []byte
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int8(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestInt8_String(t *testing.T) {
	for i := math.MinInt8; i <= math.MaxInt8; i++ {
		a := Int8(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int8(%d).String() = %q, want %q", i, got, want)
		}
	}
}
