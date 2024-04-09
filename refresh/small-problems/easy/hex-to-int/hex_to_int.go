package hextoint

import "math"

var hexToDec = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
}

func HexToInt(hex string) int {
	var decimal int

	for i := 0; i < len(hex); i += 1 {
		distToEnd := (len(hex) - 1) - i
		pow := int(math.Pow(16, float64(distToEnd)))
		char := string(hex[i])
		decimal += pow * (hexToDec[char])
	}

	return decimal
}
