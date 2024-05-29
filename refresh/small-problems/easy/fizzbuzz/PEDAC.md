# FizzBuzz

## Problem

Write a function that takes two arguments: the first is the starting number, and the second is the ending number. Print out all numbers between the two numbers, except if a number is divisible by 3, print "Fizz", if a number is divisible by 5, print "Buzz", and finally if a number is divisible by 3 and 5, print "FizzBuzz".

__Understanding__

Input
- 2 arguments
  - number, start
  - number, end
Output
- Print
  - All numbers between start and end inclusive
  - Exceptions:
    - divisible by 3, print "Fizz"
    - divisible by 5, print "Buzz"
    - divisibly by 3 and 5, "FizzBuzz

## Examples/Test Cases

```go
FizzBuzz(1, 15)
//=> "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz"
FizzBuzz(1, 100)
//=> "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz, ..."
```

## Data Structures

## Algorithm

1. Instantiate slice `nums`
1. Assign `num` to `start`
1. Iterate from `start` until `end`
1.   Condionally push `string(num)` or phrase to `nums`
1. Join `nums` with `", "` and return
