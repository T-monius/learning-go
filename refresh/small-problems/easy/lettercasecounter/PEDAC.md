# Lettercase Counter

## Problem

Write a function that takes a string, and then returns a map that contains 3 entries: one represents the number of characters in the string that are lowercase letters, one the number of characters that are uppercase letters, and one the number of characters that are neither.

__Understanding__

Input
- String
Output
- Map
  - 3 keys
    1. lowerase letter count, int
    2. uppercase letter count, int
    3. neither count, int

Strings in go are immutable
Strings can be ranged over
Assume that characters are ASCII
Characters are runes in Go, pertinent when iterating and accessing a letter
`int` converts a byte to its ASCII equivalent

## Examples/Test Cases

```go
map[string]int{ "lowercase":5,"uppercase": 1, "neither": 4 }
const UpperMax := 90  //?
const UpperMin := 67  //?
const LowerMin := 91  //?
const LowerMax := 122 //?
LetterCaseCount("abCdef 123")
//                         ^
//=> map[string]int{ "lowercase":5,"uppercase": 1, "neither": 4 }
LetterCaseCount("AbCd +Ef")
//=> map[string]int{ "lowercase":3,"uppercase": 3, "neither": 2 }
LetterCaseCount("123")
//=> map[string]int{ "lowercase":0,"uppercase": 0, "neither": 3 }
LetterCaseCount("")
//=> map[string]int{ "lowercase":0,"uppercase": 0, "neither": 0 }
```

## Data Structure

- Map

## Algorithm

1. Instantiate map for return w/ defaults
1. Iterate over input string
1.   Access char at iteration
1.   convert current character to ASCII w/ `int`
1.   Determine where current ASCII falls w/ conditionals OR swich
1.     Increment given map property
1. Return map
