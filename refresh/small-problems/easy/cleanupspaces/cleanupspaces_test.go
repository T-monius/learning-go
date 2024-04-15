package cleanupspaces

import "testing"

func TestCleanupSpaces(t *testing.T) {
	messy := "---what's my +*& line?"
	got := CleanupSpaces(messy)
	want := " what s my line "

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func TestIsAlpha(t *testing.T) {
	t.Run("test an ascii val: 97", func(t *testing.T) {
		asciiVal := 97
		got := IsAlpha(asciiVal)
		want := true

		assertBool(t, got, want)
	})
	t.Run("test a non-ascii val: 91", func(t *testing.T) {
		nonAsciiVal := 91
		got := IsAlpha(nonAsciiVal)
		want := false

		assertBool(t, got, want)
	})
}

func TestFindNextAlphaIndex(t *testing.T) {
	messy := "---what's my +*& line?"
	currentIdx := 0
	got := FindNextAlphaIndex(messy, currentIdx)
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertBool(t testing.TB, got, want bool) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
