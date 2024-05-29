package fizzbuzz

import (
	"fmt"
	"strings"
)

const sequenceSeparator string = ", "
const primaryDividend int = 3
const secondaryDividend int = 5
const primaryVal string = "Fizz"
const secondaryVal string = "Buzz"

func isDivisibleBy(num, div int) bool {
	return num%div == 0
}

func determineNextVal(num int) string {
	var nextVal string
	switch {
	case (isDivisibleBy(num, primaryDividend)) && (isDivisibleBy(num, secondaryDividend)):
		nextVal = fmt.Sprintf("%s%s", primaryVal, secondaryVal)
	case (isDivisibleBy(num, primaryDividend)):
		nextVal = primaryVal
	case (isDivisibleBy(num, secondaryDividend)):
		nextVal = secondaryVal
	default:
		nextVal = fmt.Sprint(num)
	}
	return nextVal
}

func FizzBuzz(start, end int) string {
	var sequence []string
	for num := start; num <= end; num += 1 {
		nextVal := determineNextVal(num)
		sequence = append(sequence, nextVal)
	}
	return strings.Join(sequence, sequenceSeparator)
}
