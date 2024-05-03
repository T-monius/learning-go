# Fibonacci Location by Length

## Problem

The Fibonacci series is a series of numbers (1, 1, 2, 3, 5, 8, 13, 21, ...) such that the first 2 numbers are 1 by definition, and each subsequent number is the sum of the two previous numbers. This series appears throughout the natural world.

Computationally, the Fibonacci series is a very simple series, but the results grow at an incredibly rapid rate. For example, the 100th Fibonacci number is 354,224,848,179,261,915,075 -- that's enormous, especially considering that it takes 6 iterations before it generates the first 2 digit number.

--

Write a method that calculates and returns the index of the first Fibonacci number that has the number of digits specified as an argument. (The first Fibonacci number has index 1.)

__understanding__

Calculate fibonacci numbers
- Maintain a count/index for return
- Need previous two numbers to solve for the given index

Determine the length of a number
- Cast to string
- Radix math (lenght changes at every power of 10)

## Example/Test Cases

```go
FindFibonacciIndexByLength(2)
//=> 7
prior := 1
penultimate := 1
i := 3
current := 2
  prior := 2
  penultimate := 1
  i := 4
  current := 3 
    prior := 3
    penultimate := 2
    i := 5
    current := 5
      prior := 5
      penultimate := 3
      i := 6
      current := 8
        prior := 8
        penultimate := 5
        i := 7
        current := 13
FindFibonacciIndexByLength(10)
//= 45
FindFibonacciIndexByLength(20)
//= 93
// Below are too big for recursion
FindFibonacciIndexByLength(100)
//=> 476
FindFibonacciIndexByLength(1000)
//=> 4782
FindFibonacciIndexByLength(10000)
//=> 47847
```

## Data Structures

- ...

## Algorithm

1. Sum `prior` and `penultimate` to calculate `current`
1. Conditionally return `i` if `current` is length `l`
1. Guard clause based on length of `l`
1. Invoke recusion with `prior` as `penultimate` and `current` as prior. Increment `i`.
