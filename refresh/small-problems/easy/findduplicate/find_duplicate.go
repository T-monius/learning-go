package findduplicate

import "math"

var negativeInfinity = int(math.Inf(-1))

func FindDuplicate(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		_, ok := seen[num]
		if ok {
			return num
		}
		seen[num] = true
	}

	return negativeInfinity
}
