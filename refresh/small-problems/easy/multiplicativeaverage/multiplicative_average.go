package multiplicativeaverage

import (
	"fmt"
	"math"
)

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func MultiplicativeAverage(nums []int) string {
	multiple := 1

	for _, num := range nums {
		multiple *= num
	}
	avg := float64(multiple) / float64(len(nums))
	avg = Round(avg, 0.001)

	return fmt.Sprintf("%.3f", avg)
}
