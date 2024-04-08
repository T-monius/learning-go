package reverseit

import (
	"strings"
)

func ReverseIt(str string) string {
	words := strings.Split(str, " ")
	reversed := make([]string, 0)

	for idx := len(words) - 1; idx > -1; idx -= 1 {
		word := words[idx]
		reversed = append(reversed, word)
	}

	return strings.Join(reversed, " ")
}
