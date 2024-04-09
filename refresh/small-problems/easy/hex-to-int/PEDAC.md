# Hexadecimal to Integer

## Problem

Write a `HexadecimalToInteger` function that converts a string representing a hexadecimal number to its integer value.

__Understanding__

- Hex chars reps 0-16 in roman?
  - 0-10, a-f
- Each additional character of a hexadecimal number is equivalent to 16 to the power of it's distance from the end of the string.

## Examples/Test Cases

```go
total = 0
num := "123"
//       ^
// 
// diff betw idx and len = 1 || pos -1
// 2 || pos -2
//      13, 1, 12, 14
str := "D1CE"
//        ^
//     (16**3) => 4096 * 13 => 53248
//     (16**2) => 256 * 1 =>     256
//     (16**1) => 16 * 12 =>     192
//                                14
// + 3328
//   
//=> 53710

str1 := "2E6"
//=> 742
```

## Data Structures

- Struct mapping char to roman value

## Algorithm

1. Instantiate a roman total var for return
1. Iterate over input str
1.   determine diff from current idx to end (len - 1)
1.   determine roman value of char at idx
1.   Detemine hex power (i.e. (16**diff))
1.   Multiply val hex power && add to total
1. Return total
