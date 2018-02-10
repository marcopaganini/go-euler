package main

import (
	"fmt"
	"math/big"
)

func main() {
	fact := big.NewInt(1)
	for x := int64(100); x > 1; x-- {
		fact = fact.Mul(fact, big.NewInt(x))
	}
	sum := 0
	for _, r := range fmt.Sprint(fact) {
		sum += int(r) - 0x30
	}
	fmt.Println(sum)
}
