package stringystrings

import (
	"strings"
)

func StringyStrings(num int) string {
	binaryDigs := make([]string, num)

	for idx := 0; idx < num; idx += 1 {
		if idx%2 == 0 {
			binaryDigs = append(binaryDigs, "1")
		} else {
			binaryDigs = append(binaryDigs, "0")
		}
	}

	binaryStr := strings.Join(binaryDigs, "")

	return binaryStr
}
