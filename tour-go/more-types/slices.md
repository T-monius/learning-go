# Slices

A slice is a dynamically-sized, flexible view into the elements of an array.

the type `[]T` is a slice with elements of type `T`.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

`a[low : high]`

This selects a half-open range which includes the first element but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of `a`:

`a[1:4]`

```go
package main

import "fmt"

func main() {
  primes := [6]int{2, 3, 5, 7, 11, 13}

  var s []int = primes[1:4]
  fmt.Println(s)
}
```

## Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

Other slices that share the same underlying array will see those changes.

## Slice literals

A slice literal is like an array literal without the length.

This is an array literal:

`[3]bool{true, true, false}`

And this creates the same array as above, then builds a slice that references it:

`[]bool{true, true, false}`

```go
package main

import "fmt"

func main() {
  q := []int{2, 3, 5, 7, 11, 13}
  fmt.Println(q)

  r := []bool{true, false, true, true, false, true}
  fmt.Println(r)

  s := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)
}

/* prints
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}] */
```

## Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

`var a [10]int`

these slice expressions are equivalent:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

## Slice length and capacity

A slice has both a _length_ and a _capacity_.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions _len(s)_ and _cap(s)_.

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.

## Nil slices

The zero value of a slice is `nil`.

A nil slice has a length and capacity of 0 and has no underlying array.

```go
package main

import "fmt"

func main() {
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil!")
  }
}
```

## Creating a slice with make

`make` function
- Built-in
- Manner of creating dynamically-sized arrays.

`make` allocates a zeroed array and returns a slice that refers to that array

```go
a := make([]int, 5) // len(a)=5
```

Specify a capacity passing a third argument:

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

## Slices of slices

Slices can contain any type, including other slices.

```go
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Create a tic-tac-toe board.
  board := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
  }

  // The players take turns.
  board[0][0] = "X"
  board[2][2] = "O"
  board[1][2] = "X"
  board[1][0] = "O"
  board[0][2] = "X"

  for i := 0; i < len(board); i++ {
    fmt.Printf("%s\n", strings.Join(board[i], " "))
  }
```

## Appending to a slice

`append` function ([docs](https://go.dev/pkg/builtin/#append))
