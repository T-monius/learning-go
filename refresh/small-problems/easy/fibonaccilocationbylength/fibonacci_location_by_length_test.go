package fibonaccilocationbylength

import "testing"

func TestFindFibonacciIndexByLength(t *testing.T) {
	defaultParams := fibParams{prior: 1, pen: 1, i: 3}
	tests := []struct {
		description string
		l           int
		want        int
	}{
		{
			description: "Index of first 2 digit num",
			l:           2,
			want:        7,
		},
		{
			description: "Index of first 10 digit num",
			l:           10,
			want:        45,
		},
		{
			description: "Index of first 20 digit num",
			l:           20,
			want:        93,
		},
	}

	for _, test := range tests {
		got := FindFibonacciIndexByLength(test.l, defaultParams)

		if got != test.want {
			t.Errorf("got %d want %d", got, test.want)
		}
	}
}
