package doubleconsonants

var consonants map[string]bool = map[string]bool{
	"B": true,
	"C": true,
	"D": true,
	"F": true,
	"G": true,
	"H": true,
	"J": true,
	"K": true,
	"L": true,
	"M": true,
	"N": true,
	"P": true,
	"Q": true,
	"R": true,
	"S": true,
	"T": true,
	"V": true,
	"W": true,
	"X": true,
	"Y": true,
	"Z": true,
	"b": true,
	"c": true,
	"d": true,
	"f": true,
	"g": true,
	"h": true,
	"j": true,
	"k": true,
	"l": true,
	"m": true,
	"n": true,
	"p": true,
	"q": true,
	"r": true,
	"s": true,
	"t": true,
	"v": true,
	"w": true,
	"x": true,
	"y": true,
	"z": true,
}

func DoubleConsonants(str string) string {
	doubled := ""

	for _, b := range str {
		char := string(b)
		if consonants[char] {
			doubled += char
		}
		doubled += char
	}

	return doubled
}
