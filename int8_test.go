package ints

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

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
