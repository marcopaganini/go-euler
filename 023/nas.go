package main

import (
	"fmt"
	"sort"
)

const (
	limit = 28123
)

func main() {
	anums := []int{}
	amap := make([]bool, limit)

	for n := 1; n < limit; n++ {
		div := divisors(n)
		sort.Ints(div)
		if sum(div) > n {
			//fmt.Println(n, div, sum(div))
			anums = append(anums, n)
		}
	}
	//fmt.Println(len(anums), "abundant numbers")

	// Find all possible sums of two numbers less than limit
	for _, i := range anums {
		for _, j := range anums {
			if i+j < limit {
				amap[i+j] = true
			}
		}
	}

	// Print all numbers that are not a sum.
	sum := 0
	for i, v := range amap {
		if !v {
			sum += i
		}
	}
	fmt.Println("Total", sum)
}

func divisors(n int) []int {
	ret := []int{1}

	limit := n
	for d := 2; d < limit; d++ {
		if n%d == 0 {
			ret = append(ret, d)
			limit = n / d
			if limit != d {
				ret = append(ret, limit)
			}
		}
	}
	return ret
}

func sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}
