package listinclude

import (
	"slices"
)

func DoesInclude(nums []int, target int) bool {
	slices.Sort(nums)
	l := 0
	r := len(nums) - 1
	for l < r-1 {
		m := (r + l) / 2
		val := nums[m]

		switch {
		case val == target:
			return true
		case val < target:
			l = m
		case val > target:
			r = m
		}
	}

	if nums[l] == target || nums[r] == target {
		return true
	}

	return false
}
