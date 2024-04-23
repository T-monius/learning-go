package combineslice

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	sl := []int{1, 3, 5}
	sl1 := []int{3, 6, 9}
	got := Merge(sl, sl1)
	want := []int{1, 3, 5, 6, 9}

	if slices.Compare(got, want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}
