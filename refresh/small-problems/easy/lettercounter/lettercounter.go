package lettercounter

import "strings"

const space string = " "

func LetterCounter(str string) map[int]int {
	lc := make(map[int]int)
	if len(str) < 1 {
		return lc
	}
	words := strings.Split(str, space)

	for _, word := range words {
		l := len(word)
		_, ok := lc[l]
		if !ok {
			lc[l] = 1
			continue
		}
		lc[l] += 1
	}
	return lc
}
