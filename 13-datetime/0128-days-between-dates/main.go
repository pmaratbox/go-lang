package main

import "fmt"

// daysFromCivil returns days since 1970-01-01 (Howard Hinnant's algorithm).
func daysFromCivil(y, m, d int) int {
	if m <= 2 {
		y--
	}
	era := y
	if y < 0 {
		era = y - 399
	}
	era /= 400
	yoe := y - era*400
	var mp int
	if m > 2 {
		mp = m - 3
	} else {
		mp = m + 9
	}
	doy := (153*mp+2)/5 + d - 1
	doe := yoe*365 + yoe/4 - yoe/100 + doy
	return era*146097 + doe - 719468
}

func main() {
	a := daysFromCivil(2000, 1, 1)
	b := daysFromCivil(2000, 12, 31)
	fmt.Println(b - a)
}
