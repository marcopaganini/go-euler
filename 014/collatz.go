package main

import "fmt"

func main() {
	const (
		start = 13
		limit = 1e6
	)

	maxlen := 0
	maxn := 0
	for n := start; n < limit; n++ {
		length := collatz(n)
		if length > maxlen {
			maxn = n
			maxlen = length
		}
	}
	fmt.Println("Result: ", maxn, maxlen)
}

func collatz(n int) int {
	count := 1

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		count++
		if n == 1 {
			break
		}
	}
	return count
}
