# Pointer receivers

Can declare methods w/ pointer receivers.
- Receiver type has the literal syntax `*T` for some `T`
- `T` itself cannot be a reciever

Methods w/ pointer receivers can modify the value to which the receiver points.
- Methods often need to modify their receiver.
- Pointer receivers are more common than value receivers.

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

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
}

//=> 50
```

`Scale` defined on `*Vertex` allows modifying the reciever `v`
- Removing the `*` prefixed to `&Vertex` in the `Scale` method definition would cause the program to return `5` since the function `Scale` would not be able to modify `v`.

## Pointers and functions

`Abs` and `Scale` methods rewritten as functions.

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

func Scale(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  Scale(&v, 10)
  fmt.Println(Abs(v))
}

//=> 50
```

Removing `*` in `Scale` definition and `&` in `Scale` invocation would allow program to return `5` but not execute the intent of the program which would be to 'scale' `v`.

### Methods and pointer indirection

Comparing the previous two programs, you might notice that functions with a pointer argument must take a pointer:

```go
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

while methods with pointer receivers take either a value or a pointer as the receiver when they are called:

```go
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

For the statement `v.Scale(5)`, even though `v` is a value and not a pointer, the method with the pointer receiver is called automatically. That is, as a convenience, Go interprets the statement `v.Scale(5)` as `(&v).Scale(5)` since the `Scale` method has a pointer receiver.

```go
package main

import "fmt"

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(2)
  ScaleFunc(&v, 10)

  p := &Vertex{4, 3}
  p.Scale(3)
  ScaleFunc(p, 8)

  fmt.Println(v, p)
}

//=> {60 80} &{96 72}
```

The equivalent thing happens in the reverse direction.

Functions that take a value argument must take a value of that specific type:

```go
var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

while methods with value receivers take either a value or a pointer as the receiver when they are called:

```go
var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

In this case, the method call `p.Abs()` is interpreted as `(*p).Abs()`.

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

func AbsFunc(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
  fmt.Println(AbsFunc(v))

  p := &Vertex{4, 3}
  fmt.Println(p.Abs())
  fmt.Println(AbsFunc(*p))
}
/*
=> 5
5
5
5 */
```

## Choosing a value or pointer receiver

Two reasons to use a pointer receiver.

1. So that the method can modify the value that its receiver points to.
2. To avoid copying the value on each method call.
- Can be more efficient if the receiver is a large struct, for example.

In this example, both `Scale` and `Abs` are with receiver type `*Vertex`, even though the `Abs` method needn't modify its receiver:

```go
package main

import (
  "fmt"
  "math"
)

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := &Vertex{3, 4}
  fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
  v.Scale(5)
  fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}

/*
=>
Before scaling: &{X:3 Y:4}, Abs: 5
After scaling: &{X:15 Y:20}, Abs: 25 */
```

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
