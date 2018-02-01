package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	const maxnum = 2e6

	var (
		sum int64
		n   int64
	)

	// Already includes 2 + 3
	sum = 2
	for n = 5; n < maxnum; n += 2 {
		if fermatPrime(big.NewInt(int64(n))) {
			sum += n
			//fmt.Println("Adding", n, "sum is now", sum)
		}
	}

	fmt.Println("Result:", sum)
}

func fermatPrime(p *big.Int) bool {
	const iterations = 64

	// Any positive number less than 4 is prime by default.
	if p.Cmp(big.NewInt(4)) == -1 {
		return true
	}

	// Even number? Return composite
	z := big.NewInt(0)
	z.Rem(p, big.NewInt(2))
	if z.Cmp(big.NewInt(0)) == 0 {
		return false
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	a := big.NewInt(0)
	res := big.NewInt(0)

	for i := 0; i < iterations; i++ {
		// Generate a as a random from [1-p)
		z.Sub(p, big.NewInt(1))
		a.Rand(rnd, z)
		a.Add(a, big.NewInt(1))

		z.Sub(p, big.NewInt(1))
		res.Exp(a, z, p)
		//fmt.Println("Tested ", p, " with random a=", a, "result is ", res)
		if res.Cmp(big.NewInt(1)) != 0 {
			return false
		}
	}
	return true
}
