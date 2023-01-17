# Flow Control

## Looping Constructs

`for` loop
- Only looping construct
- 3 components separated by semicolons (no parens):
  1. init statement: executed before first iteration
    - Often short variable declaration
    - Variables declared only in scope of `for` statement
  2. condition expression: evaluated before every iteration
    - Loop will stop iterating once boolean condition evaluates to `false`
  3. post statement: executed at the end of every iteration

```go
package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }
  fmt.Println(sum)
}
```

init and post statements are optional

```go
package main

import "fmt"

func main() {
  sum := 1
  for ; sum < 1000; {
    sum += sum
  }
  fmt.Println(sum)
}
```

Can drop semicolons when only using a condition expression which is the equivalent of a `while` loop.

```go
package main

import "fmt"

func main() {
  sum := 1
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)
}
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```go
package main

func main() {
  for {
  }
}
```

## `if` statements

Like `for` loops.
- The expression need not be surrounded by parentheses `()`.
- Braces `{}` are requires.

```go
package main

import (
  "fmt"
  "math"
)

func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func main() {
  fmt.Println(sqrt(2), sqrt(-4))
}
```

`if` statement can also start w/ a short statement to execute before the condition.
- Variables declared only in scope of the `if` block or any `else` blocks

```go
package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim // v out of scope
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
```

```go
package main

import (
  "fmt"
  "math"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}

func main() {
  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )
}
```

## Switch Statement

`switch`: shorter way of writing a sequence of `if - else` statements.
- Runs the first case whose value is equal to the condition expression.

Go only runs the selected case, not all the cases that follow.

Switch cases need not be constants, and the values need not be integers.

```go
package main

import (
  "fmt"
  "runtime"
)

func main() {
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Println("OS X.")
  case "linux":
    fmt.Println("Linux.")
  default:
    // freebsd, openbsd,
    // plan9, windows...
    fmt.Printf("%s.\n", os)
  }
}
```

### Switch evaluation order

Cases evaluate from top to bottom.

### Switch with no condition

Same as `switch true`
- This construct can be a clean way to write long if-then-else chains.

```go
package main

import (
  "fmt"
  "time"
)

func main() {
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon.")
  default:
    fmt.Println("Good evening.")
  }
}
```
