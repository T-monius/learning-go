package multiplicativeaverage

import "testing"

func TestMultiplicativeAverage(t *testing.T) {
	tests := []struct {
		description string
		want        string
		nums        []int
	}{
		{
			description: "Average two numbers: 3, 5",
			nums:        []int{3, 5},
			want:        "7.500",
		},
		{
			description: "Average one number: 6",
			nums:        []int{6},
			want:        "6.000",
		},
		{
			description: "Average several numbers: 2, 5, 7, 11, 13, 17",
			nums:        []int{2, 5, 7, 11, 13, 17},
			want:        "28361.667",
		},
	}

	for _, test := range tests {
		got := MultiplicativeAverage(test.nums)

		if got != test.want {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
