# Double Consonants

## Problem

Write a function that takes a string, and returns a new string in which every consonant character is doubled. Vowels (a,e,i,o,u), digits, punctuation, and whitespace should not be doubled.

__Understanding__

Input
- String
Output
- String
  - Consonants of original string doubled

## Examples/Test Cases

```go
str := "String"
//         ^
doubled := 'SSttrri...'
//=> "SSttrrinngg"
str1 := "Hello-World!"
//=> "HHellllo-WWorrlldd!"
str2 := "July 4th"
//=> "JJullyy 4tthh"
str3 := ""
//=> ""
```

## Data Structures

- Map

## Algorithm

1. Create `CONSONANTS` lookup map
1. Instantiate `doubled` string for return
1. Iterate over input `str`
1.   Convert rune/byte to `string`
1.   Loodup current char in consonant map
1.   Conditionally push twice
1.   Else push once
1. Return `doubled`
