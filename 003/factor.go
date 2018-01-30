package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	var num int64

	num = 600851475143
	maxNum := int64(math.Sqrt(float64(num)))

	for n := maxNum; n > 1; n-- {
		if num%n == 0 && fermatPrime(n) {
			fmt.Println("Result: ", n)
			break
		}
	}
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
