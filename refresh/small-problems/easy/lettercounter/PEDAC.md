# Letter Counter

## Problem

Write a function that takes a string with one or more space separated words and returns a map that shows the number of words of different sizes.

_NOTE:_ Words consist of any string of characters that do not include a space.

__Understanding__

Input
- One arg
  - String
    - One or more space separated words
    - Word: non-space char sequence
      - Assume:
        - `const space := " "`
        - Not considering tabs, line sep, etc.
Output
- Map
  - Key: int => word length
  - Val: int => count of words w/ that word len

## Examples/Test Cases

```go
const space := " "
count := map[int]int{  }
words := make([]strings, 0)

str := "Four score and seven."
//                             ^
//                            ^
// words -> { "Four" "score" "and" "seven."}
//                                   ^
// count { 4:1  5:1 3:1 6:1 }
//=> { 3: 1, 4: 1, 5: 1, 6: 1 }
str1 := "Hey diddle diddle, the cat and the fiddle!"
//=> { 3: 5, 6: 1, 7: 2 }
str2 := "What's up doc?"
//=> { 6: 1, 2: 1, 4: 1 }
emptyStr := ""
//=> {}
```

## Data Structures

- Map
- Slice

## Algorithm

1. Assign space char to a constant var
1. Instantiate a map for counts
1. Split input by space -> words?
1. Iterate over words
1.   Determine word len
1.   Conditionally add key for len OR increment count
1. Return count map
