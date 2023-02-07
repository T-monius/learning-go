# Variables

`var` _statement_ declares a list of variables
- Type is last
- Can be at package or function level

```go
package main

import "fmt"

var c, python, java bool

func main() {
  var i int
  fmt.Println(i, c, python, java)
}
```

## Variables w/ initializers
- Var declarations can include initializers
  - One per variable
  - __If an initializer is present, the type can be omitted__

```go
package main

import "fmt"

var i, j int = 1, 2

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}
```

## Short variable declaratoins

Short assignment statement
- `:=`
- Can be used inside a function
  - In place of `var`
  - __Only for implicit type__

Outside a function, _every statement begins w/ a keyword_

## Basic Types

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

Variable declarations can be "factored" into blocks.

```go
package main

import (
  "fmt"
  "math/cmplx"
)

var (
  ToBe   bool       = false
  MaxInt uint64     = 1<<64 - 1
  z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
  fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
  fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
  fmt.Printf("Type: %T Value: %v\n", z, z)
}
```

## Zero values

Variables declared w/o an explicit initial value are given their _zero value_:
- `0` for numeric types
- `false` for boolean types
- `""` (empty string) for strings.

```go
package main

import "fmt"

func main() {
  var i int
  var f float64
  var b bool
  var s string
  fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

## Type conversions

`T(v)`: expression that converts a value (`v`) to the type of `T`.

Numeric conversions:
```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)

// simplified
i := 42
f := float(i)
u := uint(f)
```

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y int = 3, 4
  var f float64 = math.Sqrt(float64(x*x + y*y))
  var z uint = uint(f)
  fmt.Println(x, y, z)
}
```

Go assignment between items of different types requires an explicit conversion
- Unlike C
- _Can be shown by removing type conversions in variable declarations in above example._

## Type inference

Variable's type can be inferred from the value on the right hand side of assignment.

When the right hand side of the declaration is typed, the new variable is of the same type:

```go
var i int
j := i // j is an int
```

When the right hand side contains an untyped numeric constant, the new variable may be an `int`, `float64`, or `complex128` depending on the precision of the constant.

```go
i := 42            // int
f := 3.142         // float64
g := 0.8567 + 0.5i //
```

## Constants

Constants declared like variables but with the `const` keyword.
- Character, string, boolean, or numeric values
- Cannot be declared w/ short form (e.g. `:=`)

```go
package main

import "fmt"

const Pi = 3.14

func main() {
  const World = "世界"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Truth = true
  fmt.Println("Go rules?", Truth)
}
```

## Numeric Constants

High-precision values.

An untyped constant takes the type needed by its context.

```go
package main

import "fmt"

const (
  // Create a huge number by shifting a 1 bit left 100 places.
  // In other words, the binary number that is 1 followed by 100 zeroes.
  Big = 1 << 100
  // Shift it right again 99 places, so we end up with 1<<1, or 2.
  Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
  return x * 0.1
}

func main() {
  fmt.Println(needInt(Small))
  fmt.Println(needFloat(Small))
  fmt.Println(needFloat(Big))
  // fmt.Println(needInt(Big))
  // (An int can store at maximum a 64-bit integer, and sometimes less.)
}
```
