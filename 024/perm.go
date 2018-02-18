// Note: This is not very efficient. Should add a "maxdepth" to permutate which
// returns immediately once we compute the millionth permutation.

package main

import "fmt"

func main() {
	res := permutate([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(res[1e6-1])
}

func permutate(digits []int) [][]int {
	var ret [][]int

	// If we only have two left, immediately return both permutations.
	if len(digits) <= 2 {
		ret = append(ret, digits)
		ret = append(ret, []int{digits[1], digits[0]})
		return ret
	}

	for i := 0; i < len(digits); i++ {
		// Remaining digits.
		current := digits[i]

		// Remove current from ndigits
		ndigits := []int{}
		for k := range digits {
			if k != i {
				ndigits = append(ndigits, digits[k])
			}
		}

		for _, perms := range permutate(ndigits) {
			s := append([]int{current}, perms...)
			ret = append(ret, s)
		}
	}
	return ret
}
