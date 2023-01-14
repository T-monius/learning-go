# Functions

Typed language
- Type comes after variable name in function param

```go
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
}
```

[Go's Declaration Syntax](https://go.dev/blog/declaration-syntax)
- Go declaration differs from C family
  - C ex: `int x`
- Go Syntax
  - Non C languages usually declare type differently
  - Name usually comes first (often followed by a colon)
  - Go declares name first and drops colon: `x int`
  - Reading left-to-right clear as type complexity increases
  - Pointers
    - Declaring arrays and slices puts brackets on left while expression syntax puts brackets on right:
    ```go
    var a []int
    x = a[1]
    ```
    - `*` notation from C maintained for pointers:
    ```go
    var p *int
    x = *p
    ```

Two or more consecutive named function parameters can use short-form (omit a type)

```go
x int, y int
// becomes
x, y int
```
