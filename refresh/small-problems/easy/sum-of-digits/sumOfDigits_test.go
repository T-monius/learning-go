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
