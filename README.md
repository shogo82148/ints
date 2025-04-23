# ints

[![Go Reference](https://pkg.go.dev/badge/github.com/shogo82148/ints.svg)](https://pkg.go.dev/github.com/shogo82148/ints)
[![Test](https://github.com/shogo82148/ints/actions/workflows/test.yaml/badge.svg)](https://github.com/shogo82148/ints/actions/workflows/test.yaml)

The ints package is a package for handling integer types of various sizes that are not included in the built-in types.

## SYNOPSIS

```go
import "github.com/shogo82148/bits"

func main() {
  // a and b are 128-bit integer.
  a := bits.Uint64(math.MaxUint64).Uint128()
  b := bits.Uint64(10).Uint128()

  var c ints.Uint128
  c = a.Add(b) // c = a + b
  c = a.Sub(b) // c = a - b
  c = a.Mul(b) // c = a * b
  c = a.Quo(b) // c = a / b
}
```
