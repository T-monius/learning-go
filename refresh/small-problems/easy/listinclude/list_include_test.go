package listinclude

import "testing"

func TestDoesInclude(t *testing.T) {
	tests := []struct {
		description string
		nums        []int
		target      int
		want        bool
	}{
		{
			description: "target 3 present in sorted array",
			nums:        []int{1, 2, 3, 4, 5},
			target:      3,
			want:        true,
		},
		{
			description: "target 4 present in sorted array",
			nums:        []int{1, 2, 3, 4, 5},
			target:      4,
			want:        true,
		},
		{
			description: "target 3 present in unsorted array",
			nums:        []int{2, 3, 1, 5, 4},
			target:      3,
			want:        true,
		},
		{
			description: "target 6 present",
			nums:        []int{1, 2, 3, 4, 5},
			target:      6,
			want:        false,
		},
	}

	for _, test := range tests {
		got := DoesInclude(test.nums, test.target)

		t.Run(test.description, func(t *testing.T) {
			if got != test.want {
				t.Errorf("got %t want %t", got, test.want)
			}
		})
	}
}
