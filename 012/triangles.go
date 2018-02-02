package main

import (
	"fmt"
)

func main() {
	var trinum int

	for n := 1; ; n++ {
		trinum += n
		if numFactors(trinum) > 500 {
			break
		}
	}
	fmt.Printf("Result: %d\n", trinum)
}

func numFactors(n int) int {

	// Special case 1
	if n == 1 {
		return 1
	}

	var tot int
	limit := n
	for factor := 1; factor < limit; factor++ {
		if n%factor == 0 {
			limit = n / factor
			tot += 2
		}
	}
	return tot
}
