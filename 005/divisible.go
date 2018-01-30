package main

import "fmt"

func main() {
	num := 20
	var div int

	for {
		for div = 2; div <= 20; div++ {
			if num%div != 0 {
				break
			}
		}
		if div > 20 {
			fmt.Println("Result:", num)
			break
		}
		num++
	}
}
