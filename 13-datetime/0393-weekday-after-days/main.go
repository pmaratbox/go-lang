package main

import "fmt"

func main() {
	names := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	start := 6 // Saturday
	result := (start + 3) % 7
	fmt.Println(names[result])
}
