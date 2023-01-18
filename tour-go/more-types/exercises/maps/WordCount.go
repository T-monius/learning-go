package main

import (
	"fmt"
	"strings"
	// "golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	// fmt.Printf("words: %v, len: %d, cap: %d\n", words, len(words), cap(words))
	var wc = make(map[string]int)

	for _, word := range(words) {
		count := wc[word]
		wc[word] = count + 1
	}

	fmt.Println(wc)
	return wc
}

// var testString = "I am learning Go!"
var testStringOne = "I am learning Go, and I like it!"

func main() {
	WordCount(testStringOne)
	// wc.Test(WordCount)
}
