package main

import "fmt"

type Color struct {
	r, g, b int
}

// fromHex is a static factory building a Color from "#rrggbb".
func fromHex(hex string) Color {
	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)
	return Color{r, g, b}
}

func main() {
	c := fromHex("#ff0000")
	fmt.Println(c.r, c.g, c.b)
}
