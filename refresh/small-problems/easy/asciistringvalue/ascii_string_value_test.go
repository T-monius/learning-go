package asciistringvalue

import "testing"

func TestASCIIStringValue(t *testing.T) {
	testCases := []struct {
		message string
		str     string
		want    int
	}{
		{message: "ASCII val: Four Score", str: "Four score", want: 984},
		{message: "ASCII val: Launch School", str: "Launch School", want: 1251},
		{message: "ASCII val: a", str: "a", want: 97},
		{message: "ASCII val: ab", str: "ab", want: 195},
		{message: "ASCII val for empty string", str: "", want: 0},
	}

	for _, testCase := range testCases {
		t.Run(testCase.message, func(t *testing.T) {
			got := ASCIIStringValue(testCase.str)

			assertInt(t, got, testCase.want)
		})
	}

}

func assertInt(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
