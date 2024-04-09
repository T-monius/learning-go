# Number to String

## Problem

Write a function that takes a positive integer or zero, and converts it to a string representation.

Input
- 1 argument
- Positive int OR zero
Output
- String
  - Represents the digits of the number

__Understanding__

- Decimal num (base 10)
- Can modulus off positions of a number given a base
- Can map digits to string values

## Examples/Test Cases

```go
num := 4321
NumToString(num)
//=> "4321"
//    ^
// stack { }
//         ^
// str = "4321"
// return str
num1 := 0
NumToString(num1)
//=> "0"
NumToString(num2)
num2 := 5000
//=> "5000"
```

## Data Structures

- Stack (empty slice)
- Empty str

## Algorithm

1. If the number is zero return `"0"`
1. Instantiate a `digits` stack
1. Iterate while the input int is greater than 0
1.   Modulus off last digit (i.e. mod 10)
1.   Push new dig to stack
1.   Divive input num by 10 and re-assign
1. Instantiate a `s` var for return
1. Iterate while stack isn't empty
1.   Pop last digit
1.   Convert digit to char and push to `s`
1. Return `s`
