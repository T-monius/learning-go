# List Include

## Problem

Write a fucntion named `DoesInclude` that takes a Slice of Ints and a search value as arguments. This fucntion should return `true` if the search value is in the array, `false` if it is not. You may not use `slices` functions that do this in your solution.

__Understanding__

Input
- Slice
  - Int elements
  - Arbitrary order
Output
- Bool

Determine whether a particular element is in the slice

Arbitrary order

## Examples/Test Cases

```go
nums := []int{1,2,3,4,5}
target := 3
//=> true
nums1 := []int{1,3,2,4,5}
//                   l
//                   m
//                      r
target1 := 6
//=> false
nums := []int{}
target2 := 3
//=> false
```

## Data Structures

- Slice

## Algorithm

1. Sort input array
1. Set `l` and `r` to slice bounds
1. While `l < r`
1.   Calculate `m` (`r + l / 2`)
1.   If `m` is `target` early return it
1.   Else if `m < target`
1.     set `l` to `m`
1.   Else if `m > target`
1.     set `r` to `m`
1. If `l` or `r` is `target` return `true`
1. return `false`