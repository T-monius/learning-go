package numbertostring

import "testing"

func TestNumberToString(t *testing.T) {
	t.Run("test zero", func(t *testing.T) {
		num := 0
		got := NumberToString(num)
		want := "0"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test 4321", func(t *testing.T) {
		num := 4321
		got := NumberToString(num)
		want := "4321"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("test 5000", func(t *testing.T) {
		num := 5000
		got := NumberToString(num)
		want := "5000"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func TestDigitsFromNumber(t *testing.T) {
	t.Run("test 4321", func(t *testing.T) {
		num := 4321
		got := DigitsFromNumber(num)
		want := []int{1, 2, 3, 4}

		if got[0] != want[0] || len(got) != len(want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
