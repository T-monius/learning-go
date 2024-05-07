package lettercasecounter

import (
	"reflect"
	"testing"
)

func TestLetterCaseCount(t *testing.T) {
	tests := []struct {
		description string
		str         string
		want        map[string]int
	}{
		{
			description: "A string of abcs spaces and numbers: abCdef 123",
			str:         "abCdef 123",
			want:        map[string]int{"lowercase": 5, "uppercase": 1, "neither": 4},
		},
		{
			description: "A string of alphas, spaces, and special chars: AbCd +Ef",
			str:         "AbCd +Ef",
			want:        map[string]int{"lowercase": 3, "uppercase": 3, "neither": 2},
		},
		{
			description: "A string of numbers: 123",
			str:         "123",
			want:        map[string]int{"lowercase": 0, "uppercase": 0, "neither": 3},
		},
		{
			description: "An empty string",
			str:         "",
			want:        map[string]int{"lowercase": 0, "uppercase": 0, "neither": 0},
		},
	}

	for _, test := range tests {
		got := LetterCaseCount(test.str)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
