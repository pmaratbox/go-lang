package main

import "fmt"

func main() {
	total := 10*60 + 45
	total = (total + 90) % (24 * 60)
	h, m := total/60, total%60
	fmt.Printf("%02d:%02d\n", h, m)
}
