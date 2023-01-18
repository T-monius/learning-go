# Word Count

## Problem

Implement `WordCount`. It should return a map of the counts of each "word" in the string `s`. The `wc.Test` function runs a test suite agains the provided function and prints success or Failure.

You might find `strings.Fields` helpful.

## Understanding

Input
- String
Output
- Map
  - keys: each word in string
  - values: word count

Words are any consecutive non-whitespace chars

## Examples / Test Cases

```go
var str string = "I am learning Go!"
var wc = make(map[string]int)

// wc =>
//      map[string]int{"I": 1}
// words =>
// a[4]string{"I", "am", "learning", "Go!"}
//                   ^


// output =>
// map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
```

## Data Structures

- Array/slice
- Map

## Algorithm

- Split `s` into an array/slice (w/ `string.Fields`)
- Instantiate a map w/ string keys and in vals
- Iterate array/slice `words`
  - Instantiate `elem`, `ok` vars from `wc` access for current `word`
  - If `ok`
    - Access value
    - Re-assign value to value plus one
  - Else
    - Add key w/ `word` assigning one

- Return `wc`
