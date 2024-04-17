package lettercounter

import (
	"reflect"
	"testing"
)

func TestLetterCounter(t *testing.T) {
	tests := []struct {
		str  string
		want map[int]int
	}{
		{
			str: "Four score and seven.",
			want: map[int]int{
				3: 1, 4: 1, 5: 1, 6: 1,
			},
		},
		{
			str: "Hey diddle diddle, the cat and the fiddle!",
			want: map[int]int{
				3: 5, 6: 1, 7: 2,
			},
		},
		{
			str: "What's up doc?",
			want: map[int]int{
				6: 1, 2: 1, 4: 1,
			},
		},
		{
			str:  "",
			want: map[int]int{},
		},
	}

	for _, test := range tests {
		got := LetterCounter(test.str)
		assertMap(t, got, test.want)
	}

}

func assertMap(t testing.TB, got, want map[int]int) {
	t.Helper()

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %v want %v", got, want)
	}
}
