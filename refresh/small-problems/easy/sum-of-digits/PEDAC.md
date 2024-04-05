# Sum of Digits

## Problem

Write a method that takes one argument, a positive integer, and 
returns the sum of its digits.

## Examples/Test Cases

num = 23
Output = 5

total = 0

dig = 23 % 10
dig = 3
total += 3

num =/ 10
//=> 2.3

dig = 2.3 % 10
dig = 2

num =/ 10
//=> .23

num = 496
Output = 19

num = 123_456_789
Output = 45

## Data Structures

- N/A

## Algorithm

1. Set a `total` var to zero
1. Iterate while num is greater than 0
1.   Modulus off the last digit with remainder 10
1.   Add `dig` to total
1.   Set `num` to 10 division (w/ `math.Floor`?)
1. Return `total`
