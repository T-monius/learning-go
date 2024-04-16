package dailydouble

import "strings"

func SqueezeDuplicates(str string) string {
	nonDupChars := make([]string, 0)

	for _, b := range str {
		char := string(b)
		if len(nonDupChars) == 0 || (nonDupChars[len(nonDupChars)-1] != char) {
			nonDupChars = append(nonDupChars, char)
		}
	}

	return strings.Join(nonDupChars, "")
}
