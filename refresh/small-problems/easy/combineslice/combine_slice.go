package combineslice

import (
	"slices"
)

func Merge(ints, otherInts []int) []int {
	merged := make([]int, 0)
	slices.Sort(ints)
	slices.Sort(otherInts)
	i := 0
	j := 0
	for i < len(ints) && j < len(otherInts) {
		val := ints[i]
		otherVal := otherInts[j]
		if val <= otherVal {
			merged = append(merged, val)
		}
		if otherVal < val {
			merged = append(merged, otherVal)
		}
		if otherVal <= val {
			j += 1
		}
		i += 1
	}

	restOfSlice := ints
	k := i
	if j < len(otherInts) {
		restOfSlice = otherInts
		k = j
	}

	for k < len(restOfSlice) {
		lastVal := merged[len(merged)-1]
		val := restOfSlice[k]
		if val > lastVal {
			merged = append(merged, val)
		}
		k += 1
	}

	return merged
}
