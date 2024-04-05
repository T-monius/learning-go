package odd

import "testing"

func TestOdd(t *testing.T) {
	num := 1
	got := Odd(num)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestOdd2(t *testing.T) {
	num := 2
	got := Odd(num)
	want := false

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestOddNeg(t *testing.T) {
	num := -1
	got := Odd(num)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}

func TestOddOtherNeg(t *testing.T) {
	num := -17
	got := Odd(num)
	want := true

	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
