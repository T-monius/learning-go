# Stringy Strings

## Problem

Write a method that takes one argument, a positive integer, and returns a string of alternating 1s and 0s, always starting with 1

_NOTE:_ The length of the string should match the given integer.

__Understanding__

Input
- Integer
Output
- String
  - alternating `1`s and `0`s
  - Starting w/ `1`
  - Length of input num
  - Even idx: even
  - Odd idx: zero

## Examples/Test Cases
```go
num = 6
//=> "101010"

num = 9
//=> "101010101"

num = 4
//=> "1010"

num = 7
//=> "1010101"
```

## Data Structures

## Algorithm

1. ~~Declare string var for return~~
1. Declare an idx var starting at 0
1. Iterate for times input num reresents w/ an index
1.   If idx is even, push `"1"` to str else `"0"`
1.   Increment the idx
1. ~~Return string~~
