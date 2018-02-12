package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	fdata, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		log.Fatalln(err)
	}

	fs := strings.Split(string(fdata), ",")
	sort.Strings(fs)

	var sum int
	for k, w := range fs {
		sum += wordweight(w) * (k + 1)
	}
	fmt.Println(sum)
}

func wordweight(s string) int {
	var sum int
	for _, c := range s {
		if c != '"' {
			sum += int(c) - 0x40
		}
	}
	return sum
}
