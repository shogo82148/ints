package ints

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

func FuzzInt16_Add(f *testing.F) {
	f.Add(int16(0), int16(0))
	f.Add(int16(1), int16(0))
	f.Add(int16(math.MaxInt16), int16(math.MaxInt16))
	f.Add(int16(math.MinInt16), int16(math.MinInt16))
	f.Fuzz(func(t *testing.T, x, y int16) {
		a := Int16(x)
		b := Int16(y)
		got := a.Add(b)
		want := Int16(int16(x + y))
		if got != want {
			t.Errorf("Int16(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt16_Sub(f *testing.F) {
	f.Add(int16(0), int16(0))
	f.Add(int16(1), int16(0))
	f.Add(int16(math.MaxInt16), int16(-1))
	f.Add(int16(math.MinInt16), int16(1))
	f.Fuzz(func(t *testing.T, x, y int16) {
		a := Int16(x)
		b := Int16(y)
		got := a.Sub(b)
		want := Int16(int16(x - y))
		if got != want {
			t.Errorf("Int16(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzInt16_Mul(f *testing.F) {
	f.Add(int16(0), int16(0))
	f.Add(int16(1), int16(0))
	f.Add(int16(math.MaxInt16), int16(math.MaxInt16))
	f.Add(int16(math.MinInt16), int16(math.MinInt16))
	f.Fuzz(func(t *testing.T, x, y int16) {
		a := Int16(x)
		b := Int16(y)
		got := a.Mul(b)
		want := Int16(int16(x * y))
		if got != want {
			t.Errorf("Int16(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestInt16_Lsh(t *testing.T) {
	testCases := []struct {
		x    Int16
		i    uint
		want Int16
	}{
		{0, 1, 0},
		{1, 1, 2},
		{-1, 1, -2},
		{127, 1, 254},
		{-128, 1, -256},
	}

	for _, tc := range testCases {
		got := tc.x.Lsh(tc.i)
		if got != tc.want {
			t.Errorf("Int16(%d).Lsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt16_Rsh(t *testing.T) {
	testCases := []struct {
		x    Int16
		i    uint
		want Int16
	}{
		{0, 1, 0},
		{1, 1, 0},
		{-1, 1, -1},
		{127, 1, 63},
		{-128, 1, -64},
	}

	for _, tc := range testCases {
		got := tc.x.Rsh(tc.i)
		if got != tc.want {
			t.Errorf("Int16(%d).Rsh(%d) = %d, want %d", tc.x, tc.i, got, tc.want)
		}
	}
}

func TestInt16_Sign(t *testing.T) {
	testCases := []struct {
		x    Int16
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{math.MaxInt16, 1},
		{math.MinInt16, -1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Int16(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt16_Neg(t *testing.T) {
	testCases := []struct {
		x    Int16
		want Int16
	}{
		{0, 0},
		{1, -1},
		{-1, 1},
		{math.MaxInt16, -math.MaxInt16},
		{math.MinInt16, math.MinInt16},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Int16(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestInt16_Cmp(t *testing.T) {
	testCases := []struct {
		a, b Int16
		want int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, -1},
		{math.MaxInt16, math.MinInt16, 1},
		{math.MinInt16, math.MaxInt16, -1},
	}

	for _, tc := range testCases {
		got := tc.a.Cmp(tc.b)
		if got != tc.want {
			t.Errorf("Int16(%d).Cmp(%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestInt16_Text(t *testing.T) {
	var b big.Int
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int16(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt16_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Int16(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestInt16_AppendText(t *testing.T) {
	var buf []byte
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int16(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestInt16_String(t *testing.T) {
	for i := math.MinInt16; i <= math.MaxInt16; i++ {
		a := Int16(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Int16(%d).String() = %q, want %q", i, got, want)
		}
	}
}
