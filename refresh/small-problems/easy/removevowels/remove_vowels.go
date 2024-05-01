package removevowels

var vowels map[int]bool = map[int]bool{
	97:  true,
	101: true,
	105: true,
	111: true,
	117: true,
	65:  true,
	69:  true,
	73:  true,
	79:  true,
	85:  true,
}

func RemoveVowels(strings []string) []string {
	noVowels := make([]string, 0)
	for _, str := range strings {
		bytes := []byte(str)
		w := 0
		for r := 0; r < len(bytes); r += 1 {
			char := bytes[r]
			_, ok := vowels[int(char)]
			if !ok {
				bytes[w] = char
				w += 1
			}
		}
		noVowels = append(noVowels, string(bytes[0:w]))
	}
	return noVowels
}
