package hextoint

import "testing"

func TestHextToInt(t *testing.T) {
	t.Run("Get decimal of D1CE", func(t *testing.T) {
		hex := "D1CE"
		got := HexToInt(hex)
		want := 53710

		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	})

	t.Run("Get decimal of 2E6", func(t *testing.T) {
		hex := "2E6"
		got := HexToInt(hex)
		want := 742

		if got != want {
			t.Errorf("got %d want %d\n", got, want)
		}
	})
}
