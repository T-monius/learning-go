# Fibonacci Closure

## Problem

Let's have some fun with functions.

Implement a `fibonacci` function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

### Understanding

A Fibonnaci number is the result of the addition of the previous two numbers in the series.
- The series starts w/ zero and one.
- Therefore, the third number in the series is 1 (i.e. 0 + 1 => 2)

The loop invoking the `f` (e.g. `fibonacci`) function does not pass it any arguments.
- Not similar to the recursive solution which would unfold all previous numbers in the series to calculate the nth number, argument passed.

Use a closure.
- Two variables like `previousNum` and `numTwoPreious` can be used to calculate and return `currentNum`
- First two invocations are exceptional
  - Can instantiate closure vars w/ `-1`

## Examples / Test Cases

```go
previousNum    //=> -1
numTwoPrevious //=> -1


```

## Data Structures

- n/a (closure?)

## Algorithm

Instantiate `previousNum` and `numTwoPrevious`
Return a function that
  If `previuosNum` is `-1`, sets it to `0` and return it
  If `twoPreviousNum` is `-1`, sets it to `0`, set `previousNum` to `1`, and return it.

  Sets `currentNum` to the sum of two closure vars
  Sets `numTwoPrevious` to `previousNum`
  Sets `previousNum` to `currentNum`
  Returns `currentNum`
