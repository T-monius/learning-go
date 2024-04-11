# ASCII String Value

## Problem

Write a function that determines and returns the ASCII string value of a string that is passed in as an argument. The ASCII string value is the sum of the ASCII values of every character in the string. (You may use String#ord to determine the ASCII value of a character.)

__Understanding__

Input
- Single arg
  - String
Output
- Integer
  - Represents sum
  - Corresponds to values of individual chars in str

## Examples/Test Cases

```go

var sum int
str := "Four score"
//=> 984
str1 := "Launch School"
//=> 1251
str2 := "a"
//       ^
//=> 97
str3 := ""
//=> 0
str4 := "ab"
//        ^
//=> 195
```

## Data Structures

- n/a

## Algorithm

High Level

1. Instantiate a `sum` var
1. Iterate over the str
1.  Convert char at iteration to ASCII
1.  Add val to total
1. Return total

