package doubleconsonants

import "testing"

func TestDouleConsonants(t *testing.T) {
	tests := []struct {
		want        string
		str         string
		description string
	}{
		{
			str:         "String",
			want:        "SSttrrinngg",
			description: "Just alphas",
		},
		{
			str:         "Hello-World!",
			want:        "HHellllo-WWorrlldd!",
			description: "A string of alphas and special characters",
		},
		{
			str:         "July 4th",
			want:        "JJullyy 4tthh",
			description: "Alpha numeric and special characters",
		},
		{
			str:         "",
			want:        "",
			description: "An empty string",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			got := DoubleConsonants(test.str)

			if got != test.want {
				t.Errorf("got %s want %s", got, test.want)
			}
		})
	}
}
