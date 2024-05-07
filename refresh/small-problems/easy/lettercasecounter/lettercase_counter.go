package lettercasecounter

const UpperMax int = 90
const UpperMin int = 65
const LowerMin int = 97
const LowerMax int = 122
const lower string = "lowercase"
const upper string = "uppercase"
const neither string = "neither"

func LetterCaseCount(str string) map[string]int {
	counts := map[string]int{
		lower:   0,
		upper:   0,
		neither: 0,
	}

	for _, b := range str {
		a := int(b)

		switch {
		case (a >= UpperMin) && (a <= UpperMax):
			counts[upper] += 1
		case (a >= LowerMin) && (a <= LowerMax):
			counts[lower] += 1
		default:
			counts[neither] += 1
		}
	}

	return counts
}
