package odd

func Odd(num int) bool {
	mod := num % 2

	return mod == 1 || mod == -1
}
