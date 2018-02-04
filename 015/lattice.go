// Note: This can be simplified to 2n!/(n!)^2 (http://oeis.org)
// This take TOO LONG (24m)

package main

import "fmt"

const (
	xmax = 3
	ymax = 3
)

var counter int
var depth int

func main() {
	navigate(0, 0)
	fmt.Println(counter)
}

func navigate(x, y int) {
	depth++
	//fmt.Println(depth, ":", x, y)
	// At destination? Increment and stop
	if x == xmax || y == ymax {
		//fmt.Println(depth, "Bingo! Returning")
		counter++
		depth--
		return
	}
	// Go right if we can
	if x < xmax {
		//fmt.Println(depth, "Calling with X incremented", x+1, y)
		navigate(x+1, y)
	}
	// Go down if we can
	if y < ymax {
		//fmt.Println(depth, "Calling with y incremented", x, y+1)
		navigate(x, y+1)
	}
	//fmt.Println(depth, "Returning")
	depth--
	return
}
