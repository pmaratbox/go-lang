package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	width := 10
	filled := int(math.Round(float64(width) * 0.4))
	bar := strings.Repeat("#", filled) + strings.Repeat("-", width-filled)
	fmt.Printf("[%s]\n", bar)
}
