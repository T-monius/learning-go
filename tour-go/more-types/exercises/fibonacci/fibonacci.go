package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	twoPreviousNum, previousNum := -1, -1
	return func() int {
		if previousNum == -1 {
			previousNum = 0
			return previousNum
		}
		if twoPreviousNum == -1 {
			twoPreviousNum = 0
			previousNum = 1
			return previousNum
		}

		currentNum := previousNum + twoPreviousNum
		twoPreviousNum = previousNum
		previousNum = currentNum

		return currentNum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}