package sumOfDigits

func SumOfDigits(num int) int {
	var total int

	for num > 0 {
		dig := num % 10

		total += dig
		num /= 10
	}

	return total
}
