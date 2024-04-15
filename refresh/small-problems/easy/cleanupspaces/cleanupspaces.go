package cleanupspaces

const space string = " "
const upperMin int = 65
const upperMax int = 90
const lowerMin int = 97
const lowerMax int = 122

func IsAlpha(charVal int) bool {
	if (charVal <= upperMax && charVal >= upperMin) || (charVal <= lowerMax && charVal >= lowerMin) {
		return true
	}

	return false
}

func FindNextAlphaIndex(str string, currentIdx int) int {
	for i := currentIdx + 1; i < len(str); i += 1 {
		otherAsciiVal := int(str[i])
		if IsAlpha(otherAsciiVal) {
			return i
		}
	}
	return len(str)
}

func CleanupSpaces(str string) string {
	clean := ""

	for i := 0; i < len(str); i += 1 {
		byteVal := str[i]
		asciiVal := int(byteVal)
		if IsAlpha(asciiVal) {
			char := string(byteVal)
			clean += char
			continue
		}

		clean += space
		i = FindNextAlphaIndex(str, i) - 1
	}

	return clean
}
