package sumOfDigits

import (
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	got := SumOfDigits(23)
	want := 5

	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}

func TestSumOfDigitsAgain(t *testing.T) {
	got := SumOfDigits(496)
	want := 19

	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}

func TestSumOfDigitsYetAgain(t *testing.T) {
	got := SumOfDigits(123_456_789)
	want := 45

	if got != want {
		t.Errorf("got %d want %d\n", got, want)
	}
}
