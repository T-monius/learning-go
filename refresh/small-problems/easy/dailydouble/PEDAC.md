# Daily Double

## Problem

Write a function that takes a string argument and returns a new string that contains the value of the original string with all consecutive duplicate characters collapsed into a single character. You may not use a built-in that does this automatically.

__Understanding__

Input
- One argument
  - String
  - Arbitrary characters
Output
- String
  - No consecutive duplicates

Return a 'new string' (No 'in-place' requirement)

## Examples/Test Cases
```go
str := "ddaaiillyy ddoouubbllee"
//                   ^
// q = { d a i l y   d o... }
//=> "daily double"
str1 := "4444abcabccba"
//=> "4abcabcba"
str2 := "ggggggggggggggg"
//=> 'g'
str3 := "a"
//=> "a"
str4 := ""
//=> ""
```

## Data Structures

- Queue (empty slice)
- Empty string

## Algorithm

1. Instantiate a `queue` (empty slice) var
1. Iterate over the bytes in the input string
1.   Convert byte to a `string` char
1.   Conditionally push char to `queue`
1. Instantiate an empty string
1. Iterate while the `queue` is not empty
1.   Push current char to string
1. Return string
