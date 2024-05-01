package removevowels

import (
	"slices"
	"testing"
)

func TestRemoveVowels(t *testing.T) {
	tests := []struct {
		description string
		strings     []string
		want        []string
	}{
		{
			description: "The abcs: abcdefghijklmnopqrstuvwxyz",
			strings:     []string{"abcdefghijklmnopqrstuvwxyz"},
			want:        []string{"bcdfghjklmnpqrstvwxyz"},
		},
		{
			description: "The colors: green YELLOW black white",
			strings:     []string{"green", "YELLOW", "black", "white"},
			want:        []string{"grn", "YLLW", "blck", "wht"},
		},
		{
			description: "The consonants and vowels: ABC AEIOU XYZ",
			strings:     []string{"ABC", "AEIOU", "XYZ"},
			want:        []string{"BC", "", "XYZ"},
		},
	}

	for _, test := range tests {
		got := RemoveVowels(test.strings)

		if slices.Compare(got, test.want) != 0 {
			t.Errorf("got %v want %v", got, test.want)
		}
	}
}
