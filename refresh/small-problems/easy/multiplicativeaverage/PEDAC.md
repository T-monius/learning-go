# Multiplicative Average

## Problem

Write a function that takes an Array/Slice of integers as input, multiplies all the numbers together, divides the result by the number of entries in the Array/Slice, and then prints the result rounded to 3 decimal places. Assume the array is non-empty.

__Understanding__

Input
- Array/Slice
Output
- String
  - Average to three decimal places

Multiply input nums
Divide by length of input
Format result to a string

## Examples/Test Cases

```go
nums := []int{3, 5}
MultiplicativeAverage(nums)
//=>  "7.500"

nums1 := []int{6}
MultiplicativeAverage(nums1)
//=> "6.000"

nums2 := []int{2, 5, 7, 11, 13, 17}
MultiplicativeAverage(nums2)
//=> "28361.667"
```

## Data Structures

- N/A

## Algorithm

1. Declare `total` var for nums total
1. Iterate `nums` adding `num` at iteration to total
1. Convert `total` to a float and divide by `nums.length` setting equal to `avg` var
1. Format `avg` to a string w/ `fmt.Sprintf`?
1. Return formatted avg
