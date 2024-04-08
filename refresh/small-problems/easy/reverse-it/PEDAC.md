# Reverse It

## Problem

Write a function that takes one argument, a string, and returns the same string with the words in reverse order.

_Input_
- One argument
- String
_Output_
- String
  - Same as input, _in-place_?
  - Words in reverse order as input
    - Words v. letters

__Understanding__
- Identify words in the string first
  - There's a `strings.split()` method
  - Manually separate/parse words
- No explicit requirement to do manipluation in-place

## Examples/Test Cases

```go
str := ""
//=> "" just return empty?

stack = ["World", "Hello"]
str1 := "Hello World"
//=> "World Hello"
["World", "Hello"]
    ^
str2 := "Reverse these words"
//=> "words these Reverse"
//         s
//              e
```

## Data Structures

- Stack

## Algorithnm

1. Split string by spaces
2. Instantiate an empty slice as stack
3. Instantiate a slice for return
3. Iterate while stack isn't empty
4.   Push to a new slice
5. Join new return slice w/ `" "`

1. Instantiate an empty slice for return
2. Splic string by spaces
3. Iterate in reverse over split string
4.   Push words to slice for return
5. Join slice for return
