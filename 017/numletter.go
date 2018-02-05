package main

import (
	"fmt"
	"strings"
)

func main() {
	count := 0
	for n := 1; n <= 1000; n++ {
		count += len(spellnum(n))
	}
	fmt.Println(count)
}

func spellnum(n int) string {
	numbers := []string{
		"zero", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen",
	}
	deci := []string{
		"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
	}

	var ret []string

	if n == 1000 {
		return "onethousand"
	}

	if n > 99 {
		hundreds := n / 100
		n %= 100
		ret = append(ret, numbers[hundreds], "hundred")
	}

	if n > 0 && len(ret) > 0 {
		ret = append(ret, "and")
	}

	if n >= 20 {
		tens := n / 10
		n %= 10
		ret = append(ret, deci[tens-2])
	}

	if n > 0 {
		ret = append(ret, numbers[n])
	}

	return strings.Join(ret, "")
}
