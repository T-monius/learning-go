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
## Return Values

Multiple return valu
- Arbitrary number of values can be returned.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
  return y, x
}

func main() {
  a, b := swap("hello", "world")
  fmt.Println(a, b)
}
```

Named Return values
- Return values can be named
  - Treated as variables defined at top of function
  - __USE__: document the meaning of return values
- 'Naked return': `return` statement w/o args
  - Returns named return values
  - Best practice, use only w/ short functions

```go
package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(32))
}
```
