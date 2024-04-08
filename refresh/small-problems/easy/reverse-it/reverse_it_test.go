package reverseit

import "testing"

func TestReverseit(t *testing.T) {
	t.Run("Reverse an empty string", func(t *testing.T) {
		str := ""
		got := ReverseIt(str)
		want := ""

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Reverse string: Hello World", func(t *testing.T) {
		str := "Hello World"
		got := ReverseIt(str)
		want := "World Hello"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Reverse string: Reverse these words", func(t *testing.T) {
		str := "Reverse these words"
		got := ReverseIt(str)
		want := "words these Reverse"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
