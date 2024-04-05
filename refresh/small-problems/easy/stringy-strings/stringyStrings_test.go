package stringystrings

import (
	"testing"
)

func TestStringyStrings(t *testing.T) {
	t.Run("test num: 0", func(t *testing.T) {
		var num int
		got := StringyStrings(num)
		want := ""

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if len(want) != num {
			t.Errorf("Output an invalid string length. Want len: %d\n", len(want))
		}
	})

	t.Run("test num: 6", func(t *testing.T) {
		num := 6
		got := StringyStrings(num)
		want := "101010"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

		if len(want) != num {
			t.Errorf("Output an invalid string length. Want len: %d\n", len(want))
		}
	})
}
