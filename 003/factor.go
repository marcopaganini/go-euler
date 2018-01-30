package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	var (
		num      int64
		maxprime int64
	)

	num = 600851475143
	limit := num

	for n := int64(2); n <= limit; n++ {
		if num%n == 0 {
			limit = num / n
			if fermatPrime(n) {
				maxprime = n
			}
		}
	}
	fmt.Println("Result:", maxprime)
}

func fermatPrime(p int64) bool {
	const iterations = 20

	// Even number? Return composite
	if p&1 == 0 {
		//fmt.Println(p, "is even. Returning false.")
		return false
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	ba := big.NewInt(0)

	for i := 0; i < iterations; i++ {
		ba.Rand(rnd, big.NewInt(p))
		res := ba.Exp(ba, big.NewInt(p-1), big.NewInt(p))
		//fmt.Println("Tested ", p, " with random a=", ba, "result is ", res)
		if res.Uint64() != 1 {
			return false
		}
	}
	return true
}
