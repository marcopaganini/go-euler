// recycles.go - Please note that this is very un-mathematical and uses brute
// force.  I originally set out to write a generic "longest repeating string
// chain" algorithm and somewhat succeeded. This could be useful for future
// euler challenger (who knows...)

package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	var (
		longnum int
		maxlen  int
	)

	for i := 1; i < 1000; i++ {
		a := big.NewRat(1, int64(i))
		chain, err := cycle(a.FloatString(2000)[2:])
		if err != nil {
			log.Fatalln(err)
		}
		if chain != "" {
			//fmt.Printf("num: %d, chain: [%s]\n", i, chain)
			if len(chain) > maxlen {
				maxlen = len(chain)
				longnum = i
			}
		}
	}
	fmt.Printf("Longest number: %d, chain length: %d\n", longnum, maxlen)
}

func cycle(s string) (string, error) {
	for start := 0; start < len(s)/2; start++ {

		limitlen := (len(s) - start) / 2
		for clen := 1; clen <= limitlen; clen++ {
			original := s[start : start+clen]
			found := true
			count := 0
			for cs := start; cs < (len(s) - clen); cs += clen {
				sub := s[cs : cs+clen]
				if sub != original {
					found = false
					break
				}
				if found {
					count++
				}
			}
			if found && count > 1 {
				return original, nil
			}
		}
	}
	return "", nil
}
