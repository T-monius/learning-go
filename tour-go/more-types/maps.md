# Maps

Way to map keys to values.
- Zero value `nil`
  - `nil` map has no keys
  - keys can't be added

`make` returns a map of a type initialized for use.

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m map[string]Vertex

func main() {
  m = make(map[string]Vertex)
  m["Bell Labs"] = Vertex{
    40.68433, -74.39967,
  }
  fmt.Println(m["Bell Labs"])
}

//=> {40.68433 -74.39967}
```

## Map literals

Like struct literals
- Keys required

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": Vertex{
    40.68433, -74.39967,
  },
  "Google": Vertex{
    37.42202, -122.08408,
  },
}

func main() {
  fmt.Println(m)
}
```

If the top-level type is just a type name, it can be omitted from the elements of the literal.

```go
package main

import "fmt"

type Vertex struct {
  Lat, Long float64
}

var m = map[string]Vertex{
  "Bell Labs": {40.68433, -74.39967},
  "Google":    {37.42202, -122.08408},
}

func main() {
  fmt.Println(m)
}
```

## Mutating Maps

Insert or update an element in map `m`:

```go
m[key] = elem
```

Retrieve an element:

```go
elem = m[key]
```

Delete an element:

```go
delete(m, key)
```

Test that a key is present with a two-value assignment:
```go
elem, ok = m[key]
```

If `key` is in `m`, `ok` is `true`. If not, `ok` is `false`.

If `key` is not in the map, then `elem` is the zero value for the map's element type.

Note: If `elem` or `ok` have not yet been declared you could use a short declaration form:

```go
elem, ok := m[key]
```

```go
package main

import "fmt"

func main() {
  m := make(map[string]int)

  m["Answer"] = 42
  fmt.Println("The value:", m["Answer"])

  m["Answer"] = 48
  fmt.Println("The value:", m["Answer"])

  delete(m, "Answer")
  fmt.Println("The value:", m["Answer"])

  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}
```
