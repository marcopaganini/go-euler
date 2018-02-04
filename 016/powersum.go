package main

import (
	"fmt"
	"math/big"
)

func main() {
	n := big.NewInt(0)
	n.Exp(big.NewInt(2), big.NewInt(1000), nil)

	sum := 0
	s := fmt.Sprint(n)
	for _, c := range s {
		sum += int(c) - 0x30
	}
	fmt.Println(sum)
}
