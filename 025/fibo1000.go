package main

import (
	"fmt"
	"math/big"
)

// bigfibo calculates fibonnaci numbers until the number is has at least maxlen
// digits. It returns the iteration number and the number found.
func bigfibo(maxlen int) (int, *big.Int) {
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := big.NewInt(0)

	var iter int
	for {
		iter++
		c.Add(a, b)
		if len(c.String()) >= maxlen {
			return iter + 1, c
		}
		a.Set(b)
		b.Set(c)
	}
}

func main() {
	fmt.Println(bigfibo(1000))
}
