# Combine Slices

## Problem

Write a function that takes two Arrays/Slices as arguments, and returns an Array/Slice that contains all of the values from the argument Arrays/Slices. There should be no duplication of values in the returned Array, even if there are duplicates in the original Arrays.

### Understanding

Input
- 2 arguments
  - Arrays/Slices
    - Element type?
Output
- Array/Slice
  - Contains non-duplicate elements of inputs

No duplication
- Are the inputs ordered?

## Examples/Test Cases

```go
sl := []int{1, 3, 5}
//                  i
sl1 := []int{3, 6, 9}
//                 j
merge(sl, sl1)
//=> [1, 3, 5, 6, 9]

sl := []int{1, 7, 5}
// sorted := []int{1, 5, 7}
//                         i
sl1 := []int{3, 6, 7, 9}
// sl1 := []int{3, 6, 7, 9}
//                       j
merge(sl, sl1)
//=> [1, 3, 5, 6, 7, 9]
```

## Data Structures

- Slice

## Algorithm

1. Instantiate a new slice for return
1. Instantiate `i` and `j` to `0` for iteration
1. Sort input arrays
1. Iterate while both `i` and `j` are less than respective slices
1.   access values at `i` and `j`
1.   If value at `i` is less than that at `j`
1.     Push it to return array
1.     increment `i`
1.   If value at `j` is less than that at `i`
1.     push it to return array
1.     increment `j`
1.   If values are equal
1.     push value at `i`
1.     increment both pointers
1. Instantiate a slice var `restOfSlice`
1. Instantiate var `restOfPosition`
1. Conditionally set `restOfSlice`, `restOfPosition` based on pointer positions of `i` and `j`
1. Iterate over `restOfSlice`
1.   Access last val from return slice
1.   Conditionally push current val to return slice
1. Return return slice
