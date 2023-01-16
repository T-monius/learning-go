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
