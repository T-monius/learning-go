# Alphabetical Numbers

## Problem

Write a function that takes an Array of Integers between 0 and 19, and returns an Array of those Integers sorted based on the English words for each number:

zero, one, two, three, four, five, six, seven, eight, nine, ten, eleven, twelve, thirteen, fourteen, fifteen, sixteen, seventeen,eighteen, nineteen

_Understanding_

Input
- Array/slice?
  - Integer elements
  - between zero and 19
Output
- Array
  - Integers
    - Sorted
      - by English words
      - Ascending

## Examples/Test Cases

```go
alphaFromNums := map[int]string{
  0: "zero",
  1: "one",
  // ...
}
nums := [0, 1, 2, 3,..19]
//       ^
//=>
[
  8, 18, 11, 15, 5, 4, 14, 9, 19, 1, 7, 17,
  6, 16, 10, 13, 3, 12, 2, 0
]
```

## Data Structures

- N/A

## Algorithm

1. Instantiate a map of nums to alphas
1. Invoke sort function
1.   At each iteration of sort, access char
1. Return sorted array
