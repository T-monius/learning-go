package main

import "fmt"

func RepeatYourself(word string, times int) {
	for idx := 0; idx < times; idx += 1 {
		fmt.Printf("%s\n", word)
	}
}

func main() {
	word := "Hello"
	times := 3
	RepeatYourself(word, times)
}
