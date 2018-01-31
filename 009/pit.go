package main

import (
	"fmt"
	"math"
)

func main() {
OuterLoop:
	for a := 1; a < 1000; a++ {
		for b := 1; b < 1000; b++ {
			c2 := a*a + b*b
			c := math.Sqrt(float64(c2))
			if float64(a)+float64(b)+c == 1000 {
				fmt.Printf("Result: a=%d b=%d c=%f product=%.0f\n", a, b, c, float64(a)*float64(b)*c)
				break OuterLoop
			}
		}
	}
}
