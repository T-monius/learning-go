package reverseslice

func Backwards(list []int) []int {
	for s, e := 0, len(list)-1; s < e; s, e = s+1, e-1 {
		list[s], list[e] = list[e], list[s]
	}

	return list
}
