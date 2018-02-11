package main

import "fmt"

func main() {
	var asum int

	for n := 0; n < 10000; n++ {
		nsum := sum(divisors(n))
		if n == nsum {
			continue
		}
		dsum := sum(divisors(nsum))
		if dsum == n {
			asum += dsum
		}
	}
	fmt.Println(asum)
}

func divisors(n int) []int {
	ret := []int{1}

	limit := n
	for d := 2; d < limit; d++ {
		if n%d == 0 {
			ret = append(ret, d)
			limit = n / d
			ret = append(ret, limit)
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
