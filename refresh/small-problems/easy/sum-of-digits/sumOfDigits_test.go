package sumOfDigits

import (
	"testing"
)

func TestSumOfDigits(t *testing.T) {
	t.Run("test a positive number: 23", func(t *testing.T) {
		got := SumOfDigits(23)
		want := 5

		assertCorrectSum(t, got, want)
	})

	t.Run("test anoter positive number: 496", func(t *testing.T) {
		got := SumOfDigits(496)
		want := 19

		assertCorrectSum(t, got, want)
	})

	t.Run("test yet another positive number: 123_456_789", func(t *testing.T) {
		got := SumOfDigits(123_456_789)
		want := 45

		assertCorrectSum(t, got, want)
	})
}

func assertCorrectSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
