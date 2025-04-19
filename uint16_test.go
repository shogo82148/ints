package ints

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

func FuzzUint16_Add(f *testing.F) {
	f.Add(uint16(0), uint16(0))
	f.Add(uint16(1), uint16(0))
	f.Add(uint16(math.MaxUint16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, x, y uint16) {
		a := Uint16(x)
		b := Uint16(y)
		got := a.Add(b)
		want := Uint16(uint16(x + y))
		if got != want {
			t.Errorf("Uint16(%s).Add(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint16_Sub(f *testing.F) {
	f.Add(uint16(0), uint16(0))
	f.Add(uint16(0), uint16(1))
	f.Add(uint16(math.MaxUint16), uint16(1))
	f.Fuzz(func(t *testing.T, x, y uint16) {
		a := Uint16(x)
		b := Uint16(y)
		got := a.Sub(b)
		want := Uint16(uint16(x - y))
		if got != want {
			t.Errorf("Uint16(%s).Sub(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func FuzzUint16_Mul(f *testing.F) {
	f.Add(uint16(0), uint16(0))
	f.Add(uint16(1), uint16(0))
	f.Add(uint16(math.MaxUint16), uint16(math.MaxUint16))
	f.Fuzz(func(t *testing.T, x, y uint16) {
		a := Uint16(x)
		b := Uint16(y)
		got := a.Mul(b)
		want := Uint16(uint16(x * y))
		if got != want {
			t.Errorf("Uint16(%s).Mul(%s) = %d, want %d", a, b, got, want)
		}
	})
}

func TestUint16_Sign(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want int
	}{
		{0, 0},
		{1, 1},
		{math.MaxUint16, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Sign()
		if got != tc.want {
			t.Errorf("Uint16(%d).Sign() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint16_Neg(t *testing.T) {
	testCases := []struct {
		x    Uint16
		want Uint16
	}{
		{0, 0},
		{1, math.MaxUint16},
		{math.MaxUint16, 1},
	}

	for _, tc := range testCases {
		got := tc.x.Neg()
		if got != tc.want {
			t.Errorf("Uint16(%d).Neg() = %d, want %d", tc.x, got, tc.want)
		}
	}
}

func TestUint16_Text(t *testing.T) {
	var b big.Int
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			got := a.Text(base)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint16(%d).Text(%d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint16_Append(t *testing.T) {
	var b big.Int
	var buf []byte
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		b.SetInt64(int64(i))
		for base := 2; base <= 62; base++ {
			buf = a.Append(buf[:0], base)
			got := string(buf)
			want := b.Text(base)
			if got != want {
				t.Errorf("Uint16(%d).Append(buf, %d) = %q, want %q", i, base, got, want)
			}
		}
	}
}

func TestUint16_AppendText(t *testing.T) {
	var buf []byte
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		buf, err := a.AppendText(buf[:0])
		if err != nil {
			t.Fatal(err)
		}
		got := string(buf)
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint16(%d).AppendText() = %q, want %q", i, got, want)
		}
	}
}

func TestUint16_String(t *testing.T) {
	for i := range math.MaxUint16 + 1 {
		a := Uint16(i)
		got := a.String()
		want := strconv.FormatInt(int64(i), 10)
		if got != want {
			t.Errorf("Uint16(%d).String() = %q, want %q", i, got, want)
		}
	}
}
