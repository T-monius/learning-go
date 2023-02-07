# Interfaces

Inrerface type: set of method signatures
- Can hold any value that implements those methods

__Note:__ There is an error in the example code on line 22. `Vertex` (the value type) doesn't implement `Abser` because the `Abs` method is defined only on `*Vertex` (the pointer type).

```go
package main

import (
  "fmt"
  "math"
)

type Abser interface {
  Abs() float64
}

func main() {
  var a Abser
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}

  a = f  // a MyFloat implements Abser
  a = &v // a *Vertex implements Abser

  // In the following line, v is a Vertex (not *Vertex)
  // and does NOT implement Abser.
  a = v

  fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}

type Vertex struct {
  X, Y float64
}

func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
```

## Interfaces are Implemented Implicitly

A type implements an interface by implementing its methods.
- There is no explicit declaration of intent (i.e. 'implements' keyword)

_Implicit interfaces decouple the definition of an interface from its implementation._
- Implementation could then appear in any package w/o prearrangement.

```go
package main

import "fmt"

type I interface {
  M()
}

type T struct {
  S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
  fmt.Println(t.S)
}

func main() {
  var i I = T{"hello"}
  i.M()
}

// => hello
```

## Interface Values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

`(value, type)`

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the methods of the same name on its underlying type.

```go
package main

import (
  "fmt"
  "math"
)

type I interface {
  M()
}

type T struct {
  S string
}

func (t *T) M() {
  fmt.Println(t.S)
}

type F float64

func (f F) M() {
  fmt.Println(f)
}

func main() {
  var i I

  i = &T{"Hello"}
  describe(i)
  i.M()

  i = F(math.Pi)
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

/*
=> (&{Hello}, *main.T)
Hello
(3.141592653589793, main.F)
3.141592653589793 */
```

## Nil interface values

Holds neither value or concrete type.
- Call method on -> runtime error
- No type inside interface tuple indicating _concrete_ method.

```go
package main

import "fmt"

type I interface {
  M()
}

func main() {
  var i I
  describe(i)
  i.M()
}

func describe(i I) {
  fmt.Printf("(%v, %T)\n", i, i)
}

/*
=> (<nil>, <nil>)
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x481961]

goroutine 1 [running]:
main.main()
  /tmp/sandbox3454078585/prog.go:12 +0x61
*/
```

## The empty interface

Interface type that specifies zero methods: _empty interface_
- `interface{}`

An empty interface may hold values of any type.
- Every type implements at least zero methods.

Empty interfaces are used by code that handles values of unknown type.
- For example, `fmt.Print` takes any number of arguments of type `interface{}`.

```go
package main

import "fmt"

func main() {
  var i interface{}
  describe(i)

  i = 42
  describe(i)

  i = "hello"
  describe(i)
}

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}

/*
=> (<nil>, <nil>)
(42, int)
(hello, string)
*/
```

## Type assertions

A _type assertion_ provides access to an interface value's underlying concrete value.

```go
t := i.(T)
```

This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger panic.

_Test_ a specific type on an interface value by accessing two values through the type assertion (i.e. a value and a boolean)

```go
t, ok := i.(T)
```

__Note__: test syntax similar to reading from a map.

```go
package main

import "fmt"

func main() {
  var i interface{} = "hello"

  s := i.(string)
  fmt.Println(s)

  s, ok := i.(string)
  fmt.Println(s, ok)

  f, ok := i.(float64)
  fmt.Println(f, ok)

  f = i.(float64) // panic
  fmt.Println(f)
}

/*
=> hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
  /tmp/sandbox3217719788/prog.go:17 +0x14a
*/
```

## Type switches

_Type switch_: permits several type assertions in series.
- Like regular switch statement.
- Cases specify types (not values)

```go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

```go
package main

import "fmt"

func do(i interface{}) {
  switch v := i.(type) {
  case int:
    fmt.Printf("Twice %v is %v\n", v, v*2)
  case string:
    fmt.Printf("%q is %v bytes long\n", v, len(v))
  default:
    fmt.Printf("I don't know about type %T!\n", v)
  }
}

func main() {
  do(21)
  do("hello")
  do(true)
}
```

## Stringsers

`Stringer`, defined by `fmt`
- Ubiquitous interface

```go
type Stringer interface {
  String() string
}
```

Type that can describe itself as a string.
- Looked for to print values.

```go
package main

import "fmt"

type Person struct {
  Name string
  Age  int
}

func (p Person) String() string {
  return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
  a := Person{"Arthur Dent", 42}
  z := Person{"Zaphod Beeblebrox", 9001}
  fmt.Println(a, z)
}

//=> Arthur Dent (42 years) Zaphod Beeblebrox (9001 years)
```

## Stringers Exercise

Make the `IPAddr` type implement fmt.Stringer to print the address as a dotted quad.

For instance, `IPAddr{1, 2, 3, 4}` should print as `"1.2.3.4"`.

```go
package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (i IPAddr) String() string {
  return fmt.Sprintf("%v.%v.%v.%v", i[0], i[1], i[2], i[3])
}

func main() {
  hosts := map[string]IPAddr{
    "loopback":  {127, 0, 0, 1},
    "googleDNS": {8, 8, 8, 8},
  }
  for name, ip := range hosts {
    fmt.Printf("%v: %v\n", name, ip)
  }
}

/*
=> loopback: 127.0.0.1
googleDNS: 8.8.8.8 */
```
