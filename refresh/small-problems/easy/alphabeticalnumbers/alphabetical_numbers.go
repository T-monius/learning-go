package alphabeticalnumbers

import (
	"cmp"
	"slices"
)

var alphaFromNum = map[int]string{
	0: "zero",
	1: "one",
	2: "two",
}

func AlphabeticalNumbers(nums []int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		return cmp.Compare(alphaFromNum[a], alphaFromNum[b])
	})

	return nums
}
