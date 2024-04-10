package alphabeticalnumbers

import (
	"slices"
	"testing"
)

func TestAlphabeticalNumbers(t *testing.T) {
	t.Run("test short array: [0,1,2]", func(t *testing.T) {
		nums := []int{0, 1, 2}
		got := AlphabeticalNumbers(nums)
		want := []int{1, 2, 0}

		if slices.Equal(got, want) != true {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
