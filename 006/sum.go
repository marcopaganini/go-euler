package main

import "fmt"

func main() {
	var (
		sum1 int64
		sum2 int64
	)
	for n := int64(1); n <= 100; n++ {
		sum1 += n * n
		sum2 += n
	}
	sum2 = sum2 * sum2
	fmt.Println(sum1, sum2, sum2-sum1)
}
