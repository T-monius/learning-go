package asciistringvalue

func ASCIIStringValue(str string) int {
	var sum int

	for _, char := range str {
		asciivalue := int(char)
		sum += asciivalue
	}

	return sum
}
