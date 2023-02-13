# Readers

`io` pkg specifies `io.Reader` interface.
- Represents read end of a stream of data.
- Go standard library contains many implementations of this interface including:
  - Files
  - Network connections
  - Compressors
  - Ciphers
  - Etc.
- Has `Read` method

```go
func (T) Read(b []byte) (n int, err error)
```

`Read` populates the given byte slice w/ data and returns the number of bytes populated and an error value.
- Returns an `io.EOF` error when the stream ends

Example code creates a `strings.Reader` and consumes the output 8 bytes at a time.

```go
package main

import (
  "fmt"
  "io"
  "strings"
)

func main() {
  r := strings.NewReader("Hello, Reader!")

  b := make([]byte, 8)
  for {
    n, err := r.Read(b)
    fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
    fmt.Printf("b[:n] = %q\n", b[:n])
    if err == io.EOF {
      break
    }
  }
}
```
