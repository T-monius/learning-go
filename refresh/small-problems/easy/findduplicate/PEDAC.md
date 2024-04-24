# Find Duplicate

## Problem

Given an unordered array/slice and the information that exactly one value in the array occurs twice (every other value occurs exactly once), how would you determine which value occurs twice?

Write a funtion that will find and return the duplicate value that is known to be in the array.

__Understanding__

Input
- Array/slice
  - Unordered
  - Elements __all__ appear once except __one__
Output
- Single value
  - Duplicate

Determine duplicate value

## Examples/Test Cases

```go
//      []int{1, 1, 3, 5}
// seen := map[int]bool
nums := []int{1, 5, 3, 1}
//                     ^
//=> 1
moreNums := []int{18,  9, 36, 96, 31, 19, 54, 75, 42, 15,
          38, 25, 97, 92, 46, 69, 91, 59, 53, 27,
          14, 61, 90, 81,  8, 63, 95, 99, 30, 65,
          78, 76, 48, 16, 93, 77, 52, 49, 37, 29,
          89, 10, 84,  1, 47, 68, 12, 33, 86, 60,
          41, 44, 83, 35, 94, 73, 98,  3, 64, 82,
          55, 79, 80, 21, 39, 72, 13, 50,  6, 70,
          85, 87, 51, 17, 66, 20, 28, 26,  2, 22,
          40, 23, 71, 62, 73, 32, 43, 24,  4, 56,
          7,  34, 57, 74, 45, 11, 88, 67,  5, 58}
//=> 73
```

## Data Structures

- Map

## Algorithm

1. Instantiate a `seen` map
1. Iterate over the input array
1.   Check `seen` for int at iteration
1.   Conditionally return that int else add to `seen`
1. Return `-Infinity`?
