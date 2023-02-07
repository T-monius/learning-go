# Function Values

First-class functions.

```go
package main

import (
  "fmt"
  "math"
)

func compute(fn func(float64, float64) float64) float64 {
  return fn(3, 4)
}

func main() {
  hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
  }
  fmt.Println(hypot(5, 12))

  fmt.Println(compute(hypot))
  fmt.Println(compute(math.Pow))
}
```

## Function Closures

Go functions may be closures.
- Closure: a function value that references variables from outside its body.
- May access and assign to referenced variables.
  - The function is "bound" to the variables.

For example, the `adder` function returns a closure. Each closure is bound to its own `sum` variable:

```go
package main

import "fmt"

func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return sum
  }
}

func main() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}

/* prints:
0 0
1 -2
3 -6
6 -12
10 -20
15 -30
21 -42
28 -56
36 -72
45 -90 */
```
