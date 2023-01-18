# Slices

## Problem

Implement `Pic`. It should return a slice of length `dy`, each element of which is a slice of `dx` 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include `(x+y)/2`, `x*y`, and `x^y`.

(You need to use a loop to allocate each `[]uint8` inside the `[][]uint8`.)

(Use `uint8(intValue)` to convert between types.)

### Understanding

Input
- 2 parameters
  - Both `int`
  - `dy`
    - Designates the size of output slice
  - `dx`
    - Designates length of inner slices
Output
- Array/Slice
  - Slices?
    - `uint8` elements

## Examples / Test Cases

```go
var dy int = 4
var dx int = 7

[
  [uint8, uint8, uint8, uint8, uint8, uint8, uint8, uint8],
  [uint8, uint8, uint8, uint8, uint8, uint8, uint8, uint8],
  [uint8, uint8, uint8, uint8, uint8, uint8, uint8, uint8],
  [uint8, uint8, uint8, uint8, uint8, uint8, uint8, uint8]
]
```

## Algorithm

Create a slice of length `dy`
Range over the slice with an index
  For index make new slice of length `dx`
  Range over inner slice
    Assign a calculated value to inner slice at index

Return outer slice
