package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	start := 1
	end := 15
	got := FizzBuzz(start, end)
	want := "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, FizzBuzz"

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
