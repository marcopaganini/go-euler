package main

import "fmt"

func main() {
	sundays := 0
	d := 6
	m := 1
	y := 1901

	for y <= 2000 {
		if d == 1 {
			sundays++
		}
		mdays := monthDays(m, y)
		d += 7
		if d > mdays {
			d %= mdays
			m++
			if m > 12 {
				m = 1
				y++
			}
		}
	}
	fmt.Println(sundays)
}

func monthDays(month, year int) int {
	mday := map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	// Leap year
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		mday[2] = 29
	}

	numdays, ok := mday[month]
	if ok {
		return numdays
	}
	return -1
}
