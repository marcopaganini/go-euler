package main

import "fmt"

func fibo(total, a, b int) int {
	sum := a + b
	if sum > 4e6 {
		return total
	}
	fmt.Println(sum)
	if sum&1 == 0 {
		total += sum
	}
	return fibo(total, b, sum)
}

func main() {
	// Add two in "1, 2"
	fmt.Println("Sum of even numbers is", fibo(0, 1, 2)+2)
}
