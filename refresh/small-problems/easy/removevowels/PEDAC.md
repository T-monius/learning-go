# Remove Vowels

## Problem

Write a function that takes an slice of strings, and returns an slice of the same string values, except with the vowels (a, e, i, o, u) removed.

__Understanding__

Input
- Slice
  - String elements
  - Strings have characters/bytes
Output
- Slice
  - String elments
  - Removed vowel characters

Operate on each element
- Iteratitively
- Recursively

While operating on each element
- Operate on each char of the string

## Examples/Test Cases

```go
RemoveVowels([]string{"abcdefghijklmnopqrstuvwxyz"})
//          ^
//=> []string{"bcdfghjklmnpqrstvwxyz"}
// "bcdfghjklmnpqrstvwxyz"
//                       w
//                            r
RemoveVowels([]string{"green", "YELLOW", "black", "white"})
//=> []string{"grn", "YLLW", "blck", "wht"}
RemoveVowels([]string{"ABC", "AEIOU", "XYZ"})
//=> []string{"BC", "", "XYZ"}
```

## Data Structures

- ...

## Algorithm

1. Instantiate a slice for return
1. Iterate over strings slice
1. At each iteration
1.   Convert `str` to a byte array
1.   Instantiate `w` and `r` pointers to `0`
1.   Iterate over the byte array while `r` less than the length of the byte array
1.     If char at `r` is not a vowel, assign it to `w` in byte array and increment `w`
1.     increment `r`
1.   Slice byte array up to `w`, convert to string, and push to return slice
1. Return return slice
