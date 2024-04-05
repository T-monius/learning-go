package stringystrings

import (
	"testing"
)

func TestStringyStrings(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		var num int
		got := StringyStrings(num)
		want := "something"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
