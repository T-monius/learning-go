package numbertostring

import (
	"fmt"
	"slices"
)

const zero string = "0"
const numericBase int = 10

var digitToChar = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func DigitsFromNumber(num int) []int {
	digits := make([]int, 0)

	for num > 0 {
		dig := num % numericBase
		digits = append(digits, dig)
		num /= numericBase
	}

	return digits
}

func NumberToString(num int) string {
	if num == 0 {
		return zero
	}

	digits := DigitsFromNumber(num)
	var s string = ""

	for len(digits) > 0 {
		lastIdx := len(digits) - 1
		digit := digits[lastIdx]
		s += fmt.Sprint(digit)
		digits = slices.Delete(digits, lastIdx, len(digits))
	}

	return s
}
