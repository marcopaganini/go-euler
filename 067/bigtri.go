package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("p067_triangle.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var tri [][]int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		trinum := []int{}
		for _, v := range strings.Split(line, " ") {
			n, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			trinum = append(trinum, n)
		}
		tri = append(tri, trinum)
	}

	// Calculate
	for y := len(tri) - 2; y >= 0; y-- {
		for x := 0; x < len(tri[y]); x++ {
			left := tri[y+1][x]
			right := tri[y+1][x+1]
			if left > right {
				tri[y][x] += left
			} else {
				tri[y][x] += right
			}
		}
	}

	fmt.Println("Max is", tri[0][0])
}
