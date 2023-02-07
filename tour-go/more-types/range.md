# Range

The `range` form of the `for` loop iterates over a slice or a map.
- Two values returned per iteration for a slice
  1. Index
  2. Copy of the element at index

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
  for i, v := range pow {
    fmt.Printf("2**%d = %d\n", i, v)
  }
}
```

You can skip the index or value by assigning to `_`.

If you only want the index, you can omit the second variable.
