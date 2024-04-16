package dailydouble

import "testing"

func TestSqueezeDuplicates(t *testing.T) {
	duplicateTests := []struct {
		str  string
		want string
	}{
		{
			str: "ddaaiillyy ddoouubbllee", want: "daily double",
		},
		{
			str: "4444abcabccba", want: "4abcabcba",
		},
		{
			str: "ggggggggggggggg", want: "g",
		},
		{
			str: "a", want: "a",
		},
		{
			str: "", want: "",
		},
	}

	for _, testDup := range duplicateTests {
		t.Run("Squeeze: ddaaiillyy ddoouubbllee", func(t *testing.T) {
			got := SqueezeDuplicates(testDup.str)

			assertString(t, got, testDup.want)
		})
	}
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
