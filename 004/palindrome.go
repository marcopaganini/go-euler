package main

import "fmt"

func main() {
	max := 0

	for n1 := 100; n1 < 1000; n1++ {
		for n2 := 100; n2 < 1000; n2++ {
			n := n1 * n2
			if n > max && palindrome(n) {
				max = n
			}
		}
	}
	fmt.Println("Result is", max)
}

func palindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-(i+1)] {
			return false
		}
	}
	return true
}
