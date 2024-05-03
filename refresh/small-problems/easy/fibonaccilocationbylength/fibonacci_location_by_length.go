package fibonaccilocationbylength

import "fmt"

type fibParams struct {
	prior int
	pen   int
	i     int
}

func FindFibonacciIndexByLength(l int, p fibParams) int {
	c := p.prior + p.pen
	if len(fmt.Sprint(c)) == l {
		return p.i
	}

	return FindFibonacciIndexByLength(l, fibParams{prior: c, pen: p.prior, i: p.i + 1})
}
