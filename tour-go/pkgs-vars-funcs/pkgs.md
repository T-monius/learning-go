# Packages

## Packages

- Make up Go programs

`main`
- Startoint of programs

Package name
- Same as last 'element' of import path
- Ex, `math/rand` files begin w/ `package rand`

NOTE: environment is deterministic (`rand.Intn` returns same number)

## Imports

Imports can be grouped (parenthesized/factored) or written individually.

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
```

Good style to factor.

## Exported names

Name exported if it begins w/ a capital letter
- `Pizza` or `Pi`
Refer to export names to import

```go
package main

import (
  "fmt"
  "math"
)

func main() {
  fmt.Println(math.Pi)
}
```
