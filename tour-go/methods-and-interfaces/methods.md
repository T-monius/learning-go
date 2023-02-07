# Methods

Go does not have classes
- Methods can be defined on types.

Method: function with a special `receiver` argument
- Receiver appears in its own argument list
  - Between the `func` keyword and method name

In this example, the `Abs` method has a receiver of type Vertex named `v`:

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
}
```

`Abs` written as a regular function with no change in functionality:

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(Abs(v))
}
```

Methods can be declared on non-struct types.

Can only declare a method with a receiver whose type is defined in the same pachage as the method.

```go
package main

import (
  "fmt"
  "math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

func main() {
  f := MyFloat(-math.Sqrt2)
  fmt.Println(f.Abs())
}
```
