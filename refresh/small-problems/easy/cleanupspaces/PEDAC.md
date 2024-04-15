# Cleanup Spaces

## Problem

Given a string that consists of some words and an assortment of non-alphabetic characters, write a function that returns that string with all of the non-alphabetic characters replaced by spaces. If one or more non-alphabetic characters occur in a row, you should only have one space in the result (the result should never have consecutive spaces).

Input
- String
  - Alpha and non alpha chars
Output
- String
  - Non-alphas removed

Understanding
- Assume English
- What's an alpha?
  - Letters
- We can define a collection to hold alphas
- ASCII table
  - Each char corresponds to a number on the table

## Examples/Test Cases

```go
messy := "---what's my +*& line?"
//                             ^
// clean := " what s my line "
//=> " what s my line "
```

## Data Structures

- N/A

## Algorithm

1. Define alpha ranges from ASCII table
1. Instantiate a `clean` string var for return
1. Iterate for `i` of input string
1.   Determine if char at iteration is Alpha
       Compare to upper and lower alpha min/max
     If non-alpha
        Push a space to `clean`
        Iterate setting `j` to `i + 1` while `j` is less than string len
          If the char at iteration is alpha, set `i` to `j-1` and break
     Else
       Push char to `clean`
1. Return `clean`

