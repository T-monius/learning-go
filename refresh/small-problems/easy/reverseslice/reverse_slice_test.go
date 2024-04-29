package reverseslice

import (
	"slices"
	"testing"
)

func TestBackwards(t *testing.T) {
	list := []int{1, 2, 3, 4}
	got := Backwards(list)
	want := []int{4, 3, 2, 1}

	if slices.Compare(got, want) != 0 {
		t.Errorf("got %v want %v", got, want)
	}
}
