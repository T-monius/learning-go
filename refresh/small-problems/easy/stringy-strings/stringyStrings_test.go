package stringystrings

import (
	"testing"
)

func TestStringyStrings(t *testing.T) {
	t.Run("test num: 0", func(t *testing.T) {
		var num int
		got := StringyStrings(num)
		want := ""

		assertCorrectBinaryStr(t, num, got, want)
	})

	t.Run("test num: 6", func(t *testing.T) {
		num := 6
		got := StringyStrings(num)
		want := "101010"

		assertCorrectBinaryStr(t, num, got, want)
	})

	t.Run("test num: 9", func(t *testing.T) {
		num := 9
		got := StringyStrings(num)
		want := "101010101"

		assertCorrectBinaryStr(t, num, got, want)
	})

	t.Run("test num: 4", func(t *testing.T) {
		num := 4
		got := StringyStrings(num)
		want := "1010"

		assertCorrectBinaryStr(t, num, got, want)
	})

	t.Run("test num: 7", func(t *testing.T) {
		num := 7
		got := StringyStrings(num)
		want := "1010101"

		assertCorrectBinaryStr(t, num, got, want)
	})
}

func assertCorrectBinaryStr(t testing.TB, num int, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if len(want) != num {
		t.Errorf("Output an invalid string length. Want len: %d\n", len(want))
	}
}
